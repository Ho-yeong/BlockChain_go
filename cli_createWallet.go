package main

import "fmt"

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := NewWallets(nodeID)
	address := wallets.createWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
