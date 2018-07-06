package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) printChain() {
	// 判断数据库是否存在
	if dbExists() == false {
		cli.printUsage()
		os.Exit(1)
	}

	bc := NewBlockChain("")
	defer bc.DB.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		// pow := NewProofOfWork(*block)
		// fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		for _, tranx := range block.Transactions {
			fmt.Printf("Transactions: %x\n", tranx.ID)
			fmt.Println("\nVin:")
			for _, in := range tranx.Vin {
				fmt.Printf("%x %d %s\n", in.Txid, in.Vout, in.PubKey)
			}

			fmt.Println("\nVout:")
			for _, out := range tranx.Vout {
				fmt.Println(out.Value, "  ", out.PubKeyHash)
			}

		}
		fmt.Println("\n=========================\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
