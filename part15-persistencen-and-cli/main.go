package main

import (
	"fmt"
	"kyxy/block-chain/part15-persistencen-and-cli/BLC"
	"log"
	"math/big"
	"time"

	"github.com/boltdb/bolt"
)

const blocksBucket = "blocks"

// 16 进制
// 64位哈希值
// 32字节
// 256 bit/位 （32*8）
func main() {
	blockchain := BLC.NewBlockChain()

	// fmt.Println(blockchain)
	// fmt.Printf("tip: %x\n", blockchain.Tip)
	// fmt.Println(block.Data)
	// fmt.Printf("%x\n", block.Hash)

	// fmt.Println(block)
	// fmt.Println(blockchain)

	blockchain.AddBlock("send 20 BTC to NORMAN")
	blockchain.AddBlock("send 30 BTC to JEN")
	blockchain.AddBlock("send 100 BTC to ZACK")
	// fmt.Printf("tip: %x\n", blockchain.Tip)

	var blockchanIterator *BLC.BlockchainIterator
	blockchanIterator = blockchain.Iterator()

	var hashInt big.Int

	for {
		fmt.Printf("%x\n", blockchanIterator.CurrentHash)

		err := blockchanIterator.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blocksBucket))
			blockBytes := b.Get(blockchanIterator.CurrentHash)
			block := BLC.DeserializeBlock(blockBytes)

			fmt.Printf("Data: %s \n", string(block.Data))
			fmt.Printf("prevHash: %x \n", block.PrevBlockHash)
			fmt.Printf("Timestamp: %s \n", time.Unix(block.Timestamp, 0))
			fmt.Printf("Hash: %x \n", block.Hash)
			fmt.Printf("Nonce: %d \n\n", block.Nonce)

			return nil
		})
		if err != nil {
			log.Panic(err)
		}

		blockchanIterator = blockchanIterator.Next()
		hashInt.SetBytes(blockchanIterator.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
}
