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
	if amount > w.balance {
		return fmt.Errorf("Debit amount can not be higher than balance!")
	}

	w.balance -= amount
	return nil
}
