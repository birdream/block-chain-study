package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

type Block struct {
	// timestamp creating block time
	Timestamp int64
	// last block hash
	PrevBlockHash []byte
	// data transtaction data
	Transactions []*Transaction
	// current block hash
	Hash []byte
	// Nonce 满足某个难度的随机数
	Nonce int
}

// 提供给挖矿使用
// 将区块里面所有的id拼接并且生成hash
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func (b *Block) SetHash() {
	// 时间戳转化为字节数组
	//  将int64转字符串
	//  将字符串转字节数组
	timeString := []byte(strconv.FormatInt(b.Timestamp, 10))
	// 将除了hash以外的其他属性以字节数组的形式全拼接起来
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.HashTransactions(), timeString}, []byte{})
	// 将拼接起来的数据进行256hash
	hash := sha256.Sum256(headers)
	// 将hash赋值
	b.Hash = hash[:]
}

// 工厂方法
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		prevBlockHash,
		transactions,
		[]byte{},
		0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	// pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGeneisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
