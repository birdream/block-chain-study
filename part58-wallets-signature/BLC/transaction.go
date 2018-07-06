package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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

func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	if tx.IsCoinbase() {
		return
	}

	for _, vin := range tx.Vin {
		if prevTXs[hex.EncodeToString(vin.Txid)].ID == nil {
			log.Panic("ERROR: Previous transaction is not correct")
		}
	}

	txCopy := tx.TrimmedCopy()

	for inID, vin := range txCopy.Vin {
		prevTx := prevTXs[hex.EncodeToString(vin.Txid)]
		txCopy.Vin[inID].Signature = nil
		txCopy.Vin[inID].PubKey = prevTx.Vout[vin.Vout].PubKeyHash
		txCopy.ID = txCopy.Hash()
		txCopy.Vin[inID].PubKey = nil

		// 签名代码
		r, s, err := ecdsa.Sign(rand.Reader, &privKey, txCopy.ID)
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)

		tx.Vin[inID].Signature = signature
	}
}

// 拷贝一份新的Transaction用于签名
func (tx *Transaction) TrimmedCopy() Transaction {
	var ins []*TXInput
	var outs []*TXOutput

	for _, in := range tx.Vin {
		ins = append(ins, &TXInput{in.Txid, in.Vout, nil, nil})
	}

	for _, out := range tx.Vout {
		outs = append(outs, &TXOutput{out.Value, out.PubKeyHash})
	}

	return Transaction{tx.ID, ins, outs}
}

func (tx *Transaction) Hash() []byte {

	txCopy := tx

	txCopy.ID = []byte{}

	hash := sha256.Sum256(txCopy.Serialize())
	return hash[:]
}

func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}
