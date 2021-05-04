package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PrevHash []byte
	Hash     []byte
	Data     []byte
}

func (b *Block) DeriveHash() {
	prevInfo := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hashed := sha256.Sum256(prevInfo)
	b.Hash = hashed
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data, prevHash)}
	block.DeriveHash()
	return block
}

func main() {
	fmt.Println("hello world")
}
