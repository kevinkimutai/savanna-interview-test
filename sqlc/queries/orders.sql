-- name: GetOrder :one
SELECT * FROM orders
WHERE order_id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
WHERE (order_id ILIKE '%' || $1 || '%' OR $1 IS NULL)
  AND (total_amount >= $2 OR $2 IS NULL)
  AND (total_amount <= $3 OR $3 IS NULL)
  AND (created_at >= $4 OR $4 IS NULL)
  AND (created_at <= $5 OR $5 IS NULL)
ORDER BY created_at DESC
LIMIT $6 OFFSET $7;

-- name: CreateOrder :one
INSERT INTO orders (
  order_id,customer_id,total_amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_id = $1;

-- name: TotalOrderPrice :one
SELECT
    SUM(p.price * oi.quantity) AS total_price
FROM
    orders o
JOIN
    order_items oi ON o.order_id = oi.order_id
JOIN
    products p ON oi.product_id = p.product_id
WHERE
    o.order_id = $1;

-- name: UpdateTotalPrice :one
UPDATE orders
  SET total_amount = $2 
WHERE order_id = $1
RETURNING *;

-- name: CountOrders :one
SELECT COUNT(*) FROM orders;