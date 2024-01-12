package repository

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}

type Customer struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Status int    `db:"status"`
}
