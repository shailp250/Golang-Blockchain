package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

/* Already in proof-of-work Algo in "proof.go"
func (b *Block) DeriveHash() {
	prevInfo := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hashed := sha256.Sum256(prevInfo)
	b.Hash = hashed[:]
}
*/
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	//block.DeriveHash()
	return block
}

func InitBlock() *Block {
	return CreateBlock("Prime", []byte{})
}

// badger db only accepts arrays of bytes so you need to serialize and deserialize
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	Handle(err)
	return &block

}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
