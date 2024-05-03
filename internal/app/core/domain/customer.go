package domain

import (
	"errors"
)

type CustomerResponse struct {
	StatusCode uint     `json:"status_code"`
	Message    string   `json:"message"`
	Data       Customer `json:"data"`
}

type ErrorResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
}

type CustomersFetch struct {
	Page          uint     `json:"page"`
	NumberOfPages uint     `json:"number_of_pages"`
	Total         uint     `json:"total"`
	Data          Customer `json:"data"`
}

type CustomersResponse struct {
	StatusCode    uint     `json:"status_code"`
	Message       string   `json:"message"`
	Page          uint     `json:"page"`
	NumberOfPages uint     `json:"number_of_pages"`
	Total         uint     `json:"total"`
	Data          Customer `json:"data"`
}

type CustomerParams struct {
	Search string
	Page   string
	Limit  string
}

type Customer struct {
	CustomerID string `json:"customer_id"`
	Name       string `json:"name"`
}

func NewCustomerDomain(customer Customer) (Customer, error) {
	if customer.Name == "" {
		return customer, errors.New("missing name field")
	}

	return customer, nil
}

func CheckCustomerParams(m map[string]string) CustomerParams {

	cusParams := CustomerParams{}

	if m["search"] != "" {
		cusParams.Search = m["search"]
	}

	if m["page"] != "" {
		cusParams.Page = m["page"]
	}

	if m["limit"] != "" {
		cusParams.Limit = m["limit"]
	}

	return cusParams

}
