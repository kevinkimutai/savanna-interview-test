package db

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
)

func TestCreateOrder(t *testing.T) {
	db := SetupTestDB()

	// Create a test product
	product := domain.Product{
		Name:     "Test Product",
		Price:    50.00,
		ImageURL: "http://example.com/image.png",
	}
	createdProduct, err := db.CreateProduct(product)
	require.NoError(t, err)

	// Create a test customer
	customer := queries.CreateCustomerParams{
		CustomerID: "xyz",
		Name:       "Test Customer",
		Email:      "test@example.com",
	}
	createdCustomer, err := db.CreateCustomer(customer)
	require.NoError(t, err)

	// Test data
	orderItems := []domain.OrderItem{
		{
			OrderID:   "order-1",
			ProductID: strconv.Itoa(createdProduct.ProductID),
			Quantity:  2},
	}
	customerID := createdCustomer.CustomerID

	// Test the CreateOrder method
	createdOrder, err := db.CreateOrder(orderItems, customerID)

	// Assert expectations
	require.NoError(t, err)
	assert.Equal(t, "order-1", createdOrder.OrderID)
	assert.Equal(t, customerID, createdOrder.CustomerID)
	assert.Equal(t, float64(100.00), createdOrder.TotalAmount)
}

func TestGetOrderByID(t *testing.T) {
	db := SetupTestDB()

	// test order
	orderID := "order-1"
	customerID := "xyz"

	// Test the GetOrderByID method
	fetchedOrder, err := db.GetOrderByID(orderID)

	// Assert expectations
	require.NoError(t, err)
	assert.Equal(t, orderID, fetchedOrder.OrderID)
	assert.Equal(t, customerID, fetchedOrder.CustomerID)
}

func TestDeleteOrder(t *testing.T) {
	db := SetupTestDB()
	defer TeardownTestDB(db)

	// test order
	orderID := "order-1"

	// Test the DeleteOrder method
	err := db.DeleteOrder(orderID)

	// Assert expectations
	require.NoError(t, err)

	// Verify the order is deleted
	_, err = db.GetOrderByID(orderID)
	require.Error(t, err)

	if err != nil {
		require.Contains(t, err.Error(), "no rows in result set")
	}
}

// func TestGetAllOrders(t *testing.T) {
// 	db := SetupTestDB()
// 	defer TeardownTestDB(db)

// 	// Create test orders
// 	orderParams := queries.ListOrdersParams{
// 		Offset: 0,
// 		Limit:  10,
// 	}

// 	order1 := queries.CreateOrderParams{
// 		OrderID:    "order-1",
// 		CustomerID: pgtype.Text{String: "customer-1", Valid: true},
// 	}
// 	_, err := db.queries.CreateOrder(context.Background(), order1)
// 	require.NoError(t, err)

// 	order2 := queries.CreateOrderParams{
// 		OrderID:    "order-2",
// 		CustomerID: pgtype.Text{String: "customer-2", Valid: true},
// 	}
// 	_, err = db.queries.CreateOrder(context.Background(), order2)
// 	require.NoError(t, err)

// 	// Test the GetAllOrders method
// 	ordersFetch, err := db.GetAllOrders(orderParams)

// 	// Assert expectations
// 	require.NoError(t, err)
// 	assert.Equal(t, 1, ordersFetch.Page)                // Assuming offset/limit gives first page
// 	assert.Equal(t, uint(1), ordersFetch.NumberOfPages) // Assuming 10 orders per page
// 	assert.Equal(t, uint(2), ordersFetch.Total)         // Total number of orders
// 	assert.Len(t, ordersFetch.Data, 2)                  // Number of orders fetched
// 	assert.Equal(t, "order-1", ordersFetch.Data[0].OrderID)
// 	assert.Equal(t, "customer-1", ordersFetch.Data[0].CustomerID)
// 	assert.Equal(t, "order-2", ordersFetch.Data[1].OrderID)
// 	assert.Equal(t, "customer-2", ordersFetch.Data[1].CustomerID)
// }
