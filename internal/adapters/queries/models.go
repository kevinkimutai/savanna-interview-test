// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package queries

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Customer struct {
	CustomerID string
	Name       string
	CreatedAt  pgtype.Timestamptz
	Email      string
}

type Order struct {
	OrderID     string
	CustomerID  pgtype.Text
	TotalAmount pgtype.Numeric
	CreatedAt   pgtype.Timestamptz
}

type OrderItem struct {
	OrderItemID int64
	OrderID     string
	ProductID   int64
	Quantity    int32
	CreatedAt   pgtype.Timestamptz
}

type Product struct {
	ProductID int64
	Name      string
	Price     pgtype.Numeric
	ImageUrl  string
	CreatedAt pgtype.Timestamptz
}