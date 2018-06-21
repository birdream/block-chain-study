package BLC

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type Blockchain struct {
	Tip []byte   //最后一个区块的hash
	DB  *bolt.DB // 数据库
}

type BlockchainIterator struct {
	currentHash []byte
}

// 新增区块

// func (Blockchain *Blockchain) AddBlock(data string) {
// 	// 创建新block
// 	prevBlock := Blockchain.Blocks[len(Blockchain.Blocks)-1]

// 	newBlock := NewBlock(data, prevBlock.Hash)

// 	// 将区块添加到Blocks里面
// 	Blockchain.Blocks = append(Blockchain.Blocks, newBlock)
// }

// 创建一个带有创世区块的区块链
func NewBlockChain() *Blockchain {
	var tip []byte //获取最后一个区块的HASH

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		// 表
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("No existing blockchain found, Creating a new one")
			genesis := NewGeneisBlock()

			// 创建表
			b, err = tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			// 创建创世区块
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{tip, db}
}

func (blockchain *Blockchain) AddBlock(data string) {
	newBlock := NewBlock(data, blockchain.Tip)

	err := blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		blockchain.Tip = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}
