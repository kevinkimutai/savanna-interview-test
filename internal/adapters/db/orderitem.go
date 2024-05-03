package db

import (
	"fmt"

	"github.com/kevinkimutai/savanna-app/internal/adapters/queries"
	"github.com/kevinkimutai/savanna-app/internal/app/core/domain"
	"github.com/kevinkimutai/savanna-app/internal/utils"
)

func (db *DBAdapter) CreateOrderItem(orderItem domain.OrderItem) (domain.OrderItem, error) {

	productID, err := utils.ConvertStringToInt64(orderItem.ProductID)
	if err != nil {
		fmt.Print("Error converting string to int64", err)
	}

	orderItemParams := queries.CreateOrderItemParams{
		OrderID:   orderItem.OrderID,
		ProductID: productID,
		Quantity:  int32(orderItem.Quantity),
	}

	oitem, err := db.queries.CreateOrderItem(db.ctx, orderItemParams)
	if err != nil {
		return domain.OrderItem{}, err
	}

	return domain.OrderItem{
		OrderItemID: int(oitem.OrderItemID),
		OrderID:     oitem.OrderID,
		ProductID:   utils.ConvertInt64ToString(oitem.ProductID),
		Quantity:    uint(oitem.Quantity),
		CreatedAt:   oitem.CreatedAt.Time,
	}, nil
}
