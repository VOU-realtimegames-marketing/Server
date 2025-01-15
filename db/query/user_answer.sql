-- name: CreateUserAnswer :one
INSERT INTO user_answers (
  username,
  event_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: UpdateUserAnswer :one
UPDATE user_answers
SET
  num_correct = num_correct + 1
WHERE username = sqlc.arg(username) AND event_id = sqlc.arg(event_id)
RETURNING *;