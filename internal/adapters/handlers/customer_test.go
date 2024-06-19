package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllCustomers(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockCustomerApiPort)
	customerService := NewCustomerService(mockAPI)
	app.Get("/customers", customerService.GetAllCustomers)

	params := domain.CustomerParams{}
	customersFetch := domain.CustomersFetch{
		Page:          1,
		NumberOfPages: 1,
		Total:         1,
		Data: []domain.Customer{
			{
				CustomerID: "123",
				Name:       "John Doe",
				Email:      "john.doe@example.com",
			},
		},
	}
	mockAPI.On("GetAllCustomers", params).Return(customersFetch, nil)

	req := httptest.NewRequest(http.MethodGet, "/customers", nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response domain.CustomersResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, customersFetch.Page, response.Page)
	assert.Equal(t, customersFetch.Total, response.Total)
	assert.Equal(t, customersFetch.Data[0].CustomerID, response.Data[0].CustomerID)
	assert.Equal(t, customersFetch.Data[0].Name, response.Data[0].Name)
	assert.Equal(t, customersFetch.Data[0].Email, response.Data[0].Email)
}

func TestGetCustomerByID(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockCustomerApiPort)
	customerService := NewCustomerService(mockAPI)
	app.Get("/customers/:customerID", customerService.GetCustomerByID)

	customerID := "123"
	customer := domain.Customer{
		CustomerID: customerID,
		Name:       "John Doe",
		Email:      "john.doe@example.com",
	}
	mockAPI.On("GetCustomerByID", customerID).Return(customer, nil)

	req := httptest.NewRequest(http.MethodGet, "/customers/"+customerID, nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response domain.CustomerResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, customer.CustomerID, response.Data.CustomerID)
	assert.Equal(t, customer.Name, response.Data.Name)
	assert.Equal(t, customer.Email, response.Data.Email)
}

func TestDeleteCustomer(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockCustomerApiPort)
	customerService := NewCustomerService(mockAPI)
	app.Delete("/customers/:customerID", customerService.DeleteCustomer)

	customerID := "123"
	mockAPI.On("DeleteCustomer", customerID).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/customers/"+customerID, nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
