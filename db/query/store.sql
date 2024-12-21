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

-- name: GetStoreByIdAndOwner :one
SELECT * FROM stores
WHERE id = $1 and owner = $2 LIMIT 1;

-- name: UpdateStore :one
UPDATE stores
SET
  name = COALESCE(sqlc.narg(name),name),
  business_type = COALESCE(sqlc.narg(business_type),business_type)
WHERE id = sqlc.arg(id) and owner = sqlc.arg(owner)
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1 and owner = $2;
