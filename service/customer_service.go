package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
)

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) customerService {
	return customerService{repository: repository}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.repository.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewInternalServerError()
	}
	response := []CustomerResponse{}
	for _, c := range customers {
		response = append(response, CustomerResponse{
			Id:     c.Id,
			Name:   c.Name,
			Status: c.Status,
		})
	}
	return response, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.repository.GetById(id)
	if err != nil {
		logs.Error(err)
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		return nil, errs.NewInternalServerError()
	}
	response := &CustomerResponse{
		Id:     customer.Id,
		Name:   customer.Name,
		Status: customer.Status,
	}
	return response, err
}
