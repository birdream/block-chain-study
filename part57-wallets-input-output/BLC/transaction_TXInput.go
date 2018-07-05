package BLC

import (
	"bytes"
)

type TXInput struct {
	Txid      []byte //交易的ID
	Vout      int    //储存在TXOutput在Vout里面的索引
	Signature []byte //数字签名
	PubKey    []byte //公钥
}

// 检查账号地址
func (in *TXInput) UnLockRipemd160Hash(ripemd160Hash []byte) bool {
	lockingHash := Ripemd160Has(in.PubKey)

	return bytes.Compare(lockingHash, ripemd160Hash) == 0
}
