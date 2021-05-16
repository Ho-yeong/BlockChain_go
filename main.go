package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("send 1 BTC to Simon")
	bc.AddBlock("Send 2 more BTC to Simon")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
