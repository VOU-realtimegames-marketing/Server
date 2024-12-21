-- name: ListBranchs :many
SELECT * FROM branchs
WHERE store_id = $1;

-- name: CreateBranch :one
INSERT INTO branchs (
  store_id,
  name,
  position,
  city_name,
  country,
  address,
  emoji
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: DeleteBranch :exec
DELETE FROM branchs
WHERE id = $1;