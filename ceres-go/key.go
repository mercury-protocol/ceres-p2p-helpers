package ceresgo

import (
	"bufio"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
)

func GenKey() (*crypto.PrivKey, error) {
	privK := open()
	if privK != nil {
		return privK, nil
	}

	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	if err != nil {
		return nil, err
	}

	pemPriv, err := os.Create("p2ppriv.pem")
	if err != nil {
		return nil, err
	}

	defer pemPriv.Close()

	pb, err := crypto.MarshalPrivateKey(priv)
	if err != nil {
		return nil, err
	}

	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: pb,
	}
	err = pem.Encode(pemPriv, pemBlock)
	if err != nil {
		return nil, err
	}
	return &priv, nil
}

func open() *crypto.PrivKey {
	privFile, err := os.Open("p2ppriv.pem")
	if os.IsNotExist(err) {
		fmt.Println("Private key file doesn't exist")
		return nil
	} else if err != nil {
		fmt.Println("Error opening key file: ", err)
		return nil
	}
	defer privFile.Close()

	peminfo, _ := privFile.Stat()
	size := peminfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data, _ := pem.Decode(pembytes)

	privKey, err := crypto.UnmarshalPrivateKey(data.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Private key already exists, loading it")
	return &privKey
}
