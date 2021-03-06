package BLC

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

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

// func (Block chain *Blockchain) AddBlock(data string) {
// 	// 创建新block
// 	prevBlock := Blockchain.Blocks[len(Blockchain.Blocks)-1]

// 	newBlock := NewBlock(data, prevBlock.Hash)

// 	// 将区块添加到Blocks里面
// 	Blockchain.Blocks = append(Blockchain.Blocks, newBlock)
// }

func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

func BlockchainObj() *Blockchain {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	var tip []byte

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b != nil {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &Blockchain{tip, db}
}

func CreateBlockchain(address string) *Blockchain {
	if dbExists() {
		fmt.Println("Blockchain already exists.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := NewCoinBaseTX(address, genesisCoinbaseData)
		genesis := NewGeneisBlock(cbtx)

		b, err := tx.CreateBucket([]byte(blocksBucket))
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

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}
	return &bc
}

// 创建一个带有创世区块的区块链
func NewBlockChain(address string) *Blockchain {
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
			cbtx := NewCoinBaseTX(address, genesisCoinbaseData)
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

// 返回当前用户未花费输出的所有交易的集合并返回交易的数组
func (bc *Blockchain) FindUnspentTransactions(address string, txs []*Transaction) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)

	for i := len(txs) - 1; i >= 0; i-- {
		tx := txs[i]
		if tx.IsCoinbase() == false {
			for _, in := range tx.Vin {
				if in.CanUnlockOutputWith(address) {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
				}
			}
		}

		txID := hex.EncodeToString(tx.ID)

	Work:
		for outIdx, out := range tx.Vout {

			if spentTXOs[txID] != nil {
				for _, spentOut := range spentTXOs[txID] {
					if spentOut == outIdx {
						continue Work
					}
				}
			}

			if out.CanBeUnlockedWith(address) {
				unspentTXs = append(unspentTXs, *tx)
			}
		}
	}

	bci := bc.Iterator()
	var hashInt big.Int

	for {
		block := bci.Next()

		for i := len(block.Transactions) - 1; i >= 0; i-- {
			tran := block.Transactions[i]
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
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}

		hashInt.SetBytes(bci.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}

	return unspentTXs
}

// 查找可用的未花费的输出信息
func (bc *Blockchain) FindSpendableOutputs(address string, amount int, txs []*Transaction) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)

	unspentTXs := bc.FindUnspentTransactions(address, txs)

	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}

// 根据交易的数组 打包新的区块
func (bc *Blockchain) MineBlock(txs []*Transaction) {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		// 新建区块
		newBlock := NewBlock(txs, bc.Tip)

		b := tx.Bucket([]byte(blocksBucket))

		if b != nil {
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			bc.Tip = newBlock.Hash
		}

		// 将区块存储到数据库
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// FindUTXO finds and returns all unspent transaction outputs
func (bc *Blockchain) FindUTXO(address string) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := bc.FindUnspentTransactions(address, []*Transaction{})

	for _, tx := range unspentTransactions {
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

// 挖掘新的多个区块
func (bc *Blockchain) MineManyBlock(from, to, amount []string) {

	var txs []*Transaction

	for i, address := range from {
		value, _ := strconv.Atoi(amount[i])
		tx := NewUTXOTransaction(address, to[i], value, bc, txs)
		txs = append(txs, tx)
	}

	var block *Block

	bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b != nil {
			hash := b.Get([]byte("l"))
			blockByte := b.Get(hash)
			block = DeserializeBlock(blockByte)

		}
		return nil
	})

	block = NewBlock(txs, block.Hash)

	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b != nil {
			err := b.Put(block.Hash, block.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), block.Hash)
			if err != nil {
				log.Panic(err)
			}

			bc.Tip = block.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
