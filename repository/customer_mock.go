package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	return customerRepositoryMock{customers: []Customer{
		{Id: 1, Name: "test-1", Status: 1},
		{Id: 2, Name: "test-2", Status: 1},
		{Id: 3, Name: "test-3", Status: 1},
	}}
}

func (m customerRepositoryMock) GetAll() ([]Customer, error) {
	return m.customers, nil
}

func (m customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, c := range m.customers {
		if c.Id == id {
			return &c, nil
		}
	}
	return nil, errors.New("Customer not found")
}
