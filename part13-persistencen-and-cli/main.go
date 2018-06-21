package main

import (
	"fmt"
	"kyxy/block-chain/part12-persistencen-and-cli/BLC"
)

// 16 进制
// 64位哈希值
// 32字节
// 256 bit/位 （32*8）
func main() {
	blockchain := BLC.NewBlockChain()

	// fmt.Println(blockchain)
	fmt.Printf("tip: %x\n", blockchain.Tip)
	// fmt.Println(block.Data)
	// fmt.Printf("%x\n", block.Hash)

	// fmt.Println(block)
	// fmt.Println(blockchain)

	blockchain.AddBlock("send 20 BTC to NORMAN")
	blockchain.AddBlock("send 30 BTC to JEN")
	blockchain.AddBlock("send 100 BTC to ZACK")
	fmt.Printf("tip: %x\n", blockchain.Tip)

	// for _, block := range blockchain.Blocks {
	// 	fmt.Printf("Data: %s \n", string(block.Data))
	// 	fmt.Printf("prevHash: %x \n", block.PrevBlockHash)
	// 	fmt.Printf("Timestamp: %s \n", time.Unix(block.Timestamp, 0))
	// 	fmt.Printf("Hash: %x \n", block.Hash)
	// 	fmt.Printf("Nonce: %d \n\n", block.Nonce)
	// }
}
