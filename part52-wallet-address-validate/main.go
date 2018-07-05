package main

import (
	"fmt"
	"kyxy/block-chain/part52-wallet-address-validate/BLC"
)

func main() {
	wallet := BLC.NewWallet()

	address := wallet.GetAddress()

	isValid := wallet.IsValidForAddress(address)

	fmt.Printf("%s validation: %v", address, isValid)
}
