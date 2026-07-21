package account

type Account struct {
	owner string
	balance float64
}

func New(owner string, initialBalance float64) Account {
	return Account{
		owner: owner,
		balance: initialBalance,
	}	
}

func (account Account) Owner() string {
	return account.owner
}
func (account Account) Balance() float64 {
	return account.balance
}
func (account *Account) Deposit(amount float64) {
	account.balance += amount
}