package main

type Wallet struct {
	id      int
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Credit(amount int) error {
	return nil
}

func (w *Wallet) Debit(amount int) error {
	return nil
}
