package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	msg := "Norman.com"

	encode := base64.StdEncoding.EncodeToString([]byte(msg))

	fmt.Printf("encode: %x \n", encode)
	fmt.Printf("encode: %s \n", encode)

	decoded, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("decode: %x \n", decoded)
	fmt.Printf("decode: %s \n", decoded)
}
