package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	CurrentHash []byte   // 当前正在遍历区块HASH
	DB          *bolt.DB // 数据库
}

func (blockchain *Blockchain) Iterator() *BlockchainIterator {

	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}

func (bi *BlockchainIterator) Next() *BlockchainIterator {
	var nextHash []byte

	err := bi.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		currentBlockBytes := b.Get(bi.CurrentHash)

		currentBlock := DeserializeBlock(currentBlockBytes)
		nextHash = currentBlock.PrevBlockHash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &BlockchainIterator{nextHash, bi.DB}
}
