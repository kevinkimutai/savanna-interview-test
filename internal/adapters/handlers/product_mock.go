package handler

import (
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockProductApiPort struct {
	mock.Mock
}

func (m *MockProductApiPort) CreateProduct(product domain.Product) (domain.Product, error) {
	args := m.Called(product)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductApiPort) GetAllProducts(params domain.ProductParams) (domain.ProductsFetch, error) {
	args := m.Called(params)
	return args.Get(0).(domain.ProductsFetch), args.Error(1)
}

func (m *MockProductApiPort) GetProduct(id string) (domain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductApiPort) DeleteProduct(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
