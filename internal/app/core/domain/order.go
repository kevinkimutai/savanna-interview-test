package domain

import (
	"time"

	"github.com/mr-tron/base58"

	"github.com/google/uuid"
)

type OrderResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
	Data       Order  `json:"data"`
}

type OrdersFetch struct {
	Page          uint    `json:"page"`
	NumberOfPages uint    `json:"number_of_pages"`
	Total         uint    `json:"total"`
	Data          []Order `json:"data"`
}

type OrdersResponse struct {
	StatusCode    uint    `json:"status_code"`
	Message       string  `json:"message"`
	Page          uint    `json:"page"`
	NumberOfPages uint    `json:"number_of_pages"`
	Total         uint    `json:"total"`
	Data          []Order `json:"data"`
}

type OrderParams struct {
	Search           string
	Page             string
	Limit            string
	TotalAmountStart string
	TotalAmountEnd   string
	StartDate        string
	EndDate          string
}

type Order struct {
	OrderID     string    `json:"order_id"`
	CustomerID  string    `json:"customer_id"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

// func NewOrderDomain(order Order) (Order, error) {
// 	//Get CustomerId From Order Domain
// 	if customer.Name == "" {
// 		return customer, errors.New("missing name field")
// 	}

// 	return customer, nil
// }

func CheckOrderParams(m map[string]string) OrderParams {

	oParams := OrderParams{}

	if m["search"] != "" {
		oParams.Search = m["search"]
	}

	if m["page"] != "" {
		oParams.Page = m["page"]
	}

	if m["limit"] != "" {
		oParams.Limit = m["limit"]
	}
	if m["total_amount_start"] != "" {
		oParams.TotalAmountStart = m["total_amount_start"]
	}
	if m["total_amount_end"] != "" {
		oParams.TotalAmountEnd = m["total_amount_end"]
	}
	if m["start_date"] != "" {
		oParams.StartDate = m["start_date"]
	}
	if m["end_date"] != "" {
		oParams.EndDate = m["end_date"]
	}

	return oParams

}

func GenerateUUID() string {
	u := uuid.New()
	// Encode UUID to Base64
	uBase58 := base58.Encode(u[:])

	return uBase58
}
