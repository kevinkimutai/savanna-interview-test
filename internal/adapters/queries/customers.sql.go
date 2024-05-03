// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: customers.sql

package queries

import (
	"context"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers (
  customer_id,name,email
) VALUES (
  $1, $2, $3
)
RETURNING customer_id, name, created_at, email
`

type CreateCustomerParams struct {
	CustomerID string
	Name       string
	Email      string
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRow(ctx, createCustomer, arg.CustomerID, arg.Name, arg.Email)
	var i Customer
	err := row.Scan(
		&i.CustomerID,
		&i.Name,
		&i.CreatedAt,
		&i.Email,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE customer_id = $1
`

func (q *Queries) DeleteCustomer(ctx context.Context, customerID string) error {
	_, err := q.db.Exec(ctx, deleteCustomer, customerID)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT customer_id, name, created_at, email FROM customers
WHERE customer_id = $1 LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, customerID string) (Customer, error) {
	row := q.db.QueryRow(ctx, getCustomer, customerID)
	var i Customer
	err := row.Scan(
		&i.CustomerID,
		&i.Name,
		&i.CreatedAt,
		&i.Email,
	)
	return i, err
}

const getCustomerByEmail = `-- name: GetCustomerByEmail :one
SELECT customer_id, name, created_at, email FROM customers
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetCustomerByEmail(ctx context.Context, email string) (Customer, error) {
	row := q.db.QueryRow(ctx, getCustomerByEmail, email)
	var i Customer
	err := row.Scan(
		&i.CustomerID,
		&i.Name,
		&i.CreatedAt,
		&i.Email,
	)
	return i, err
}

const listCustomers = `-- name: ListCustomers :many
SELECT customer_id, name, created_at, email FROM customers
ORDER BY name
`

func (q *Queries) ListCustomers(ctx context.Context) ([]Customer, error) {
	rows, err := q.db.Query(ctx, listCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.CustomerID,
			&i.Name,
			&i.CreatedAt,
			&i.Email,
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

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE customers
  set name = $2
WHERE customer_id = $1
RETURNING customer_id, name, created_at, email
`

type UpdateCustomerParams struct {
	CustomerID string
	Name       string
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error) {
	row := q.db.QueryRow(ctx, updateCustomer, arg.CustomerID, arg.Name)
	var i Customer
	err := row.Scan(
		&i.CustomerID,
		&i.Name,
		&i.CreatedAt,
		&i.Email,
	)
	return i, err
}