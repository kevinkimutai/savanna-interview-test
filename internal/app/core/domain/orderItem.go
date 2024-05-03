package domain

import (
	"errors"
	"time"
)

type OrderItemRequest struct {
	ProductID string `json:"product_id"`
	Quantity  uint   `json:"quantity"`
}

type OrderItem struct {
	OrderItemID int       `json:"order_item_id"`
	OrderID     string    `json:"order_id"`
	ProductID   string    `json:"product_id"`
	Quantity    uint      `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewOrderItemDomain(orderItem OrderItem) (OrderItem, error) {
	if orderItem.ProductID == "" {
		return orderItem, errors.New("missing product_id field")
	}
	if orderItem.Quantity == 0 {
		return orderItem, errors.New("missing quantity field")
	}

	return orderItem, nil
}
