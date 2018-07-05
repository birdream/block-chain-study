package BLC

import (
	"fmt"
)

func (cli *CLI) createWallet() {
	ws, _ := NewWallets()
	ws.CreateNewWallet()

	fmt.Println(len(ws.Wallets))
}
