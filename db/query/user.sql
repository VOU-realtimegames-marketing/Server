-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email,
  role
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = sqlc.narg(username)
  OR email = sqlc.narg(email) LIMIT 1;

-- name: ListAllOtherUsers :many
SELECT * FROM users
WHERE username <> $1;

-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password),hashed_password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at),password_changed_at),
  full_name = COALESCE(sqlc.narg(full_name),full_name),
  email = COALESCE(sqlc.narg(email),email),
  photo = COALESCE(sqlc.narg(photo),photo),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified),is_email_verified),
  active = COALESCE(sqlc.narg(active),active)
WHERE username = sqlc.arg(username)
RETURNING *;