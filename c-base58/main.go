package main

import (
	"fmt"
	"kyxy/block-chain/c-base58/BLC"
)

func main() {
	bytes := []byte("Norman.com")
	encode := BLC.Base58Encode(bytes)
	decode := BLC.Base58Decode(encode)

	fmt.Printf("encode: %x \n", encode)
	fmt.Printf("encode: %s \n", encode)
	fmt.Printf("decode: %x \n", decode)
	fmt.Printf("decode: %s \n", decode)
}
