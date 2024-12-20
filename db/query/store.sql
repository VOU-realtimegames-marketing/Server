-- name: CreateStore :one
INSERT INTO stores (
  owner,
  name,
  business_type
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ListStoresOfOwner :many
SELECT * FROM stores
WHERE owner = $1;
