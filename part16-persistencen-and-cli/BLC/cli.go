package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
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
		flag.Usage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		fmt.Println("data:" + *addBlockData)
	}

	if printChainCmd.Parsed() {
		fmt.Println("printChainCmdprintChainCmdprintChainCmd")
	}
}
