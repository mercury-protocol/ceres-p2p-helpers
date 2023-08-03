import { generateKeyPair, marshalPrivateKey, unmarshalPrivateKey, marshalPublicKey } from "@libp2p/crypto/keys";
import {peerIdFromKeys} from "@libp2p/peer-id";
import {createLibp2p} from 'libp2p';
import {tcp} from '@libp2p/tcp';
import {mplex} from '@libp2p/mplex';
import {yamux} from '@chainsafe/libp2p-yamux';
import {noise} from '@chainsafe/libp2p-noise';
import {multiaddr} from '@multiformats/multiaddr';
import {pipe} from 'it-pipe';
import {toString as uint8ArrayToString} from 'uint8arrays/to-string';
import {fromString as uint8ArrayFromString} from 'uint8arrays/from-string';
import {CID} from 'multiformats/cid';
import * as codec from '@ipld/dag-pb';
import {sha256} from 'multiformats/hashes/sha2';
import axios from "axios";
import fs from "fs";
import ethers from 'ethers';
import { recoverAddress } from 'ethers/lib/utils.js';
import daoAbi from "./dao.json" assert {type: "json" };
import 'dotenv/config';

class P2P {
    constructor(isBrowser) {
        this.node;
        this.isBrowser = isBrowser;
    }

    setup = async (port, keyBytes) => {
        const node = await this.createNode(this.isBrowser, port, keyBytes);
        await node.start();

        console.log(`Node peer ID is ${node.peerId.toString()}`);

        // Listen for new connections to peers
        node.addEventListener('peer:connect', (evt) => {
            const connection = evt.detail
            console.log(`Connected to ${connection.remotePeer.toString()}`)
        })

        // Listen for peers disconnecting
        node.addEventListener('peer:disconnect', (evt) => {
            const connection = evt.detail
            console.log(`Disconnected from ${connection.remotePeer.toString()}`)
        })

        this.node = node;
    }

    transferFile = async (targetMultiaddr, payload) => {
        const ma = multiaddr(targetMultiaddr);
        let resp;

        const len = payload.length;
        try {
            const stream = await this.node.dialProtocol(ma, '/mcy/transfer/dao/data/0.0.1')
            await pipe(
                // Source data
                [uint8ArrayFromString(`${len}\n${payload} \n`)],
                // write to the stream
                stream,
                // sink function
                async function (source) {
                    for await (const data of source) {
                        resp = uint8ArrayToString(data.subarray());
                    }
                }
            )
        } catch (e) {
            console.log("error: ", e);
        }

        if (!resp.includes("FILE_RECEIVED")) {
            return false;
        }
        return true;
    }

    sendDataCollectorMsgToNode = async (targetMultiaddr, msg) => {
        const ma = multiaddr(targetMultiaddr);
        let resp;

        try {
            const stream = await this.node.dialProtocol(ma, '/mcy/msg/dao/0.0.1');
            
            await pipe(
                // Source data
                [uint8ArrayFromString(`${msg}`)],
                // Write to the stream, and pass its output to the next function
                stream,
                // Sink function
                async function (source) {
                    // For each chunk of data
                    for await (const data of source) {
                        const respStr = uint8ArrayToString(data.subarray())
                        console.log("response: ", respStr);
                        resp = respStr;
                    }
                }
            )
        } catch(e) {
            console.log("error : ", e);
        }
        if (resp != undefined && !resp.includes("MSG_OK")) {
            return false;
        }
        return true;
    }

    cid = async (data) => {
        const hash = await sha256.digest(data);
        const cid = CID.create(1, codec.code, hash);
        return cid.toString();
    }

    createCollectorMsg = async (cid) => {
        const sig = await this.sign(cid, this.isBrowser);
        // create msg
        let msg = {
            cid: cid,
            sig: sig
        }

        const jsonStr = JSON.stringify(msg);
        
        return jsonStr;
    }

    getDaoAdminPeerInfo = async (baseUrl, address) => {
        axios.defaults.baseURL = baseUrl;
        const url = `/peer/${address}`;
        const res = await axios.get(url);
        return res.data.peer;
    }

    createFromPrivKey = async (privateKey) => {
        return peerIdFromKeys(marshalPublicKey(privateKey.public), marshalPrivateKey(privateKey))
    }

    createNode = async(isBrowser, port, keyBytes) => {
        if (isBrowser) {
            let key;
            if (keyBytes != undefined) {
                key = await unmarshalPrivateKey(keyBytes);
            } else {
                key = await generateKeyPair('RSA', 2048);
            }

            const peerId = await this.createFromPrivKey(key);

            const node = await createLibp2p({
                transports: [webTransport()],
                peerId: peerId,
                streamMuxers: [yamux(), mplex()],
                connectionEncryption: [noise()],
            })
            return node;
        } else {
             // check if file exists, if yes open and load key from that -> unmarshal
            const keyPath = "key.bin";
            let key;
            
            if (fs.existsSync(keyPath)) {
                const keyBytes = fs.readFileSync(keyPath);
                key = await unmarshalPrivateKey(keyBytes);
            } else {
                key = await generateKeyPair('RSA', 2048);
                const keyBytes = marshalPrivateKey(key);
                // save key to file
                fs.writeFileSync(keyPath, keyBytes);
            }

            const peerId = await this.createFromPrivKey(key);

            const node = await createLibp2p({
                transports: [
                    tcp()
                ],
                addresses: {
                    listen: [`/ip4/0.0.0.0/tcp/${port}`]
                },
                peerId: peerId,
                streamMuxers: [yamux(), mplex()],
                connectionEncryption: [
                    noise()
                ],
            })

            return node;
        }
    }

    sign = async(data, isBrowser) => {
        const signer = this.getSigner(isBrowser);
        const sig = await signer.signMessage(data);
        return sig;
    }
    
    verifyPeerInfo = async(addr, mu, sig, isBrowser) => {
        const signer = this.getSigner(isBrowser);
        const dao = new ethers.Contract(addr, daoAbi, signer);
        
        const hash = ethers.utils.keccak256(ethers.utils.toUtf8Bytes(mu));
        const recoveredAddress = recoverAddress(hash, sig);
        
        const hasRole = await dao.hasRole("0x0000000000000000000000000000000000000000000000000000000000000000", recoveredAddress);
        
        return hasRole;
    }
    
    getSigner = (isBrowser) => {
        if (isBrowser) {
            const provider = new ethers.providers.Web3Provider(window.ethereum);
            const signer = provider.getSigner();
            return signer;
        }
    
        const rpc = process.env.RPC_PROVIDER;
        const k = process.env.KEY;
        
        const provider = new ethers.providers.JsonRpcProvider(rpc);
        const signer = new ethers.Wallet(k, provider);
    
        return signer;
    }
}

export default P2P;
