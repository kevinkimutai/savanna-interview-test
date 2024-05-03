package domain

import (
	"errors"
	"time"
)

type ProductResponse struct {
	StatusCode uint    `json:"status_code"`
	Message    string  `json:"message"`
	Data       Product `json:"data"`
}

type ProductsFetch struct {
	Page          uint      `json:"page"`
	NumberOfPages uint      `json:"number_of_pages"`
	Total         uint      `json:"total"`
	Data          []Product `json:"data"`
}

type ProductsResponse struct {
	StatusCode    uint      `json:"status_code"`
	Message       string    `json:"message"`
	Page          uint      `json:"page"`
	NumberOfPages uint      `json:"number_of_pages"`
	Total         uint      `json:"total"`
	Data          []Product `json:"data"`
}

type ProductParams struct {
	//TODO:ADD MORE PARAMS
	Search     string
	Page       string
	Limit      string
	PriceStart string
	PriceEnd   string
	StartDate  string
	EndDate    string
}

type Product struct {
	ProductID int       `json:"product_id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProductDomain(product Product) (Product, error) {
	if product.Name == "" {
		return product, errors.New("missing name field")
	}
	if product.Price == 0 {
		return product, errors.New("missing price field")
	}
	if product.ImageURL == "" {
		return product, errors.New("missing image_url field")
	}

	return product, nil
}

func CheckProductsParams(m map[string]string) ProductParams {

	prodParams := ProductParams{}

	if m["search"] != "" {
		prodParams.Search = m["search"]
	}
	if m["page"] != "" {
		prodParams.Page = m["page"]
	}
	if m["limit"] != "" {
		prodParams.Limit = m["limit"]
	}
	if m["price_start"] != "" {
		prodParams.PriceStart = m["price_start"]
	}
	if m["price_end"] != "" {
		prodParams.PriceEnd = m["price_end"]
	}
	if m["start_date"] != "" {
		prodParams.StartDate = m["start_date"]
	}
	if m["end_date"] != "" {
		prodParams.EndDate = m["end_date"]
	}

	return prodParams

}
