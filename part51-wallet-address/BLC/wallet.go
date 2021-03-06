package BLC

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const (
	version            = byte(0x00)
	addressCjecksumLen = 4
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()

	// fmt.Println(&privateKey)
	// fmt.Println(publicKey)

	return &Wallet{privateKey, publicKey}
}

// build publicKey through privateKey
func newKeyPair() (ecdsa.PrivateKey, []byte) {

	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}

func (w *Wallet) GetAddress() []byte {
	// 1 hash160 -》》 长度是 20
	ripemd160Has := w.Ripemd160Has(w.PublicKey)

	version_ripemd160Has := append([]byte{version}, ripemd160Has...)

	checkSumBytes := w.CheckSum(version_ripemd160Has)

	bytes := append(version_ripemd160Has, checkSumBytes...)

	return Base58Encode(bytes)
}

func (w *Wallet) Ripemd160Has(pubKey []byte) []byte {
	// 256
	hash256 := sha256.New()
	hash256.Write(pubKey)
	hash := hash256.Sum(nil)

	// 160
	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(hash)

	return ripemd160Hash.Sum(nil)
}

func (w *Wallet) CheckSum(b []byte) []byte {
	// 两次hash256
	h_1 := sha256.Sum256(b)
	h_2 := sha256.Sum256(h_1[:])

	return h_2[:addressCjecksumLen]

}
