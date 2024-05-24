package db

import (
	"context"

	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
)

func (db *DBAdapter) CreateCustomer(customer queries.CreateCustomerParams) (queries.Customer, error) {
	ctx := context.Background()
	cus, err := db.queries.CreateCustomer(ctx, customer)
	if err != nil {
		return cus, err
	}

	return cus, nil
}

func (db *DBAdapter) GetCustomerByEmail(email string) (queries.Customer, error) {
	ctx := context.Background()
	customer, err := db.queries.GetCustomerByEmail(ctx, email)
	if err != nil {
		return customer, err
	}

	return customer, nil
}
