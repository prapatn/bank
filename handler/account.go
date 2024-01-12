package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) accountHandler {
	return accountHandler{service: service}
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(mux.Vars(r)["customerId"])
	customers, err := h.service.GetAccounts(customerId)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(mux.Vars(r)["customerId"])

	if r.Header.Get("content-type") != "application/json" {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	req := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	customer, err := h.service.NewAccount(customerId, req)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}
