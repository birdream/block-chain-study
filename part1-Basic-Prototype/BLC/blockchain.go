package BLC

type Blockchain struct {
	Blocks []*Block // 存储有序区块
}

// 新增区块

func (Blockchain *Blockchain) AddBlock(data string) {
	// 创建新block
	prevBlock := Blockchain.Blocks[len(Blockchain.Blocks)-1]

	newBlock := NewBlock(data, prevBlock.Hash)

	// 将区块添加到Blocks里面
	Blockchain.Blocks = append(Blockchain.Blocks, newBlock)
}

// 创建一个带有创世区块的区块链
func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGeneisBlock()}}
}
