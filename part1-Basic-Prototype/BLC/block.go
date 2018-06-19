package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	// timestamp creating block time
	Timestamp int64
	// last block hash
	PrevBlockHash []byte
	// data transtaction data
	Data []byte
	// current block hash
	Hash []byte
}

func (b *Block) SetHash() {
	// 时间戳转化为字节数组
	//  将int64转字符串
	//  将字符串转字节数组
	timeString := []byte(strconv.FormatInt(b.Timestamp, 10))
	// 将除了hash以外的其他属性以字节数组的形式全拼接起来
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timeString}, []byte{})
	// 将拼接起来的数据进行256hash
	hash := sha256.Sum256(headers)
	// 将hash赋值
	b.Hash = hash[:]
}

// 工厂方法
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Hash:          []byte{}}

	block.SetHash()

	return block
}

func NewGeneisBlock() *Block {
	return NewBlock("Genenis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
