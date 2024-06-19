package handler

import (
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockCustomerApiPort struct {
	mock.Mock
}

func (m *MockCustomerApiPort) CreateCustomer(customer domain.Customer) (domain.Customer, error) {
	args := m.Called(customer)
	return args.Get(0).(domain.Customer), args.Error(1)
}

func (m *MockCustomerApiPort) GetAllCustomers(params domain.CustomerParams) (domain.CustomersFetch, error) {
	args := m.Called(params)
	return args.Get(0).(domain.CustomersFetch), args.Error(1)
}

func (m *MockCustomerApiPort) GetCustomerByID(id string) (domain.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Customer), args.Error(1)
}

func (m *MockCustomerApiPort) DeleteCustomer(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
