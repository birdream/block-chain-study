package main

import "kyxy/block-chain/part31-transaction-multi/BLC"

// 16 进制
// 64位哈希值
// 32字节
// 256 bit/位 （32*8）
func main() {
	cli := BLC.CLI{}
	cli.Run()
}
