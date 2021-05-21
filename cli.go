package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage: ")
	fmt.Println("	create Blockchain -address ADDRESS - Create a blockchain and genesis block reward to ADDRESS")
}

func (cli *CLI) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.ValidateArgs()

	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockhchainCmd := flag.NewFlagSet("createBlockchain", flag.ExitOnError)
	createBlockchainData := createBlockhchainCmd.String("address", "", "address data")

	switch os.Args[1] {

	case "createBlockchain":
		err := createBlockhchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if createBlockhchainCmd.Parsed() {
		if *createBlockchainData == "" {
			createBlockhchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) createBlockchain(address string) {
	CreateBlockchain(address)
	fmt.Println("Done !")
}
