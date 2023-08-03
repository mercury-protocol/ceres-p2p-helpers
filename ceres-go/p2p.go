package ceresgo

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	chunker "github.com/ipfs/boxo/chunker"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
	mh "github.com/multiformats/go-multihash"
)

const (
	DataTransferProtocol     = "/mcy/transfer/dao/data/0.0.1"
	DataCollectorMsgProtocol = "/mcy/msg/dao/0.0.1"
)

var PendingTransfers map[string]PendingTransfer

func Init(port int) (*host.Host, error) {
	priv, err := GenKey()
	if err != nil {
		return nil, err
	}
	opts := []libp2p.Option{
		libp2p.Identity(*priv),
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)),

		libp2p.DefaultTransports,
		libp2p.DisableRelay(),
	}

	h, err := libp2p.New(opts...)
	if err != nil {
		return nil, err
	}

	PendingTransfers = make(map[string]PendingTransfer)

	return &h, nil
}

func TransferFile(n host.Host, filename string, target ma.Multiaddr) error {
	n.SetStreamHandler(DataTransferProtocol, func(s network.Stream) {
		if err := handleFileSendResp(s); err != nil {
			s.Reset()
		} else {
			s.Close()
		}
	})

	info, err := peer.AddrInfoFromP2pAddr(target)
	if err != nil {
		return err
	}

	AddVerifiedPeer(n, info)

	s, err := n.NewStream(context.Background(), info.ID, DataTransferProtocol)

	if err != nil {
		return err
	}

	fmt.Println("Opening stream to send file")

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fstat, err := file.Stat()
	if err != nil {
		return err
	}

	fileLenStr := strconv.FormatInt(fstat.Size(), 10)

	_, err = s.Write([]byte(fmt.Sprintf("%s\n", fileLenStr)))
	if err != nil {
		return err
	}

	_, err = io.Copy(s, file)
	if err != nil {
		return err
	}

	buf := bufio.NewReader(s)
	out, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	if !strings.Contains(out, "FILE_RECEIVED") {
		return errors.New("data dao admin node did not confirm data receipt")
	}

	fmt.Println("Data transfer successful")
	return nil
}

func SendDataCollectorMsgToNode(n host.Host, target ma.Multiaddr, msg DataCollectorMsg) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	n.SetStreamHandler(DataCollectorMsgProtocol, func(s network.Stream) {
		s.Close()
	})

	info, err := peer.AddrInfoFromP2pAddr(target)
	if err != nil {
		return err
	}
	AddVerifiedPeer(n, info)

	s, err := n.NewStream(context.Background(), info.ID, DataCollectorMsgProtocol)

	if err != nil {
		return err
	}

	_, err = s.Write(data)
	s.CloseWrite()
	if err != nil {
		return err
	}

	buf := bufio.NewReader(s)
	out, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	if !strings.Contains(out, "MSG_OK") {
		return errors.New("data dao admin did not confirm message")
	}

	return nil
}

func Cid(data []byte) (string, error) {
	// Create an IPLD UnixFS chunker with size 1 MiB
	chunks := chunker.NewSizeSplitter(bytes.NewReader(data), 1024*1024)

	// Concatenate the chunks to build the DAG
	var buf bytes.Buffer
	for {
		chunk, err := chunks.NextBytes()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		buf.Write(chunk)
	}

	// Calculate the CID for the DAG
	hash, err := mh.Sum(buf.Bytes(), mh.SHA2_256, -1)
	if err != nil {
		return "", err
	}

	// Create a CID version 1 (with multibase encoding base58btc)
	c := cid.NewCidV1(cid.DagProtobuf, hash)

	// Print the CID as a string
	fmt.Println("CID:", c.String())

	return c.String(), nil
}

func CreateCollectorMsg(cid string, k string) (DataCollectorMsg, error) {
	priv, err := crypto.HexToECDSA(k)
	if err != nil {
		return DataCollectorMsg{}, err
	}

	hash := crypto.Keccak256Hash([]byte(cid))

	sig, err := crypto.Sign(hash.Bytes(), priv)
	if err != nil {
		return DataCollectorMsg{}, err
	}

	msg := DataCollectorMsg{
		Cid:       cid,
		Signature: sig,
	}

	return msg, nil
}

func GetDaoAdminPeerInfo(daoAdminAddr string, db string) (PeerInfo, error) {
	var dbUrl = fmt.Sprintf("%s/api/v1/peer", db)
	url := fmt.Sprintf("%s/%s", dbUrl, daoAdminAddr)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return PeerInfo{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return PeerInfo{}, err
	}
	defer res.Body.Close()
	var target map[string]PeerInfo

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return PeerInfo{}, err
	}

	err = json.Unmarshal(resBody, &target)
	if err != nil {
		return PeerInfo{}, err
	}

	return target["peer"], nil
}

func VerifyPeerInfo(p PeerInfo, clientUrl string) (bool, error) {
	// recover signer address from signature for multiaddress
	hash1 := crypto.Keccak256Hash([]byte(p.TCP))
	sig1, err := hexutil.Decode(p.Sig1)
	if err != nil {
		return false, err
	}

	tcpOk, err := verifyPeerInfoSignature(common.HexToAddress(p.Address), hash1, sig1, clientUrl)
	if err != nil {
		return false, err
	}

	hash2 := crypto.Keccak256Hash([]byte(p.Webtransport))
	sig2, err := hexutil.Decode(p.Sig2)
	if err != nil {
		return false, err
	}

	webtOk, err := verifyPeerInfoSignature(common.HexToAddress(p.Address), hash2, sig2, clientUrl)
	if err != nil {
		return false, err
	}

	return tcpOk && webtOk, nil
}

func IsAuthorizedPeer(n *host.Host, peerId peer.ID) bool {
	peers := GetVerifiedPeers(*n)
	for i := 0; i < len(peers); i++ {
		if peerId == peers[i] {
			return true
		}
	}
	return false
}

func AddVerifiedPeer(n host.Host, info *peer.AddrInfo) {
	n.Peerstore().AddAddr(info.ID, info.Addrs[0], peerstore.PermanentAddrTTL)
}

func GetVerifiedPeers(n host.Host) peer.IDSlice {
	return n.Peerstore().Peers()
}

func handleFileSendResp(s network.Stream) error {
	buf := bufio.NewReader(s)
	_, err := buf.ReadString('\n')
	if err != nil {
		return err
	}
	_, err = s.Write([]byte("ok"))
	return err
}

func verifyPeerInfoSignature(addr common.Address, hash common.Hash, sig []byte, clientUrl string) (bool, error) {
	pk, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false, err
	}
	signerAddr := crypto.PubkeyToAddress(*pk)

	var adminRole [32]byte

	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		return false, err
	}

	dao, err := NewDataDao(addr, client)
	if err != nil {
		return false, err
	}

	hasRole, err := dao.HasRole(nil, adminRole, signerAddr)
	if err != nil {
		return false, err
	}

	return hasRole, nil
}
