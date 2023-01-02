package account

import "errors"

// Account  struct
type Account struct {
	owner   string //public : 대문자, private : 소문로
	balance int
}

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
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	return nil
}
