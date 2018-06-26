package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const subsidy = 10

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type TXInput struct {
	Txid      []byte //交易的ID
	Vout      int    //储存在TXOutput在Vout里面的索引
	ScriptSig string //用户名
}

type TXOutput struct {
	Value        int    //分
	ScriptPubKey string // 用户名
}

//
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// 创建一新的coinbase交易
func NewCoinBaseTX(to string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("reward to '%s'", to)
	}

	// 创建特殊的输入
	txin := TXInput{[]byte{}, -1, data}
	// 创建输出
	txout := TXOutput{subsidy, to}
	// 创建交易
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && tx.Vin[0].Vout == -1 && len(tx.Vin[0].Txid) == 0
}

// 检查账号地址
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

// 检查是否能解锁账号
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

// 建立转账交易
func NewUTXOTransaction(from string, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("\nNot enough funds ...!!!\n")
	}

	// 建立输入
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// 建立输出， 转账
	output := TXOutput{amount, to}
	outputs = append(outputs, output)
	// 建立输出， 找零
	output = TXOutput{acc - amount, from}
	outputs = append(outputs, output)

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
