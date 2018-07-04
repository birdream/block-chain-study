package main

import (
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

func main() {
	//160位 40个数字 20个字节
	hasher := ripemd160.New()
	hasher.Write([]byte("Norman.com"))
	bytes := hasher.Sum(nil)

	fmt.Printf("%x\n", bytes)
}
