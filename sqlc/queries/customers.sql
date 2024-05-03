-- name: GetCustomer :one
SELECT * FROM customers
WHERE customer_id = $1 LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM customers
ORDER BY name;

-- name: CreateCustomer :one
INSERT INTO customers (
  customer_id,name,email
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers
  set name = $2
WHERE customer_id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE customer_id = $1;

-- name: GetCustomerByEmail :one
SELECT * FROM customers
WHERE email = $1
LIMIT 1;