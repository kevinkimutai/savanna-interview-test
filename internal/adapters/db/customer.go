package db

import "github.com/kevinkimutai/savanna-app/internal/adapters/queries"

func (db *DBAdapter) CreateCustomer(customer queries.CreateCustomerParams) (queries.Customer, error) {
	cus, err := db.queries.CreateCustomer(db.ctx, customer)
	if err != nil {
		return cus, err
	}

	return cus, nil
}

func (db *DBAdapter) GetCustomerByEmail(email string) (queries.Customer, error) {
	customer, err := db.queries.GetCustomerByEmail(db.ctx, email)
	if err != nil {
		return customer, err
	}

	return customer, nil
}
