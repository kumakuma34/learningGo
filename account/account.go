package account

import (
	"errors"
	"fmt"
)

// Account  struct
type Account struct {
	owner   string //public : 대문자, private : 소문로
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount *로 복사본을 전달해서 속도를 빠르게 함
// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit method 선언
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

// Withdrawl from your account
func (a *Account) Withdrawl(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\n Has: ", a.Balance())
}
