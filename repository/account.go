package repository

type AccountRepository interface {
	Create(Account) (*Account, error)
	GetAll(int) ([]Account, error)
}

type Account struct {
	Id          int     `db:"id"`
	CustomerId  int     `db:"customer_id"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      int     `db:"status"`
}
