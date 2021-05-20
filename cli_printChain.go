package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("========= Block %x ===========", block.Hash)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)

		pow := NewProofOfWork(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
