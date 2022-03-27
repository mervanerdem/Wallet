package main

import "fmt"

type Wallet struct {
	id      int
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Credit(amount int) error {
	if amount < 0 {
		return fmt.Errorf("Credit amount can not be negative!")
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Debit(amount int) error {
	w.balance -= amount
	return nil
}
