package main

import (
	"fmt"
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	ks "github.com/ChainSafe/chainbridge-utils/keystore"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) != 3 {
		fmt.Printf("Usage %s [keystore-directory] [address]\n", os.Args[0])
		os.Exit(1)
	}

	path := os.Args[1]
	addr := os.Args[2]

	extractPKToFile(addr, path)

	fmt.Printf("Private key for %s written to ./privatekey\n", addr)
}

func extractPKToFile(addr, path string) {
	kpI, err := ks.KeypairFromAddress(addr, "ethereum", path, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	kp, _ := kpI.(*secp256k1.Keypair)
	rawPK := kp.Encode()
	pk := fmt.Sprintf("%x\n", rawPK)

	err = ioutil.WriteFile("privatekey", []byte(pk), 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
