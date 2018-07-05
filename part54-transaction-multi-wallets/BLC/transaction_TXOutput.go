package BLC

type TXOutput struct {
	Value        int    //分
	ScriptPubKey string // 用户名
}

// 检查是否能解锁账号
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}
