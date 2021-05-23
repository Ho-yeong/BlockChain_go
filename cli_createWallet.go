package main

import "fmt"

func (cli *CLI) createWallet() {
	wallets, _ := NewWallets()
	address := wallets.createWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
