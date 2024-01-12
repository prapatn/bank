package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
	"strings"
)

type accountService struct {
	repository repository.AccountRepository
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	return accountService{repository: repository}
}

func (s accountService) GetAccounts(customerId int) ([]AccountResponse, error) {
	accounts, err := s.repository.GetAll(customerId)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewInternalServerError()
	}
	response := []AccountResponse{}
	for _, a := range accounts {
		response = append(response, AccountResponse{
			Id:          a.Id,
			CustomerId:  a.CustomerId,
			AccountType: a.AccountType,
			Amount:      a.Amount,
			Status:      a.Status,
		})
	}
	return response, nil
}

func (s accountService) NewAccount(customerId int, req NewAccountRequest) (*AccountResponse, error) {
	//validate
	acctype := strings.ToLower(req.AccountType)
	if acctype != "saving" && acctype != "checking" {
		return nil, errs.NewValidationError("account type only saving or checking")
	}

	if req.Amount < 1000 {
		return nil, errs.NewValidationError("amount is least 1,000")
	}

	account, err := s.repository.Create(repository.Account{
		CustomerId:  customerId,
		AccountType: acctype,
		Amount:      req.Amount,
		Status:      1,
	})
	if err != nil {
		logs.Error(err)
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		return nil, errs.NewInternalServerError()
	}
	response := &AccountResponse{
		Id:          account.Id,
		CustomerId:  account.CustomerId,
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.Status,
	}
	return response, err
}
