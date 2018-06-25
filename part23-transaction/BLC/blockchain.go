package BLC

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Time 03/jan/2009 chancellor on brink of second bailout for banks"

type Blockchain struct {
	Tip []byte   //最后一个区块的hash
	DB  *bolt.DB // 数据库
}

// 新增区块

// func (Blockchain *Blockchain) AddBlock(data string) {
// 	// 创建新block
// 	prevBlock := Blockchain.Blocks[len(Blockchain.Blocks)-1]

// 	newBlock := NewBlock(data, prevBlock.Hash)

// 	// 将区块添加到Blocks里面
// 	Blockchain.Blocks = append(Blockchain.Blocks, newBlock)
// }

// 创建一个带有创世区块的区块链
func NewBlockChain() *Blockchain {
	var tip []byte //获取最后一个区块的HASH

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		// 表
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("No existing blockchain found, Creating a new one")
			cbtx := NewCoinBaseTX("Norman", genesisCoinbaseData)
			genesis := NewGeneisBlock(cbtx)

			// 创建表
			b, err = tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			// 创建创世区块
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{tip, db}
}

func (blockchain *Blockchain) AddBlock(transactions []*Transaction) {
	newBlock := NewBlock(transactions, blockchain.Tip)

	err := blockchain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		blockchain.Tip = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

// 返回当前用户未花费输出的所有交易的集合并返回交易的数组
func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)

	bci := bc.Iterator()
	var hashInt big.Int
	for {
		err := bci.DB.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blocksBucket))
			blockBytes := b.Get(bci.CurrentHash)
			block := DeserializeBlock(blockBytes)

			for _, tran := range block.Transactions {
				fmt.Printf("Transactions: %x", tran.ID)
				txID := hex.EncodeToString(tran.ID)

			Outputs:
				for outIdx, out := range tran.Vout {
					//是否已经花费
					if spentTXOs[txID] != nil {
						for _, spentOut := range spentTXOs[txID] {
							if spentOut == outIdx {
								continue Outputs
							}
						}
					}

					if out.CanBeUnlockedWith(address) {
						unspentTXs = append(unspentTXs, *tran)
					}
				}

				if tran.IsCoinbase() == false {
					for _, in := range tran.Vin {
						if in.CanUnlockOutputWith(address) {
							inTxId := hex.EncodeToString(in.Txid)
							spentTXOs[inTxId] = append(spentTXOs[inTxId], in.Vout)
						}
					}
				}
			}

			return nil
		})
		if err != nil {
			log.Panic(err)
		}

		bci = bci.Next()
		hashInt.SetBytes(bci.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	return nil
}
