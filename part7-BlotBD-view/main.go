package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

func main() {
	// 数据库创建
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// 查询
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		valueByte := b.Get([]byte("Norman"))
		fmt.Printf("%s", valueByte)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}
