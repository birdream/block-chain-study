package main

import (
	"fmt"
	"kyxy/block-chain/part3-Serialize-DeserializeBlock/BLC"
)

func main() {
	block := BLC.Block{[]byte("send 3 BTC to Norman"), 1000}

	fmt.Printf("%s\n", block.Data)
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("\n\n")

	bytes := block.Serialize()
	fmt.Println(bytes)
	fmt.Printf("\n\n")

	blc := BLC.DeserializeBlock(bytes)
	fmt.Printf("%s\n", blc.Data)
	fmt.Printf("%d\n", blc.Nonce)
	fmt.Printf("\n\n")
}
