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

// func TestCreateOrder(t *testing.T) {
// 	app := setupApp()
// 	mockAPI := new(MockOrderApiPort)
// 	orderService := NewOrderService(mockAPI)
// 	app.Post("/orders", orderService.CreateOrder)

// 	customer := domain.Customer{
// 		CustomerID: "customer-123",
// 	}
// 	orderItems := []domain.OrderItemRequest{
// 		{
// 			ProductID: 1,
// 			Quantity:  2,
// 		},
// 	}
// 	uuid := "order-uuid"
// 	items := addUUID(uuid, orderItems)

// 	order := domain.Order{
// 		OrderID:    uuid,
// 		CustomerID: customer.CustomerID,
// 		Items:      items,
// 	}

// 	mockAPI.On("CreateOrder", items, customer.CustomerID).Return(order, nil)

// 	body, _ := json.Marshal(orderItems)
// 	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	req = req.WithContext(fiber.Ctx{Locals: func(key string) interface{} {
// 		if key == "customer" {
// 			return customer
// 		}
// 		return nil
// 	}})

// 	resp, err := app.Test(req, -1)
// 	require.NoError(t, err)
// 	assert.Equal(t, http.StatusCreated, resp.StatusCode)

// 	var response domain.OrderResponse
// 	err = json.NewDecoder(resp.Body).Decode(&response)
// 	require.NoError(t, err)
// 	assert.Equal(t, order.OrderID, response.Data.OrderID)
// }

func TestGetAllOrders(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockOrderApiPort)
	orderService := NewOrderService(mockAPI)
	app.Get("/orders", orderService.GetAllOrders)

	params := domain.OrderParams{}
	ordersFetch := domain.OrdersFetch{
		Page:          1,
		NumberOfPages: 1,
		Total:         1,
		Data: []domain.Order{
			{
				OrderID:    "order-123",
				CustomerID: "customer-123",
			},
		},
	}
	mockAPI.On("GetAllOrders", params).Return(ordersFetch, nil)

	req := httptest.NewRequest(http.MethodGet, "/orders", nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response domain.OrdersResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, ordersFetch.Page, response.Page)
	assert.Equal(t, ordersFetch.Total, response.Total)
	assert.Equal(t, ordersFetch.Data[0].OrderID, response.Data[0].OrderID)
}

func TestGetOrderByID(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockOrderApiPort)
	orderService := NewOrderService(mockAPI)
	app.Get("/orders/:orderID", orderService.GetOrderByID)

	orderID := "order-123"
	order := domain.Order{
		OrderID:    orderID,
		CustomerID: "customer-123",
	}
	mockAPI.On("GetOrderByID", orderID).Return(order, nil)

	req := httptest.NewRequest(http.MethodGet, "/orders/"+orderID, nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response domain.OrderResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, order.OrderID, response.Data.OrderID)
}

func TestDeleteOrder(t *testing.T) {
	app := setupApp()
	mockAPI := new(MockOrderApiPort)
	orderService := NewOrderService(mockAPI)
	app.Delete("/orders/:orderID", orderService.DeleteOrder)

	orderID := "order-123"
	mockAPI.On("DeleteOrder", orderID).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/orders/"+orderID, nil)

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
