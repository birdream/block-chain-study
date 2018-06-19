package main

import (
	"fmt"
	"kyxy/block-chain/part2-Proof-of-work/BLC"
	"time"
)

// 16 进制
// 64位哈希值
// 32字节
// 256 bit/位 （32*8）
func main() {
	// block := BLC.NewGeneisBlock()

	// fmt.Println(block.Timestamp)
	// fmt.Printf("%x\n", block.PrevBlockHash)
	// fmt.Println(block.Data)
	// fmt.Printf("%x\n", block.Hash)

	// fmt.Println(block)

	blockchain := BLC.NewBlockChain()
	// fmt.Println(blockchain)

	blockchain.AddBlock("send 20 BTC to NORMAN")
	blockchain.AddBlock("send 30 BTC to JEN")
	blockchain.AddBlock("send 100 BTC to ZACK")
	// fmt.Println(blockchain.Blocks)

	for _, block := range blockchain.Blocks {
		fmt.Printf("Data: %s \n", string(block.Data))
		fmt.Printf("prevHash: %x \n", block.PrevBlockHash)
		fmt.Printf("Timestamp: %s \n", time.Unix(block.Timestamp, 0))
		fmt.Printf("Hash: %x \n", block.Hash)
		fmt.Printf("Nonce: %d \n\n", block.Nonce)
	}
}
