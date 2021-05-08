package main

import (
	"fmt"
	"strconv"

	"github.com/shailp250/Golang-BlockChain/Golang-BlockChain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlockToChain("First block after Init")
	chain.AddBlockToChain("Second block after Init")
	chain.AddBlockToChain("Third block after Init")
	chain.AddBlockToChain("Fourth block after Init")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevHash)
		fmt.Printf("data in block: %s\n", block.Data)
		fmt.Printf("hash: %x \n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
