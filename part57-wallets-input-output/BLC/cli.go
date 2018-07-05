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
	fmt.Println("\tcreatewallet - Create Wallet")
	fmt.Println("\taddresslist - Get All Address")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
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

	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)

	addressListCmd := flag.NewFlagSet("addresslist", flag.ExitOnError)

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
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "addresslist":
		err := addressListCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if createblockchainCmd.Parsed() {
		if IsValidForAddress([]byte(*genesisAddr)) == false {
			fmt.Printf("\n\nWrong Address...\n\n")
			cli.printUsage()
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

		for i, fromAddr := range from {
			if IsValidForAddress([]byte(fromAddr)) == false || IsValidForAddress([]byte(to[i])) == false {
				fmt.Printf("\n\nWrong Address...\n\n")
				cli.printUsage()
				os.Exit(1)
			}
		}
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

	if createWalletCmd.Parsed() {
		cli.createWallet()
	}

	if addressListCmd.Parsed() {
		cli.addressLists()
	}
}
