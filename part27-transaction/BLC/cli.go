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
	fmt.Println("    send -from FROM -to TO -amount AMOUNT")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) sendToken(from, to string, amount int) {
	tx := NewUTXOTransaction(from, to, amount, cli.BC)

	cli.BC.MineBlock([]*Transaction{tx})
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
		fmt.Printf("Nonce: %d \n", block.Nonce)
		for _, tranx := range block.Transactions {
			fmt.Printf("Transactions: %x\n", tranx.ID)
		}
		fmt.Printf("\n")

		hashInt.SetBytes(blockchanIterator.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
}

func (cli *CLI) addBlock(data string) {
	// fmt.Print("\n------Norman-----\n")
	// fmt.Println(cli.BC.FindUnspentTransactions("Norman"))
	// fmt.Print("\n------Jan-----\n")
	// fmt.Println(cli.BC.FindUnspentTransactions("Jan"))
	// fmt.Print("\n------Lu-----\n")
	// fmt.Println(cli.BC.FindUnspentTransactions("Lu"))
	// fmt.Print("\n-----------\n")
	count, outputMap := cli.BC.FindSpendableOutputs("Norman", 5)

	fmt.Println(count)
	fmt.Println(outputMap)
	// cli.sendToken()
}

/*
 hello
*/
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block Data")

	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")

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
	case "send":
		err := sendCmd.Parse(os.Args[2:])
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

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.sendToken(*sendFrom, *sendTo, *sendAmount)
	}
}
