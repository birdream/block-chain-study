package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	CurrentHash []byte   // 当前正在遍历区块HASH
	DB          *bolt.DB // 数据库
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.Tip, bc.DB}
}

func (bi *BlockchainIterator) Next() *Block {
	var block *Block

	err := bi.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		currentBlockBytes := b.Get(bi.CurrentHash)
		block = DeserializeBlock(currentBlockBytes)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bi.CurrentHash = block.PrevBlockHash

	return block
}
