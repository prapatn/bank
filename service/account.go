package service

type AccountService interface {
	GetAccounts(int) ([]AccountResponse, error)
	NewAccount(int, NewAccountRequest) (*AccountResponse, error)
}

type AccountResponse struct {
	Id          int     `json:"id"`
	CustomerId  int     `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      int     `json:"status"`
}

type NewAccountRequest struct {
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}
