// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: voucher.sql

package db

import (
	"context"
	"time"
)

const createVoucher = `-- name: CreateVoucher :one
INSERT INTO vouchers (
  event_id,
  expires_at
) VALUES (
  $1, $2
) RETURNING id, event_id, qr_code, type, status, expires_at, discount
`

type CreateVoucherParams struct {
	EventID   int64     `json:"event_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (q *Queries) CreateVoucher(ctx context.Context, arg CreateVoucherParams) (Voucher, error) {
	row := q.db.QueryRow(ctx, createVoucher, arg.EventID, arg.ExpiresAt)
	var i Voucher
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.QrCode,
		&i.Type,
		&i.Status,
		&i.ExpiresAt,
		&i.Discount,
	)
	return i, err
}

const createVoucherOwner = `-- name: CreateVoucherOwner :one
INSERT INTO voucher_owner (
  username,
  voucher_id
) VALUES (
  $1, $2
) RETURNING id, username, voucher_id, created_at
`

type CreateVoucherOwnerParams struct {
	Username  string `json:"username"`
	VoucherID int64  `json:"voucher_id"`
}

func (q *Queries) CreateVoucherOwner(ctx context.Context, arg CreateVoucherOwnerParams) (VoucherOwner, error) {
	row := q.db.QueryRow(ctx, createVoucherOwner, arg.Username, arg.VoucherID)
	var i VoucherOwner
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.VoucherID,
		&i.CreatedAt,
	)
	return i, err
}

const listMyVouchers = `-- name: ListMyVouchers :many
SELECT V.id, V.event_id, V.qr_code, V.type, V.status, V.expires_at, V.discount
FROM voucher_owner O, vouchers V
WHERE O.voucher_id = V.id AND O.username = $1
`

func (q *Queries) ListMyVouchers(ctx context.Context, username string) ([]Voucher, error) {
	rows, err := q.db.Query(ctx, listMyVouchers, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Voucher{}
	for rows.Next() {
		var i Voucher
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.QrCode,
			&i.Type,
			&i.Status,
			&i.ExpiresAt,
			&i.Discount,
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
