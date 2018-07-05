package BLC

import "fmt"

func (cli *CLI) createBlockchain(address string) {
	bc := CreateBlockchain(address)
	bc.DB.Close()
	fmt.Println("Done!")
}
