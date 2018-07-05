package main

import (
	"fmt"
	"kyxy/block-chain/part53-wallets/BLC"
)

func main() {
	wallets := BLC.NewWallets()

	wallets.CreateNewWallet()
	wallets.CreateNewWallet()

	fmt.Println(wallets.Wallets)

}
