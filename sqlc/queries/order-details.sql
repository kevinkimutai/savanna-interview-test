-- name: GetOrderItem :one
SELECT * FROM order_items
WHERE order_item_id = $1 LIMIT 1;

-- name: ListOrderItems :many
SELECT * FROM order_items
ORDER BY created_at DESC;

-- name: CreateOrderItem :one
INSERT INTO order_items (
  order_id,product_id,quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE order_item_id = $1;