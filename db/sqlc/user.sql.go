// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email,
  role
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING username, hashed_password, full_name, email, role, photo, active, is_email_verified, password_changed_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.Photo,
		&i.Active,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, full_name, email, role, photo, active, is_email_verified, password_changed_at, created_at FROM users
WHERE username = $1
  OR email = $2 LIMIT 1
`

type GetUserParams struct {
	Username pgtype.Text `json:"username"`
	Email    pgtype.Text `json:"email"`
}

func (q *Queries) GetUser(ctx context.Context, arg GetUserParams) (User, error) {
	row := q.db.QueryRow(ctx, getUser, arg.Username, arg.Email)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.Photo,
		&i.Active,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listAllOtherUsers = `-- name: ListAllOtherUsers :many
SELECT username, hashed_password, full_name, email, role, photo, active, is_email_verified, password_changed_at, created_at FROM users
WHERE username <> $1
`

func (q *Queries) ListAllOtherUsers(ctx context.Context, username string) ([]User, error) {
	rows, err := q.db.Query(ctx, listAllOtherUsers, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.Role,
			&i.Photo,
			&i.Active,
			&i.IsEmailVerified,
			&i.PasswordChangedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE($1,hashed_password),
  password_changed_at = COALESCE($2,password_changed_at),
  full_name = COALESCE($3,full_name),
  email = COALESCE($4,email),
  photo = COALESCE($5,photo),
  is_email_verified = COALESCE($6,is_email_verified),
  active = COALESCE($7,active)
WHERE username = $8
RETURNING username, hashed_password, full_name, email, role, photo, active, is_email_verified, password_changed_at, created_at
`

type UpdateUserParams struct {
	HashedPassword    pgtype.Text        `json:"hashed_password"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	FullName          pgtype.Text        `json:"full_name"`
	Email             pgtype.Text        `json:"email"`
	Photo             pgtype.Text        `json:"photo"`
	IsEmailVerified   pgtype.Bool        `json:"is_email_verified"`
	Active            pgtype.Bool        `json:"active"`
	Username          string             `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.FullName,
		arg.Email,
		arg.Photo,
		arg.IsEmailVerified,
		arg.Active,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.Photo,
		&i.Active,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
