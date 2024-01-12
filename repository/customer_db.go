package repository

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "SELECT id,name,status FROM customer"
	err := r.db.Select(&customers, query)
	return customers, err
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT id,name,status FROM customer WHERE id = ?"
	err := r.db.Get(&customer, query, id)
	return &customer, err
}
