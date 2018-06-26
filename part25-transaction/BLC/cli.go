package BLC

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

type CLI struct {
	BC *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("    addBlock -data BLOCK_DATA")
	fmt.Println("    printchain - ")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) printChain() {
	var blockchanIterator *BlockchainIterator
	blockchanIterator = cli.BC.Iterator()

	var hashInt big.Int

	for {
		// fmt.Printf("%x\n", blockchanIterator.CurrentHash)
		block := blockchanIterator.Next()

		fmt.Printf("prevHash: %x \n", block.PrevBlockHash)
		fmt.Printf("Timestamp: %s \n", time.Unix(block.Timestamp, 0))
		fmt.Printf("Hash: %x \n", block.Hash)
		fmt.Printf("Transactions: %x \n", block.Transactions)
		fmt.Printf("Nonce: %d \n\n", block.Nonce)
		for _, tranx := range block.Transactions {
			fmt.Printf("Transactions: %x\n\n", tranx.ID)
		}

		hashInt.SetBytes(blockchanIterator.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
}

func (cli *CLI) addBlock(data string) {
	// cli.BC.AddBlock([]*Transaction{})
	count, outputMap := cli.BC.FindSpendableOutputs("Norman", 5)

	fmt.Println(count)
	fmt.Println(outputMap)
}

/*
 hello
*/
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block Data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		// fmt.Println("data:" + )
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
