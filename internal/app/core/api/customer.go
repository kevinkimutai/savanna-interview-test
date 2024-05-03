package application

import "github.com/kevinkimutai/savanna-app/internal/ports"

type CustomerRepo struct {
	db ports.CustomerRepoPort
}

func NewCustomerRepo(db ports.CustomerRepoPort) *CustomerRepo {
	return &CustomerRepo{db: db}
}

func (c CustomerRepo) CreateCustomer() {}
