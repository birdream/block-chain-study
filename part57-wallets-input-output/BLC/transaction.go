package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

type Transaction struct {
	ID   []byte
	Vin  []*TXInput
	Vout []*TXOutput
}

//
func (tx *Transaction) SetID() {

	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]

}

// 创建一新的coinbase交易
func NewCoinBaseTX(to string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("reward to '%s'", to)
	}

	// 创建特殊的输入
	txin := &TXInput{[]byte{}, -1, nil, []byte{}}
	// 创建输出
	txout := NewTXOutput(subsidy, to)

	// 创建交易
	tx := &Transaction{[]byte{}, []*TXInput{txin}, []*TXOutput{txout}}

	tx.SetID()

	return tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && tx.Vin[0].Vout == -1 && len(tx.Vin[0].Txid) == 0
}
