package main

import (
	"fmt"
	"kyxy/block-chain/part13-persistencen-and-cli/BLC"
	"math/big"
)

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

		hashInt.SetBytes(blockchanIterator.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		} else {
			blockchanIterator = blockchanIterator.Next()
		}
	}
}
