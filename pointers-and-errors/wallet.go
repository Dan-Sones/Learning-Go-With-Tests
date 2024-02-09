package pointers_and_error

type Wallet struct {
	// a lowercase 'b' indicates it is private
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount

}

func (w *Wallet) Balance() int {
	return (*w).balance
}
