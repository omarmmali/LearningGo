package wallet

import "errors"

type Wallet struct{
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount int) error {
	if w.balance < amount {
        return errors.New("cannot withdraw: insufficient funds")
    }

	w.balance -= amount
	
	return nil
}