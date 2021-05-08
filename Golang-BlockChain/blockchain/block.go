package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	Blocks []*Block
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
