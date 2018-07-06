package BLC

import (
	"bytes"
)

type TXOutput struct {
	Value      int64  //分
	PubKeyHash []byte //公钥 160hash
}

func GetPubKeyFromAddress(address string) []byte {
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]

	return pubKeyHash
}

// 检查是否能解锁账号
func (out *TXOutput) CanBeUnlockedWith(address string) bool {
	pubKeyHash := GetPubKeyFromAddress(address)

	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	out.PubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
}

// func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
// 	return bytes.Compare(out.pubKeyHash, pubKeyHash) == 0
// }

func NewTXOutput(value int64, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}
