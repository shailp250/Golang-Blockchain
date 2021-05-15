package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const(
	dbPath = "./tmp/blocks"
)

type BlockChain struct {
	LastHash []*Block
	Database *badger.DB
}

func InitBlockchain() *BlockChain {
	var lastHash []byte
	opts := badger.DefaultOptions
	opts.Dir := dbPath
	opts.ValueDir := dbPath

	db, err := badger.Open(opts)
	Handle(err)
	//initalizing a blockchain, passes a closer which takes in a pointer to a badger transaction
	err := db.Update(func(txn *badger.Txn) error {
		// 4 main pieces, if blockchain has been stored - if yes then create new BC in memory and get last hash in the disc database and push in here
		//if not last hash(lh) found than genesis blockchain = lh
		if _, err := txn.Get([]byte("lh")); err = badger.ErrKeyNotFound {
			fmt.Println("No BC found")
			genesis := InitBlock()
			fmt.Println("Genesis Proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			// set initial block to "lh"
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash
			return err
		} else {
			item, err = txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.Value()
			return err
		}
	})
	Handle(err)

	blockchain := BlockChain{lastHash, db}
	return &blockchain
//	return &BlockChain{[]*Block{InitBlock()}}
}

func (chain *BlockChain) AddBlock(data string) {
	// establish a reah-only type transaction
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error{
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()

		return err
	})
	Handle(err)

	newBlock := CreateBlock(data, lastHash)
	//now a read-write type trasanction to put new block in database and assign new hash to lh
	err = chain.Database.Update(func(txn *badger.Txn) error{
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash
		return err
	})


 /* prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new) */
}
