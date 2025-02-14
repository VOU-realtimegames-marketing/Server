// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: store.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createStore = `-- name: CreateStore :one
INSERT INTO stores (
  owner,
  name,
  business_type
) VALUES (
  $1, $2, $3
) RETURNING id, owner, name, business_type
`

type CreateStoreParams struct {
	Owner        string `json:"owner"`
	Name         string `json:"name"`
	BusinessType string `json:"business_type"`
}

func (q *Queries) CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error) {
	row := q.db.QueryRow(ctx, createStore, arg.Owner, arg.Name, arg.BusinessType)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.BusinessType,
	)
	return i, err
}

const deleteStore = `-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1 and owner = $2
`

type DeleteStoreParams struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
}

func (q *Queries) DeleteStore(ctx context.Context, arg DeleteStoreParams) error {
	_, err := q.db.Exec(ctx, deleteStore, arg.ID, arg.Owner)
	return err
}

const getStoreByIdAndOwner = `-- name: GetStoreByIdAndOwner :one
SELECT id, owner, name, business_type FROM stores
WHERE id = $1 and owner = $2 LIMIT 1
`

type GetStoreByIdAndOwnerParams struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
}

func (q *Queries) GetStoreByIdAndOwner(ctx context.Context, arg GetStoreByIdAndOwnerParams) (Store, error) {
	row := q.db.QueryRow(ctx, getStoreByIdAndOwner, arg.ID, arg.Owner)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.BusinessType,
	)
	return i, err
}

const listStoresOfOwner = `-- name: ListStoresOfOwner :many
SELECT id, owner, name, business_type FROM stores
WHERE owner = $1
`

func (q *Queries) ListStoresOfOwner(ctx context.Context, owner string) ([]Store, error) {
	rows, err := q.db.Query(ctx, listStoresOfOwner, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Store{}
	for rows.Next() {
		var i Store
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Name,
			&i.BusinessType,
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

const updateStore = `-- name: UpdateStore :one
UPDATE stores
SET
  name = COALESCE($1,name),
  business_type = COALESCE($2,business_type)
WHERE id = $3 and owner = $4
RETURNING id, owner, name, business_type
`

type UpdateStoreParams struct {
	Name         pgtype.Text `json:"name"`
	BusinessType pgtype.Text `json:"business_type"`
	ID           int64       `json:"id"`
	Owner        string      `json:"owner"`
}

func (q *Queries) UpdateStore(ctx context.Context, arg UpdateStoreParams) (Store, error) {
	row := q.db.QueryRow(ctx, updateStore,
		arg.Name,
		arg.BusinessType,
		arg.ID,
		arg.Owner,
	)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.BusinessType,
	)
	return i, err
}
