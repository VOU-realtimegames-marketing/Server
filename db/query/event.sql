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