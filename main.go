package main

import (
	"fmt"

	// Modules

	"github.com/C0D2D/celerycoin-core/chain"
)

func main() {
	chain := chain.InitBlockChain() // Initializing the blockchain

	// Creating the first three blocks after the Genesis block

	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
