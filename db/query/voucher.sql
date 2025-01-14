-- name: ListMyVouchers :many
SELECT V.id, V.event_id, V.qr_code, V.type, V.status, V.expires_at, V.discount
FROM voucher_owner O, vouchers V
WHERE O.voucher_id = V.id AND O.username = $1;

-- name: CreateVoucher :one
INSERT INTO vouchers (
  event_id,
  expires_at
) VALUES (
  $1, $2
) RETURNING *;

-- name: CreateVoucherOwner :one
INSERT INTO voucher_owner (
  username,
  voucher_id
) VALUES (
  $1, $2
) RETURNING *;