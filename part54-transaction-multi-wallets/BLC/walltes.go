package BLC

import (
	"fmt"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

func NewWallets() *Wallets {
	wallets := &Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	return wallets
}

func (ws *Wallets) CreateNewWallet() {
	w := NewWallet()
	fmt.Printf("Address: %s\n", w.GetAddress())

	ws.Wallets[string(w.GetAddress())] = w
}
