package db

import (
	"math"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/utils"
)

func (db *DBAdapter) CreateOrder(orderID string, customerID string) (domain.Order, error) {

	//convert to type params
	orderParams := queries.CreateOrderParams{
		OrderID:    orderID,
		CustomerID: pgtype.Text{String: customerID, Valid: true},
	}

	order, err := db.queries.CreateOrder(db.ctx, orderParams)
	if err != nil {
		return domain.Order{}, err
	}

	//Convert Order To domain.Order
	newOrder := domain.Order{
		OrderID:    order.OrderID,
		CustomerID: order.CustomerID.String,
		CreatedAt:  order.CreatedAt.Time,
	}

	return newOrder, nil

}

func (db *DBAdapter) GetTotalPrice(orderID string) (float64, error) {
	totalPrice, err := db.queries.TotalOrderPrice(db.ctx, orderID)
	if err != nil {
		return float64(totalPrice), err
	}

	return float64(totalPrice), nil
}

func (db *DBAdapter) UpdateOrderTotalPrice(orderID string, totalPrice float64) (domain.Order, error) {
	totalPriceStr := strconv.FormatFloat(totalPrice, 'f', 2, 64)
	var numeric pgtype.Numeric
	numeric.Scan(totalPriceStr)

	updateParams := queries.UpdateTotalPriceParams{
		OrderID:     orderID,
		TotalAmount: numeric,
	}
	order, err := db.queries.UpdateTotalPrice(db.ctx, updateParams)
	if err != nil {
		return domain.Order{}, err
	}

	return domain.Order{
		OrderID:     order.OrderID,
		CustomerID:  order.CustomerID.String,
		TotalAmount: totalPrice,
		CreatedAt:   order.CreatedAt.Time,
	}, nil
}
func (db *DBAdapter) GetOrderByID(orderID string) (domain.Order, error) {

	order, err := db.queries.GetOrder(db.ctx, orderID)
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

	err := db.queries.DeleteOrder(db.ctx, orderID)
	if err != nil {
		return err
	}

	return nil
}

func (db *DBAdapter) GetAllOrders(orderParams queries.ListOrdersParams) (domain.OrdersFetch, error) {
	//Get Orders
	orders, err := db.queries.ListOrders(db.ctx, orderParams)
	if err != nil {
		return domain.OrdersFetch{}, err

	}

	//Get Count
	count, err := db.queries.CountOrders(db.ctx)
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
