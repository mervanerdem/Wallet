package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID             int
	Wallet_balance decimal.Decimal
}

type WStorage interface {
	Update(wallet Wallet) error
	Get(id int) (*Wallet, error)
}

func (w *Wallet) Balance() decimal.Decimal {
	w.Wallet_balance = decimal.NewFromFloat(w.Wallet_balance.InexactFloat64())
	return w.Wallet_balance
}

func (w *Wallet) Debit(amount decimal.Decimal) error {
	if amount.IsNegative() {
		return fmt.Errorf("Debit amount can not be negative!")
	}
	w.Wallet_balance = decimal.Sum(w.Wallet_balance, amount)
	return nil
}

func (w *Wallet) Credit(amount decimal.Decimal) error {

	if amount.GreaterThan(w.Wallet_balance) == true {
		return fmt.Errorf("Credit amount can not be higher than balance!")
	} else if amount.IsNegative() == true {
		return fmt.Errorf("Credit amount can not be negative!")
	}
	amount = amount.Neg()
	w.Wallet_balance = decimal.Sum(w.Wallet_balance, amount)
	return nil
}
