-- name: CreateEvent :one
INSERT INTO events (
  game_id,
  store_id,
  name,
  photo,
  voucher_quantity,
  status,
  start_time,
  end_time
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: UpdateEvent :one
UPDATE events
SET
  name = COALESCE(sqlc.narg(name),name),
  photo = COALESCE(sqlc.narg(photo),photo),
  status = COALESCE(sqlc.narg(status),status),
  start_time = COALESCE(sqlc.narg(start_time),start_time),
  end_time = COALESCE(sqlc.narg(end_time),end_time)
WHERE id = sqlc.arg(id)
RETURNING *;