package db

import (
	"testing"

	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	db := NewDB(config.DBURL)

	orderID := domain.GenerateUUID()
	customerID := config.CustomerID

	//Create Order
	order, err := db.CreateOrder(orderID, customerID)
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

	//check if proper details inserted

	assert.Equal(t, orderID, order.OrderID, "orderIDs should match")
	assert.Equal(t, customerID, order.CustomerID, "customersIDs should match")
}

func TestGetTotalPrice(t *testing.T) {
	db := NewDB(config.DBURL)

	orderID := domain.GenerateUUID()
	customerID := config.CustomerID

	//Create Order
	_, err := db.CreateOrder(orderID, customerID)
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

	//CreateOrderItems
	var orderItems = []domain.OrderItem{
		{OrderID: orderID,
			ProductID: "9",
			Quantity:  2,
		},

		{OrderID: orderID,
			ProductID: "11",
			Quantity:  1,
		},
	}

	for _, item := range orderItems {
		_, err := db.CreateOrderItem(item)
		if err != nil {
			t.Fatalf("Error creating orderItem: %v", err)
		}
	}

	totalPrice, err := db.GetTotalPrice(orderID)
	if err != nil {
		t.Fatalf("Error getting total price: %v", err)
	}

	assert.Equal(t, float64(7000), totalPrice, "total price should match")

	//TestOrderTotalPrice
	order, err := db.UpdateOrderTotalPrice(orderID, totalPrice)
	if err != nil {
		t.Fatalf("Error creating orderItem: %v", err)
	}

	assert.Equal(t, totalPrice, order.TotalAmount, "updated total amount should match")
	assert.Equal(t, orderID, order.OrderID, "orders id should match")
	assert.Equal(t, config.CustomerID, order.CustomerID, "customers id should match")
}

func TestGetOrderByID(t *testing.T) {
	db := NewDB(config.DBURL)

	orderID := "9e351749-2240-4c2b-b0c8-41c52ea65292"
	order, err := db.GetOrderByID(orderID)
	if err != nil {
		t.Fatalf("Error getting order by ID: %v", err)
	}

	assert.Equal(t, orderID, order.OrderID, "orders ids should match")
	assert.Equal(t, config.CustomerID, order.CustomerID, "customers id should match")
	assert.Equal(t, float64(7000), order.TotalAmount, "total amount should match")

}

func TestDeleteOrder(t *testing.T) {
	db := NewDB(config.DBURL)

	orderID := domain.GenerateUUID()
	customerID := config.CustomerID

	//Create Order
	order, err := db.CreateOrder(orderID, customerID)
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

	//Delete Order
	err = db.DeleteOrder(order.OrderID)
	if err != nil {
		t.Fatalf("Error deleting order: %v", err)
	}

	assert.Nil(t, err)

}
