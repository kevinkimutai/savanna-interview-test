package db

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/utils"
)

func (db *DBAdapter) CreateOrder(orderItems []domain.OrderItem, customerID string) (domain.Order, error) {
	ctx := context.Background()

	// Start Tx
	tx, err := db.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return domain.Order{}, errors.New("failed to start tx")
	}
	qtx := db.queries.WithTx(tx)

	// Save Order
	orderParams := queries.CreateOrderParams{
		OrderID:    orderItems[0].OrderID,
		CustomerID: pgtype.Text{String: customerID, Valid: true},
	}
	order, err := qtx.CreateOrder(ctx, orderParams)
	if err != nil {
		tx.Rollback(ctx)
		return domain.Order{}, err
	}

	// Save each orderItem
	var items []queries.OrderItem
	for _, item := range orderItems {
		productID, err := utils.ConvertStringToInt64(item.ProductID)
		if err != nil {
			tx.Rollback(ctx)
			return domain.Order{}, err
		}

		orderItem := queries.CreateOrderItemParams{
			OrderID:   item.OrderID,
			ProductID: productID,
			Quantity:  int32(item.Quantity),
		}

		orderItemResult, err := qtx.CreateOrderItem(ctx, orderItem)
		if err != nil {
			tx.Rollback(ctx)
			return domain.Order{}, err
		}

		items = append(items, orderItemResult)
	}

	// Calculate Total Price
	totalPrice, err := qtx.TotalOrderPrice(ctx, order.OrderID)
	if err != nil {
		tx.Rollback(ctx)
		return domain.Order{}, err
	}

	// Update Order
	var numeric pgtype.Numeric
	totalPricefloat64 := float64(totalPrice)
	totalPriceStr := strconv.FormatFloat(totalPricefloat64, 'f', 2, 64)

	numeric.Scan(totalPriceStr)

	updateParams := queries.UpdateTotalPriceParams{
		OrderID:     order.OrderID,
		TotalAmount: numeric,
	}

	updatedOrder, err := qtx.UpdateTotalPrice(ctx, updateParams)
	if err != nil {
		tx.Rollback(ctx)
		return domain.Order{}, err
	}

	fmt.Println(updatedOrder.TotalAmount)

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return domain.Order{}, err
	}

	return domain.Order{
		OrderID:     updatedOrder.OrderID,
		CustomerID:  updatedOrder.CustomerID.String,
		TotalAmount: utils.ConvertNumericToFloat64(updatedOrder.TotalAmount),
		CreatedAt:   order.CreatedAt.Time,
	}, nil
}

func (db *DBAdapter) GetOrderByID(orderID string) (domain.Order, error) {
	ctx := context.Background()

	order, err := db.queries.GetOrder(ctx, orderID)
	if err != nil {
		return domain.Order{}, err
	}

	return domain.Order{
		OrderID:     order.OrderID,
		CustomerID:  order.CustomerID.String,
		TotalAmount: utils.ConvertNumericToFloat64(order.TotalAmount),
		CreatedAt:   order.CreatedAt.Time,
	}, nil

}
func (db *DBAdapter) DeleteOrder(orderID string) error {
	ctx := context.Background()

	err := db.queries.DeleteOrder(ctx, orderID)
	if err != nil {
		return err
	}

	return nil
}

func (db *DBAdapter) GetAllOrders(orderParams queries.ListOrdersParams) (domain.OrdersFetch, error) {
	ctx := context.Background()

	//Get Orders
	orders, err := db.queries.ListOrders(ctx, orderParams)
	if err != nil {
		return domain.OrdersFetch{}, err

	}

	//Get Count
	count, err := db.queries.CountOrders(ctx)
	if err != nil {
		return domain.OrdersFetch{}, err

	}

	//Get Page
	page := utils.GetPage(orderParams.Offset, orderParams.Limit)

	//map struct
	var ords []domain.Order

	for _, item := range orders {
		// Convert each RequestItem to Product
		ord := domain.Order{
			OrderID:     item.OrderID,
			CustomerID:  item.CustomerID.String,
			TotalAmount: utils.ConvertNumericToFloat64(item.TotalAmount),
			CreatedAt:   item.CreatedAt.Time,
		}
		// Append the struct to the struct array
		ords = append(ords, ord)
	}

	return domain.OrdersFetch{
		Page:          page,
		NumberOfPages: uint(math.Ceil(float64(count) / float64(orderParams.Limit))),
		Total:         uint(count),
		Data:          ords,
	}, nil

}
