package BLC

type TXInput struct {
	Txid      []byte //交易的ID
	Vout      int    //储存在TXOutput在Vout里面的索引
	ScriptSig string //用户名
}

// 检查账号地址
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}
