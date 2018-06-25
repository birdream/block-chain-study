package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 12 // 23个0前面

type ProofWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	fmt.Println(target)

	pow := &ProofWork{block, target}

	return pow
}

func (pow *ProofWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Start mining the new block... \n")

	for nonce < maxNonce {
		data := pow.perparData(nonce)

		hash = sha256.Sum256(data)
		// fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	// fmt.Print("\n\n")
	return nonce, hash[:]
}

// 数据拼接，反回字节数组
func (pow *ProofWork) perparData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}
