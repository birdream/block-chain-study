package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) sendToken(from, to, amount []string) {
	// 判断数据库是否存在
	if dbExists() == false {
		cli.printUsage()
		os.Exit(1)
	}
	bc := BlockchainObj()
	// defer bc.DB.Close()

	// tx := NewUTXOTransaction(from, to, amount, bc)
	// bc.MineBlock([]*Transaction{tx})
	bc.MineManyBlock(from, to, amount)
	defer bc.DB.Close()

	fmt.Println("Success!")
}

func (cli *CLI) sendOneToken(from, to string, amount int) {
	// 判断数据库是否存在
	if dbExists() == false {
		cli.printUsage()
		os.Exit(1)
	}

	bc := NewBlockChain(from)
	// defer bc.DB.Close()

	tx := NewUTXOTransaction(from, to, amount, bc, []*Transaction{})
	bc.MineBlock([]*Transaction{tx})
	fmt.Println("Success!")
}
