-- name: GetProduct :one
SELECT * FROM products
WHERE product_id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
WHERE (name ILIKE '%' || $1 || '%' OR $1 IS NULL)
  AND (price >= $2 OR $2 IS NULL)
  AND (price <= $3 OR $3 IS NULL)
ORDER BY created_at
LIMIT $4 OFFSET $5;


-- name: CreateProduct :one
INSERT INTO products (
  name,price,image_url
) VALUES (
  $1, $2 , $3
)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
  set name = $2,
  price = $3,
  image_url = $4
WHERE product_id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1;

-- name: CountProducts :one
SELECT COUNT(*) FROM products
WHERE (name ILIKE '%' || $1 || '%' OR $1 IS NULL)
  AND (price >= $2 OR $2 IS NULL)
  AND (price <= $3 OR $3 IS NULL);
