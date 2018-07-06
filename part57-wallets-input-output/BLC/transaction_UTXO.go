package BLC

import (
	"encoding/hex"
	"log"
	"math/big"
)

// 建立转账交易
func NewUTXOTransaction(from string, to string, amount int, bc *Blockchain, txs []*Transaction) *Transaction {
	wallets, _ := NewWallets()
	wallet := wallets.WalletsMap[from]

	var inputs []*TXInput
	var outputs []*TXOutput

	money, spendableUTXODic := bc.FindSpendableOutputs(from, amount, txs)

	if money < amount {
		log.Panic("\nNot enough funds ...!!!\n")
	}

	// 建立输入
	for txid, outs := range spendableUTXODic {
		txID, _ := hex.DecodeString(txid)
		for _, out := range outs {
			input := &TXInput{txID, out, nil, wallet.PublicKey}
			inputs = append(inputs, input)
		}
	}

	// 建立输出， 转账
	// output := TXOutput{amount, to}
	output := NewTXOutput(int64(amount), to)
	outputs = append(outputs, output)
	// 建立输出， 找零
	if int64(money) > int64(amount) {
		output = NewTXOutput(int64(money)-int64(amount), from)
		outputs = append(outputs, output)
	}

	tx := &Transaction{[]byte{}, inputs, outputs}

	tx.SetID()

	return tx
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
				accumulated += int(out.Value)
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}

// 返回当前用户未花费输出的所有交易的集合并返回交易的数组
func (bc *Blockchain) FindUnspentTransactions(address string, txs []*Transaction) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)

	for i := len(txs) - 1; i >= 0; i-- {
		tx := txs[i]
		if tx.IsCoinbase() == false {
			for _, in := range tx.Vin {
				pubKeyHash := GetPubKeyFromAddress(address)

				if in.UnLockRipemd160Hash(pubKeyHash) {
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
					pubKeyHash := GetPubKeyFromAddress(address)

					if in.UnLockRipemd160Hash(pubKeyHash) {
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
