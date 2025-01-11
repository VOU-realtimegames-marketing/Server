-- name: CreateEvent :one
INSERT INTO events (
  owner,
  game_id,
  store_id,
  name,
  photo,
  voucher_quantity,
  status,
  start_time,
  end_time
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
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

-- name: ListEventsOfOwner :many
SELECT E.id, E.owner, E.game_id, E.store_id, E.name, E.photo, E.voucher_quantity, E.status, E.start_time, E.end_time, G.type as game_type, S.name as store
FROM events E, games G, stores S
WHERE E.game_id = G.id AND E.store_id = S.id AND E.owner = $1;