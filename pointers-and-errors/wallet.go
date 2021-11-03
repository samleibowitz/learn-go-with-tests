package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(sum int) int {
	w.balance += sum
	return w.balance
}

func (w *Wallet) Balance() int {
	return w.balance
}
