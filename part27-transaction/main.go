package main

import "kyxy/block-chain/part27-transaction/BLC"

const blocksBucket = "blocks"

// 16 进制
// 64位哈希值
// 32字节
// 256 bit/位 （32*8）
func main() {
	blockchain := BLC.NewBlockChain("Norman")

	cli := BLC.CLI{blockchain}

	cli.Run()
}
