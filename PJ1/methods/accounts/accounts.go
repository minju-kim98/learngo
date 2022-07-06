package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *account {
	newBankAccount := account{owner: owner, balance: 0}
	return &newBankAccount
}

// Deposit x amount on your account
func (a *account) Deposit(amount int) {
	a.balance += amount
}

var errNoMoney = errors.New("Can't withdraw. You are poor:(")

// Withdraw from your account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change owner of the account
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Get balance of your account
func (a account) GetBalance() int {
	return a.balance
}

// Get owner of your account
func (a account) GetOwner() string {
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint(a.owner, "'s account.\nBalance: ", a.balance)
}
