package pointers_and_error

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	// a lowercase 'b' indicates it is private
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount

}

func (w *Wallet) Balance() Bitcoin {
	return (*w).balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
