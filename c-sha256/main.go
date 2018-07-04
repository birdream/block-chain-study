package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	//256位 32字节
	hasher := sha256.New()
	hasher.Write([]byte("Norman.com"))
	bytes := hasher.Sum(nil)

	fmt.Printf("%x\n", bytes)
}
