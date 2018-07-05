package main

import (
	"fmt"
	"kyxy/block-chain/part51-wallet-address/BLC"
)

func main() {
	wallet := BLC.NewWallet()

	address := wallet.GetAddress()

	fmt.Printf("Address: %s\n", address)
}
