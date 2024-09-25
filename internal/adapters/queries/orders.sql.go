// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: orders.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countOrders = `-- name: CountOrders :one
SELECT COUNT(*) FROM orders
`

func (q *Queries) CountOrders(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countOrders)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
  order_id,customer_id,total_amount
) VALUES (
  $1, $2, $3
)
RETURNING order_id, customer_id, total_amount, created_at
`

type CreateOrderParams struct {
	OrderID     string
	CustomerID  pgtype.Text
	TotalAmount pgtype.Numeric
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, createOrder, arg.OrderID, arg.CustomerID, arg.TotalAmount)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.CustomerID,
		&i.TotalAmount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, orderID string) error {
	_, err := q.db.Exec(ctx, deleteOrder, orderID)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT order_id, customer_id, total_amount, created_at FROM orders
WHERE order_id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, orderID string) (Order, error) {
	row := q.db.QueryRow(ctx, getOrder, orderID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.CustomerID,
		&i.TotalAmount,
		&i.CreatedAt,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT order_id, customer_id, total_amount, created_at FROM orders
WHERE (order_id ILIKE '%' || $1 || '%' OR $1 IS NULL)
  AND (total_amount >= $2 OR $2 IS NULL)
  AND (total_amount <= $3 OR $3 IS NULL)
  AND (created_at >= $4 OR $4 IS NULL)
  AND (created_at <= $5 OR $5 IS NULL)
ORDER BY created_at DESC
LIMIT $6 OFFSET $7
`

type ListOrdersParams struct {
	Column1       pgtype.Text
	TotalAmount   pgtype.Numeric
	TotalAmount_2 pgtype.Numeric
	CreatedAt     pgtype.Timestamptz
	CreatedAt_2   pgtype.Timestamptz
	Limit         int32
	Offset        int32
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.Query(ctx, listOrders,
		arg.Column1,
		arg.TotalAmount,
		arg.TotalAmount_2,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.CustomerID,
			&i.TotalAmount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const totalOrderPrice = `-- name: TotalOrderPrice :one
SELECT
    SUM(p.price * oi.quantity) AS total_price
FROM
    orders o
JOIN
    order_items oi ON o.order_id = oi.order_id
JOIN
    products p ON oi.product_id = p.product_id
WHERE
    o.order_id = $1
`

func (q *Queries) TotalOrderPrice(ctx context.Context, orderID string) (int64, error) {
	row := q.db.QueryRow(ctx, totalOrderPrice, orderID)
	var total_price int64
	err := row.Scan(&total_price)
	return total_price, err
}

const updateTotalPrice = `-- name: UpdateTotalPrice :one
UPDATE orders
  SET total_amount = $2 
WHERE order_id = $1
RETURNING order_id, customer_id, total_amount, created_at
`

type UpdateTotalPriceParams struct {
	OrderID     string
	TotalAmount pgtype.Numeric
}

func (q *Queries) UpdateTotalPrice(ctx context.Context, arg UpdateTotalPriceParams) (Order, error) {
	row := q.db.QueryRow(ctx, updateTotalPrice, arg.OrderID, arg.TotalAmount)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.CustomerID,
		&i.TotalAmount,
		&i.CreatedAt,
	)
	return i, err
}
