package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) DeriveHash() {
	prevInfo := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hashed := sha256.Sum256(prevInfo)
	b.Hash = hashed[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlockToChain(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func InitBlock() *Block {
	return CreateBlock("Prime", []byte{})
}

func InitBlockchain() *BlockChain {
	return &BlockChain{[]*Block{InitBlock()}}
}
