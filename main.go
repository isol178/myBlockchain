package main

import (
	"fmt"
	"log"
	"myBlockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	fmt.Println("start")
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())
}
