package account

// Account  struct
type Account struct {
	Owner   string //public : 대문자, private : 소문로
	Balance int
}

// *로 복사본을 전달해서 속도를 빠르게 함
// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{Owner: owner, Balance: 0}
	return &account
}