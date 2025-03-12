package pointerr

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin
}

// This interface is defined in the fmt package
// lets you define how your type is printed when used with the %s format string in prints.
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if (*w).balance < amount {
		// fmt.Println("Insufficient Funds")
		return ErrInsufficientFunds
	}
	(*w).balance -= amount
	return nil
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit Function is %v \n", &w.balance)
	(*w).balance += amount
}

func (w *Wallet) Balance() (balance Bitcoin) {
	// fmt.Printf("Address of balance in Balance Function is %v \n", &w.balance)
	balance = (*w).balance
	return
}
