package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tgetbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("\tcreateblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("\tprintchain - ")
	fmt.Println("\tsendmany -from FROM -to TO -amount AMOUNT")
	fmt.Println("\tsendone -from FROM -to TO -amount AMOUNT")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) sendToken(from, to, amount []string) {
	// 判断数据库是否存在
	if dbExists() == false {
		cli.printUsage()
		os.Exit(1)
	}
	bc := BlockchainObj()
	defer bc.DB.Close()

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
	defer bc.DB.Close()

	tx := NewUTXOTransaction(from, to, amount, bc, []*Transaction{})
	bc.MineBlock([]*Transaction{tx})
	fmt.Println("Success!")
}

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
			fmt.Printf("Transactions: %x", tranx.ID)
		}
		fmt.Println("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) createBlockchain(address string) {
	bc := CreateBlockchain(address)
	bc.DB.Close()
	fmt.Println("Done!")
}

func (cli *CLI) getBalance(address string) {
	bc := NewBlockChain(address)
	defer bc.DB.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

/*
 hello
*/
func (cli *CLI) Run() {
	cli.validateArgs()

	createblockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	genesisAddr := createblockchainCmd.String("address", "", "Create genesis block")

	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	balanceAddr := getBalanceCmd.String("address", "", "The address of the balance")

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	sendCmd := flag.NewFlagSet("sendmany", flag.ExitOnError)
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.String("amount", "", "Amount to send")

	sendOneCmd := flag.NewFlagSet("sendone", flag.ExitOnError)
	sendOneFrom := sendOneCmd.String("from", "", "Source wallet address")
	sendOneTo := sendOneCmd.String("to", "", "Destination wallet address")
	sendOneAmount := sendOneCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "createblockchain":
		err := createblockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "sendmany":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "sendone":
		err := sendOneCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if createblockchainCmd.Parsed() {
		if *genesisAddr == "" {
			createblockchainCmd.Usage()
			os.Exit(1)
		}

		cli.createBlockchain(*genesisAddr)
	}

	if getBalanceCmd.Parsed() {
		if *balanceAddr == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*balanceAddr)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount == "" {
			sendCmd.Usage()
			os.Exit(1)
		}

		from := JSONToArray(*sendFrom)
		to := JSONToArray(*sendTo)
		amount := JSONToArray(*sendAmount)

		cli.sendToken(from, to, amount)
	}

	if sendOneCmd.Parsed() {
		if *sendOneFrom == "" || *sendOneTo == "" || *sendOneAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.sendOneToken(*sendOneFrom, *sendOneTo, *sendOneAmount)
	}
}
