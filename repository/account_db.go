package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) GetAll(customerId int) ([]Account, error) {
	accounts := []Account{}
	query := "SELECT id,customer_id,account_type,amount,status FROM account where customer_id = ?"
	err := r.db.Select(&accounts, query, customerId)
	return accounts, err
}

func (r accountRepositoryDB) Create(account Account) (*Account, error) {
	query := "INSERT account (customer_id,account_type,amount,status) values (?,?,?,?)"
	results, err := r.db.Exec(
		query,
		account.CustomerId,
		account.AccountType,
		account.Amount,
		account.Status,
	)

	if err != nil {
		return nil, err
	}

	id, err := results.LastInsertId()

	if err != nil {
		return nil, err
	}
	account.Id = int(id)

	return &account, nil
}
