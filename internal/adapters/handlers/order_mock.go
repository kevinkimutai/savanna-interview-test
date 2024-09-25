package handler

import (
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockOrderApiPort struct {
	mock.Mock
}

func (m *MockOrderApiPort) CreateOrder(items []domain.OrderItem, phonenumber int, customer queries.Customer) (domain.Order, error) {
	args := m.Called(items, customer.CustomerID)
	return args.Get(0).(domain.Order), args.Error(1)
}

func (m *MockOrderApiPort) GetAllOrders(params domain.OrderParams) (domain.OrdersFetch, error) {
	args := m.Called(params)
	return args.Get(0).(domain.OrdersFetch), args.Error(1)
}

func (m *MockOrderApiPort) GetOrderByID(orderID string) (domain.Order, error) {
	args := m.Called(orderID)
	return args.Get(0).(domain.Order), args.Error(1)
}

func (m *MockOrderApiPort) DeleteOrder(orderID string) error {
	args := m.Called(orderID)
	return args.Error(0)
}
