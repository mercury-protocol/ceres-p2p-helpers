import os
import cid
import json
from dotenv import load_dotenv
from multiformats import *
from web3 import Web3
import eth_account.messages as messages

import multiaddr
import requests

from libp2p import new_host
from libp2p.typing import TProtocol
from libp2p.peer.peerinfo import info_from_p2p_addr
from libp2p.crypto.secp256k1 import create_new_key_pair

import trio
import libp2p.security.noise.transport as noise
from libp2p.transport.typing import TMuxerOptions, TSecurityOptions
from libp2p.crypto.ed25519 import create_new_key_pair as create_ed25519_key_pair
from libp2p.stream_muxer.mplex.mplex import MPLEX_PROTOCOL_ID, Mplex

DATA_TRANSFER_PROTOCOL = TProtocol("/mcy/transfer/dao/data/0.0.1")
DATA_COLLECTOR_MSG_PROTOCOL = TProtocol("/mcy/msg/dao/0.0.1")

async def run(port, target, filepath):
    load_dotenv()

    file = await open_file(filepath)
    data = file.read()
    
    cid = get_cid(data)
    sig = sign(cid=cid)

    listen_addr = multiaddr.Multiaddr(f"/ip4/0.0.0.0/tcp/{port}")

    host = config_host()

    async with host.run(listen_addrs=[listen_addr]):
        print(f"Host peer id: {host.get_id().to_string()}")

        cid = get_cid(data)
        sig = sign(cid=cid)

        msg = create_collector_msg(cid=cid,sig=sig)

        isOk = await send_data_collector_msg_to_node(host=host,target=target,msg=msg)
        
        if isOk == False:
            print("failed to send msg to dao node")
            return
        
        isOk = await transfer_file(host=host,target=target,data=data)
        if isOk == False:
            print("failed to transfer file to dao node")
            return
        
        print("successfully transferred file to dao node")


async def send_data_collector_msg_to_node(host, target, msg):
    maddr = multiaddr.Multiaddr(target)
    info = info_from_p2p_addr(maddr)

    await host.connect(info)

    stream = await host.new_stream(info.peer_id, [DATA_COLLECTOR_MSG_PROTOCOL])

    await stream.write(msg.encode())
    await stream.close()

    resp = await stream.read()

    if "MSG_OK" in resp.decode():
        return True
    
    return False

async def transfer_file(host, target, data):
    fsize = len(data)

    maddr = multiaddr.Multiaddr(target)
    info = info_from_p2p_addr(maddr)
    
    await host.connect(info)

    stream = await host.new_stream(info.peer_id, [DATA_TRANSFER_PROTOCOL])

    # write file size to stream
    size_msg = f"{fsize}\n"
    await stream.write(size_msg.encode())
    # write file to stream
    await stream.write(data)
    # close stream
    await stream.close()

    resp = await stream.read()

    if "FILE_RECEIVED" in resp.decode():
        return True

    return False

def create_collector_msg(cid, sig):
    data_collector_msg = {
        "cid": cid.decode('utf-8'),
        "sig": sig.hex(),
    }
    json_data = json.dumps(data_collector_msg)

    return json_data

def get_dao_admin_peer_info(baseUrl, address):
    url = f"{baseUrl}/peer/{address}"
    resp = requests.get(url=url)
    return resp.data.peer

def config_host():
    import secrets
    secret = secrets.token_bytes(32)
    key_pair = create_new_key_pair(secret)

    mux_opt: TMuxerOptions = {MPLEX_PROTOCOL_ID: Mplex}

    noise_key = create_ed25519_key_pair()
    sec_opt: TSecurityOptions = {
        TProtocol(noise.PROTOCOL_ID): noise.Transport(key_pair, noise_key.private_key)
    }

    host = new_host(key_pair=key_pair, muxer_opt=mux_opt, sec_opt=sec_opt)
    return host

def get_cid(data_bytes):
    sha2_256 = multihash.get("sha2-256")
    digest = sha2_256.digest(data_bytes)

    cid_res = cid.make_cid(1, 'dag-pb', digest)

    new_cid = cid_res.encode("base32")
    print("asd: ", new_cid)
    return new_cid

def sign(cid):
    rpc_provider = os.getenv("RPC_PROVIDER")
    key = os.getenv("KEY")

    w3 = Web3(Web3.WebsocketProvider(rpc_provider))

    msg_hash = messages.defunct_hash_message(cid)

    signed_msg = w3.eth.account.signHash(msg_hash, key)

    return signed_msg.signature

async def open_file(path):
    return open(path, "rb")
