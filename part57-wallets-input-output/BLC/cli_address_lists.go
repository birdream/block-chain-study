package BLC

import "fmt"

func (cli *CLI) addressLists() {
	ws, _ := NewWallets()

	for addr, _ := range ws.WalletsMap {

		fmt.Println(addr)
	}
}
