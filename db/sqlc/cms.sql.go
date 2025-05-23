// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: cms.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkEventExists = `-- name: CheckEventExists :one
SELECT id 
FROM events
WHERE owner = $1 
  AND game_id = $2 
  AND store_id = $3 
  AND name = $4
`

type CheckEventExistsParams struct {
	Owner   string `json:"owner"`
	GameID  int64  `json:"game_id"`
	StoreID int64  `json:"store_id"`
	Name    string `json:"name"`
}

func (q *Queries) CheckEventExists(ctx context.Context, arg CheckEventExistsParams) (int64, error) {
	row := q.db.QueryRow(ctx, checkEventExists,
		arg.Owner,
		arg.GameID,
		arg.StoreID,
		arg.Name,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const checkGameExists = `-- name: CheckGameExists :one
SELECT id 
FROM games
WHERE type = $1
`

func (q *Queries) CheckGameExists(ctx context.Context, type_ string) (int64, error) {
	row := q.db.QueryRow(ctx, checkGameExists, type_)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const checkStoreExists = `-- name: CheckStoreExists :one
SELECT id
FROM stores
WHERE name = $1 AND owner = $2
LIMIT 1
`

type CheckStoreExistsParams struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func (q *Queries) CheckStoreExists(ctx context.Context, arg CheckStoreExistsParams) (int64, error) {
	row := q.db.QueryRow(ctx, checkStoreExists, arg.Name, arg.Owner)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const checkVoucherExists = `-- name: CheckVoucherExists :one
SELECT id 
FROM vouchers
WHERE event_id = $1 AND qr_code = $2 AND type = $3
`

type CheckVoucherExistsParams struct {
	EventID int64       `json:"event_id"`
	QrCode  pgtype.Text `json:"qr_code"`
	Type    string      `json:"type"`
}

func (q *Queries) CheckVoucherExists(ctx context.Context, arg CheckVoucherExistsParams) (int64, error) {
	row := q.db.QueryRow(ctx, checkVoucherExists, arg.EventID, arg.QrCode, arg.Type)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const checkVoucherOwnerExists = `-- name: CheckVoucherOwnerExists :one
SELECT id 
FROM voucher_owner
WHERE username = $1 AND voucher_id = $2
`

type CheckVoucherOwnerExistsParams struct {
	Username  string `json:"username"`
	VoucherID int64  `json:"voucher_id"`
}

func (q *Queries) CheckVoucherOwnerExists(ctx context.Context, arg CheckVoucherOwnerExistsParams) (int64, error) {
	row := q.db.QueryRow(ctx, checkVoucherOwnerExists, arg.Username, arg.VoucherID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const countVoucherOwners = `-- name: CountVoucherOwners :one
SELECT COUNT(*) AS total
FROM voucher_owner
`

func (q *Queries) CountVoucherOwners(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countVoucherOwners)
	var total int64
	err := row.Scan(&total)
	return total, err
}

const createFakeBranch = `-- name: CreateFakeBranch :one
INSERT INTO branchs (store_id, name, position, city_name, country, address, emoji)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (store_id, name) DO NOTHING
RETURNING id
`

type CreateFakeBranchParams struct {
	StoreID  int64  `json:"store_id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	CityName string `json:"city_name"`
	Country  string `json:"country"`
	Address  string `json:"address"`
	Emoji    string `json:"emoji"`
}

func (q *Queries) CreateFakeBranch(ctx context.Context, arg CreateFakeBranchParams) (int64, error) {
	row := q.db.QueryRow(ctx, createFakeBranch,
		arg.StoreID,
		arg.Name,
		arg.Position,
		arg.CityName,
		arg.Country,
		arg.Address,
		arg.Emoji,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createFakeEvent = `-- name: CreateFakeEvent :one
INSERT INTO events (game_id, store_id, owner, name, photo, voucher_quantity, status, start_time, end_time)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id
`

type CreateFakeEventParams struct {
	GameID          int64        `json:"game_id"`
	StoreID         int64        `json:"store_id"`
	Owner           string       `json:"owner"`
	Name            string       `json:"name"`
	Photo           string       `json:"photo"`
	VoucherQuantity int32        `json:"voucher_quantity"`
	Status          EventsStatus `json:"status"`
	StartTime       time.Time    `json:"start_time"`
	EndTime         time.Time    `json:"end_time"`
}

func (q *Queries) CreateFakeEvent(ctx context.Context, arg CreateFakeEventParams) (int64, error) {
	row := q.db.QueryRow(ctx, createFakeEvent,
		arg.GameID,
		arg.StoreID,
		arg.Owner,
		arg.Name,
		arg.Photo,
		arg.VoucherQuantity,
		arg.Status,
		arg.StartTime,
		arg.EndTime,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createFakeStore = `-- name: CreateFakeStore :one
INSERT INTO stores (name, owner, business_type)
VALUES ($1, $2, $3)
RETURNING id
`

type CreateFakeStoreParams struct {
	Name         string `json:"name"`
	Owner        string `json:"owner"`
	BusinessType string `json:"business_type"`
}

func (q *Queries) CreateFakeStore(ctx context.Context, arg CreateFakeStoreParams) (int64, error) {
	row := q.db.QueryRow(ctx, createFakeStore, arg.Name, arg.Owner, arg.BusinessType)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createFakeUser = `-- name: CreateFakeUser :one
INSERT INTO users (username, hashed_password, full_name, email, role, photo, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (username) DO UPDATE SET 
    hashed_password = EXCLUDED.hashed_password,
    full_name = EXCLUDED.full_name,
    email = EXCLUDED.email,
    role = EXCLUDED.role,
    photo = EXCLUDED.photo,
    created_at = EXCLUDED.created_at
RETURNING username
`

type CreateFakeUserParams struct {
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	Role           string    `json:"role"`
	Photo          string    `json:"photo"`
	CreatedAt      time.Time `json:"created_at"`
}

func (q *Queries) CreateFakeUser(ctx context.Context, arg CreateFakeUserParams) (string, error) {
	row := q.db.QueryRow(ctx, createFakeUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.Role,
		arg.Photo,
		arg.CreatedAt,
	)
	var username string
	err := row.Scan(&username)
	return username, err
}

const createFakeVoucher = `-- name: CreateFakeVoucher :one
INSERT INTO vouchers (event_id, qr_code, type, status, expires_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type CreateFakeVoucherParams struct {
	EventID   int64       `json:"event_id"`
	QrCode    pgtype.Text `json:"qr_code"`
	Type      string      `json:"type"`
	Status    string      `json:"status"`
	ExpiresAt time.Time   `json:"expires_at"`
}

func (q *Queries) CreateFakeVoucher(ctx context.Context, arg CreateFakeVoucherParams) (int64, error) {
	row := q.db.QueryRow(ctx, createFakeVoucher,
		arg.EventID,
		arg.QrCode,
		arg.Type,
		arg.Status,
		arg.ExpiresAt,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createFakeVoucherOwner = `-- name: CreateFakeVoucherOwner :one
INSERT INTO voucher_owner (username, voucher_id, created_at)
VALUES ($1, $2, $3)
RETURNING id
`

type CreateFakeVoucherOwnerParams struct {
	Username  string    `json:"username"`
	VoucherID int64     `json:"voucher_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateFakeVoucherOwner(ctx context.Context, arg CreateFakeVoucherOwnerParams) (int64, error) {
	row := q.db.QueryRow(ctx, createFakeVoucherOwner, arg.Username, arg.VoucherID, arg.CreatedAt)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createGame = `-- name: CreateGame :one
INSERT INTO games (name, photo, type, play_guide, gift_allowed)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type CreateGameParams struct {
	Name        string      `json:"name"`
	Photo       string      `json:"photo"`
	Type        string      `json:"type"`
	PlayGuide   pgtype.Text `json:"play_guide"`
	GiftAllowed bool        `json:"gift_allowed"`
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (int64, error) {
	row := q.db.QueryRow(ctx, createGame,
		arg.Name,
		arg.Photo,
		arg.Type,
		arg.PlayGuide,
		arg.GiftAllowed,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getAdminStats = `-- name: GetAdminStats :one
SELECT 
    (SELECT COUNT(*) FROM users WHERE role = 'partner') AS total_partner,
    (SELECT COUNT(*) FROM users WHERE role = 'partner' AND created_at >= NOW() - INTERVAL '1 MONTH') AS total_partner_last_month,
    (SELECT COUNT(*) FROM users WHERE role = 'user') AS total_user,
    (SELECT COUNT(*) FROM users WHERE role = 'user' AND created_at >= NOW() - INTERVAL '1 MONTH') AS total_user_last_month,
    (SELECT COUNT(*) FROM branchs) AS total_branch,
    (SELECT COUNT(*) FROM branchs) AS total_branch
`

type GetAdminStatsRow struct {
	TotalPartner          int64 `json:"total_partner"`
	TotalPartnerLastMonth int64 `json:"total_partner_last_month"`
	TotalUser             int64 `json:"total_user"`
	TotalUserLastMonth    int64 `json:"total_user_last_month"`
	TotalBranch           int64 `json:"total_branch"`
	TotalBranch_2         int64 `json:"total_branch_2"`
}

func (q *Queries) GetAdminStats(ctx context.Context) (GetAdminStatsRow, error) {
	row := q.db.QueryRow(ctx, getAdminStats)
	var i GetAdminStatsRow
	err := row.Scan(
		&i.TotalPartner,
		&i.TotalPartnerLastMonth,
		&i.TotalUser,
		&i.TotalUserLastMonth,
		&i.TotalBranch,
		&i.TotalBranch_2,
	)
	return i, err
}

const getBranchesByStore = `-- name: GetBranchesByStore :many
SELECT b.id, b.name, b.position, b.city_name, b.country, b.address, b.emoji
FROM branchs b
WHERE b.store_id = $1
`

type GetBranchesByStoreRow struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	CityName string `json:"city_name"`
	Country  string `json:"country"`
	Address  string `json:"address"`
	Emoji    string `json:"emoji"`
}

func (q *Queries) GetBranchesByStore(ctx context.Context, storeID int64) ([]GetBranchesByStoreRow, error) {
	rows, err := q.db.Query(ctx, getBranchesByStore, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetBranchesByStoreRow{}
	for rows.Next() {
		var i GetBranchesByStoreRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Position,
			&i.CityName,
			&i.Country,
			&i.Address,
			&i.Emoji,
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

const getCmsOverview = `-- name: GetCmsOverview :one
SELECT 
    -- Tổng số cửa hàng của owner
    (SELECT COUNT(*) 
     FROM stores s 
     WHERE s.owner = $1) AS total_store,

    -- Tổng số chi nhánh thuộc các cửa hàng của owner
    (SELECT COUNT(*) 
     FROM branchs b 
     JOIN stores s ON b.store_id = s.id 
     WHERE s.owner = $1) AS total_branch,

    -- Tổng số sự kiện của owner
    (SELECT COUNT(*) 
     FROM events e 
     WHERE e.owner = $1) AS total_event,

    -- Tổng số user chơi game của owner
    (SELECT COUNT(*) 
     FROM voucher_owner vo 
     JOIN vouchers v ON vo.voucher_id = v.id 
     JOIN events e ON v.event_id = e.id 
     WHERE e.owner = $1) AS total_user_play,

    -- Tổng số user chơi game tháng trước của owner
    (SELECT COUNT(*) 
     FROM voucher_owner vo 
     JOIN vouchers v ON vo.voucher_id = v.id 
     JOIN events e ON v.event_id = e.id 
     WHERE e.owner = $1 
       AND e.start_time >= DATE_TRUNC('month', NOW() - INTERVAL '1 MONTH') 
       AND e.start_time < DATE_TRUNC('month', NOW())) AS last_month_total_user_play
`

type GetCmsOverviewRow struct {
	TotalStore             int64 `json:"total_store"`
	TotalBranch            int64 `json:"total_branch"`
	TotalEvent             int64 `json:"total_event"`
	TotalUserPlay          int64 `json:"total_user_play"`
	LastMonthTotalUserPlay int64 `json:"last_month_total_user_play"`
}

func (q *Queries) GetCmsOverview(ctx context.Context, owner string) (GetCmsOverviewRow, error) {
	row := q.db.QueryRow(ctx, getCmsOverview, owner)
	var i GetCmsOverviewRow
	err := row.Scan(
		&i.TotalStore,
		&i.TotalBranch,
		&i.TotalEvent,
		&i.TotalUserPlay,
		&i.LastMonthTotalUserPlay,
	)
	return i, err
}

const getEventCreatedStats = `-- name: GetEventCreatedStats :many
SELECT 
    DATE(start_time) AS date,
    COUNT(CASE WHEN g.type = 'quiz' THEN 1 END) AS quiz_game,
    COUNT(CASE WHEN g.type = 'phone-shake' THEN 1 END) AS shake_game
FROM events e
JOIN games g ON e.game_id = g.id
WHERE start_time >= NOW() - INTERVAL '2 MONTHS'
GROUP BY DATE(start_time)
ORDER BY DATE(start_time)
`

type GetEventCreatedStatsRow struct {
	Date      pgtype.Date `json:"date"`
	QuizGame  int64       `json:"quiz_game"`
	ShakeGame int64       `json:"shake_game"`
}

func (q *Queries) GetEventCreatedStats(ctx context.Context) ([]GetEventCreatedStatsRow, error) {
	rows, err := q.db.Query(ctx, getEventCreatedStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetEventCreatedStatsRow{}
	for rows.Next() {
		var i GetEventCreatedStatsRow
		if err := rows.Scan(&i.Date, &i.QuizGame, &i.ShakeGame); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventsByStore = `-- name: GetEventsByStore :many
SELECT e.id, e.name, e.photo, e.voucher_quantity, e.status, e.start_time, e.end_time, g.type AS game_type
FROM events e
JOIN games g ON e.game_id = g.id
WHERE e.store_id = $1
`

type GetEventsByStoreRow struct {
	ID              int64        `json:"id"`
	Name            string       `json:"name"`
	Photo           string       `json:"photo"`
	VoucherQuantity int32        `json:"voucher_quantity"`
	Status          EventsStatus `json:"status"`
	StartTime       time.Time    `json:"start_time"`
	EndTime         time.Time    `json:"end_time"`
	GameType        string       `json:"game_type"`
}

func (q *Queries) GetEventsByStore(ctx context.Context, storeID int64) ([]GetEventsByStoreRow, error) {
	rows, err := q.db.Query(ctx, getEventsByStore, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetEventsByStoreRow{}
	for rows.Next() {
		var i GetEventsByStoreRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Photo,
			&i.VoucherQuantity,
			&i.Status,
			&i.StartTime,
			&i.EndTime,
			&i.GameType,
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

const getRecentPartners = `-- name: GetRecentPartners :many

SELECT 
    username AS username,      -- Partner's username
    full_name AS full_name,    -- Partner's full name
    email AS email,            -- Partner's email
    photo AS photo             -- Partner's profile photo
FROM users
WHERE role = 'partner'         -- Only include users with the "partner" role
ORDER BY created_at DESC       -- Sort by creation date (newest first)
LIMIT 5
`

type GetRecentPartnersRow struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Photo    string `json:"photo"`
}

// Description:
// This query retrieves the 5 most recently created partners from the "users" table.
// Only users with the role "partner" are considered.
// The output includes the username, full name, email, and photo of each partner.
func (q *Queries) GetRecentPartners(ctx context.Context) ([]GetRecentPartnersRow, error) {
	rows, err := q.db.Query(ctx, getRecentPartners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRecentPartnersRow{}
	for rows.Next() {
		var i GetRecentPartnersRow
		if err := rows.Scan(
			&i.Username,
			&i.FullName,
			&i.Email,
			&i.Photo,
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

const getRecentUsers = `-- name: GetRecentUsers :many
SELECT 
    u.username,
    u.full_name,
    u.email,
    u.photo,
    COUNT(vo.voucher_id) AS vouchers
FROM voucher_owner vo
JOIN users u ON vo.username = u.username
GROUP BY u.username, u.full_name, u.email, u.photo
ORDER BY MAX(vo.created_at) DESC
LIMIT 5
`

type GetRecentUsersRow struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Photo    string `json:"photo"`
	Vouchers int64  `json:"vouchers"`
}

func (q *Queries) GetRecentUsers(ctx context.Context) ([]GetRecentUsersRow, error) {
	rows, err := q.db.Query(ctx, getRecentUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRecentUsersRow{}
	for rows.Next() {
		var i GetRecentUsersRow
		if err := rows.Scan(
			&i.Username,
			&i.FullName,
			&i.Email,
			&i.Photo,
			&i.Vouchers,
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

const getRecentVoucherOwners = `-- name: GetRecentVoucherOwners :many
SELECT 
    vo.username,
    u.full_name,
    u.email,
    u.photo,
    COUNT(vo.voucher_id) AS vouchers_received
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN events e ON v.event_id = e.id
JOIN stores s ON e.store_id = s.id
JOIN users u ON vo.username = u.username
WHERE s.owner = $1
  AND vo.created_at >= NOW() - INTERVAL '1 MONTH'
GROUP BY vo.username, u.full_name, u.email, u.photo
ORDER BY MAX(vo.created_at) DESC
LIMIT 5
`

type GetRecentVoucherOwnersRow struct {
	Username         string `json:"username"`
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	Photo            string `json:"photo"`
	VouchersReceived int64  `json:"vouchers_received"`
}

func (q *Queries) GetRecentVoucherOwners(ctx context.Context, owner string) ([]GetRecentVoucherOwnersRow, error) {
	rows, err := q.db.Query(ctx, getRecentVoucherOwners, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRecentVoucherOwnersRow{}
	for rows.Next() {
		var i GetRecentVoucherOwnersRow
		if err := rows.Scan(
			&i.Username,
			&i.FullName,
			&i.Email,
			&i.Photo,
			&i.VouchersReceived,
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

const getStoresByOwner = `-- name: GetStoresByOwner :many
SELECT id, name, business_type
FROM stores
WHERE owner = $1
`

type GetStoresByOwnerRow struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	BusinessType string `json:"business_type"`
}

func (q *Queries) GetStoresByOwner(ctx context.Context, owner string) ([]GetStoresByOwnerRow, error) {
	rows, err := q.db.Query(ctx, getStoresByOwner, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetStoresByOwnerRow{}
	for rows.Next() {
		var i GetStoresByOwnerRow
		if err := rows.Scan(&i.ID, &i.Name, &i.BusinessType); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT username, hashed_password, full_name, email, role, active
FROM users
WHERE username = $1
`

type GetUserByUsernameRow struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	Active         bool   `json:"active"`
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Role,
		&i.Active,
	)
	return i, err
}

const getUserPlayCountByPartner = `-- name: GetUserPlayCountByPartner :many
    -- (SELECT COUNT(*) FROM branchs WHERE created_at >= NOW() - INTERVAL '1 MONTH') AS total_branch_last_month;


SELECT 
    u.username AS partner_username,    -- Username của partner
    COUNT(DISTINCT vo.username) AS total_user_play -- Tổng số user chơi game của partner
FROM users u
LEFT JOIN stores s ON s.owner = u.username
LEFT JOIN events e ON e.store_id = s.id
LEFT JOIN vouchers v ON v.event_id = e.id
LEFT JOIN voucher_owner vo ON vo.voucher_id = v.id
WHERE u.role = 'partner' -- Chỉ tính user là partner
GROUP BY u.username
ORDER BY total_user_play DESC
`

type GetUserPlayCountByPartnerRow struct {
	PartnerUsername string `json:"partner_username"`
	TotalUserPlay   int64  `json:"total_user_play"`
}

// todo: thêm created_at vào mỗi Table (bắt buộc)
// Description:
// Đầu tiên lấy tất cả user có role là "partner" trong table "users", ví dụ được 3 partner có user name là: ["partner_01", "partner_02", "partner_03"]
// Tiếp theo với từng Partner ở trên, ví dụ partner_01, sẽ tìm tất cả events có event.owner = username = "partner_01"
// Từ các events này tiếp tục query tất cả "vouchers" có voucher.event_id bằng event id trên
// Từ các vouchers này tiếp tục query "voucher_owner" tất cả có voucher_owner.voucher_id bằng voucher id trên
// Tổng hợp list voucher_owner của các voucher từ tất cả events trên, ta distinct theo voucher_owner.username, sẽ ra được số lượng user chơi game nhận được voucher groupby Partner.
func (q *Queries) GetUserPlayCountByPartner(ctx context.Context) ([]GetUserPlayCountByPartnerRow, error) {
	rows, err := q.db.Query(ctx, getUserPlayCountByPartner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserPlayCountByPartnerRow{}
	for rows.Next() {
		var i GetUserPlayCountByPartnerRow
		if err := rows.Scan(&i.PartnerUsername, &i.TotalUserPlay); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserPlayStats = `-- name: GetUserPlayStats :many
SELECT 
    TO_CHAR(DATE_TRUNC('month', vo.created_at), 'Month') AS month,
    COUNT(CASE WHEN g.type = 'quiz' THEN 1 END) AS quiz_game,
    COUNT(CASE WHEN g.type = 'phone-shake' THEN 1 END) AS shake_game
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN events e ON v.event_id = e.id
JOIN games g ON e.game_id = g.id
GROUP BY DATE_TRUNC('month', vo.created_at)
ORDER BY DATE_TRUNC('month', vo.created_at)
`

type GetUserPlayStatsRow struct {
	Month     string `json:"month"`
	QuizGame  int64  `json:"quiz_game"`
	ShakeGame int64  `json:"shake_game"`
}

func (q *Queries) GetUserPlayStats(ctx context.Context) ([]GetUserPlayStatsRow, error) {
	rows, err := q.db.Query(ctx, getUserPlayStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserPlayStatsRow{}
	for rows.Next() {
		var i GetUserPlayStatsRow
		if err := rows.Scan(&i.Month, &i.QuizGame, &i.ShakeGame); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserPlayStatsByGameAndDate = `-- name: GetUserPlayStatsByGameAndDate :many


WITH partner_events AS (
    SELECT e.id AS event_id, e.game_id
    FROM events e
    WHERE e.owner = $1 -- Replace $1 with the input partner username
),
partner_vouchers AS (
    SELECT v.id AS voucher_id, pe.game_id
    FROM vouchers v
    JOIN partner_events pe ON v.event_id = pe.event_id
),
user_play_stats AS (
    SELECT 
        vo.created_at::date AS play_date, 
        pv.game_id, 
        COUNT(DISTINCT vo.username) AS total_users
    FROM voucher_owner vo
    JOIN partner_vouchers pv ON vo.voucher_id = pv.voucher_id
    GROUP BY vo.created_at::date, pv.game_id
)
SELECT 
    play_date,
    CASE 
        WHEN game_id = 1 THEN 'quizGame'
        WHEN game_id = 2 THEN 'shakeGame'
        ELSE 'unknownGame'
    END AS game_type,
    total_users
FROM user_play_stats
ORDER BY play_date, game_type
`

type GetUserPlayStatsByGameAndDateRow struct {
	PlayDate   pgtype.Date `json:"play_date"`
	GameType   string      `json:"game_type"`
	TotalUsers int64       `json:"total_users"`
}

// Description: Count số lượng user chơi game quizGame và shakeGame theo từng ngày của 1 Partner.
// input: partner username, var owner = "partner_01"
// flow:
//
//	Từ owner, tìm tất cả events có event.owner = owner = "partner_01", groupby "event.game_id" (game_id 1 là quizGame, game_id 2 là shakeGame)
//	Với mỗi list events group by game_id, ví dụ từ list events có game_id = 1 (quizGame), tiếp tục query tất cả "vouchers" có voucher.event_id nằm trong list events id trên, kết quả sẽ ra được list vouchers của game_id =1 thuộc owner="partner_01). Tương tự list events của các game_id còn lại
//	Từ list vouchers đã query (của từng game_id) tiếp tục query table "voucher_owner" với voucher_owner.voucher_id nằm trong list vouchers id trên, group by từng ngày bởi field "voucher_owner.created_at", định dạng "YYYY-MM-dd". Kết quả sẽ ra được list voucher_owner thuộc từng thể loại game_id theo từng ngày
//	Tổng hợp tất cả  voucher_owner từ các vouchers trên, distinct by username, sẽ ra được số lượng user chơi game_id đó theo từng ngày
func (q *Queries) GetUserPlayStatsByGameAndDate(ctx context.Context, owner string) ([]GetUserPlayStatsByGameAndDateRow, error) {
	rows, err := q.db.Query(ctx, getUserPlayStatsByGameAndDate, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserPlayStatsByGameAndDateRow{}
	for rows.Next() {
		var i GetUserPlayStatsByGameAndDateRow
		if err := rows.Scan(&i.PlayDate, &i.GameType, &i.TotalUsers); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserStatsByStore = `-- name: GetUserStatsByStore :many
SELECT 
    s.id AS store_id,
    s.name AS store_name,
    COUNT(DISTINCT vo.username) AS total_users
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN events e ON v.event_id = e.id
JOIN stores s ON e.store_id = s.id
WHERE e.owner = $1
  AND vo.created_at >= NOW() - INTERVAL '6 MONTHS'
GROUP BY s.id, s.name
ORDER BY total_users DESC
`

type GetUserStatsByStoreRow struct {
	StoreID    int64  `json:"store_id"`
	StoreName  string `json:"store_name"`
	TotalUsers int64  `json:"total_users"`
}

func (q *Queries) GetUserStatsByStore(ctx context.Context, owner string) ([]GetUserStatsByStoreRow, error) {
	rows, err := q.db.Query(ctx, getUserStatsByStore, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserStatsByStoreRow{}
	for rows.Next() {
		var i GetUserStatsByStoreRow
		if err := rows.Scan(&i.StoreID, &i.StoreName, &i.TotalUsers); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserStoreStats = `-- name: GetUserStoreStats :many
SELECT 
    s.owner AS username,
    COUNT(DISTINCT vo.username) AS total_user_play
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN events e ON v.event_id = e.id
JOIN stores s ON e.store_id = s.id
GROUP BY s.owner
`

type GetUserStoreStatsRow struct {
	Username      string `json:"username"`
	TotalUserPlay int64  `json:"total_user_play"`
}

func (q *Queries) GetUserStoreStats(ctx context.Context) ([]GetUserStoreStatsRow, error) {
	rows, err := q.db.Query(ctx, getUserStoreStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserStoreStatsRow{}
	for rows.Next() {
		var i GetUserStoreStatsRow
		if err := rows.Scan(&i.Username, &i.TotalUserPlay); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVoucherOwnersByVoucher = `-- name: GetVoucherOwnersByVoucher :many
SELECT vo.username, u.full_name, u.email, vo.created_at
FROM voucher_owner vo
JOIN users u ON vo.username = u.username
WHERE vo.voucher_id = $1
`

type GetVoucherOwnersByVoucherRow struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) GetVoucherOwnersByVoucher(ctx context.Context, voucherID int64) ([]GetVoucherOwnersByVoucherRow, error) {
	rows, err := q.db.Query(ctx, getVoucherOwnersByVoucher, voucherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetVoucherOwnersByVoucherRow{}
	for rows.Next() {
		var i GetVoucherOwnersByVoucherRow
		if err := rows.Scan(
			&i.Username,
			&i.FullName,
			&i.Email,
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

const getVoucherStatsByMonth = `-- name: GetVoucherStatsByMonth :many
SELECT 
    TO_CHAR(DATE_TRUNC('month', vo.created_at), 'YYYY-MM') AS month,
    g.type AS game_type,
    COUNT(*) AS total_vouchers
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN events e ON v.event_id = e.id
JOIN games g ON e.game_id = g.id
WHERE e.owner = $1
  AND vo.created_at >= NOW() - INTERVAL '6 MONTHS'
GROUP BY month, game_type
ORDER BY month
`

type GetVoucherStatsByMonthRow struct {
	Month         string `json:"month"`
	GameType      string `json:"game_type"`
	TotalVouchers int64  `json:"total_vouchers"`
}

func (q *Queries) GetVoucherStatsByMonth(ctx context.Context, owner string) ([]GetVoucherStatsByMonthRow, error) {
	rows, err := q.db.Query(ctx, getVoucherStatsByMonth, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetVoucherStatsByMonthRow{}
	for rows.Next() {
		var i GetVoucherStatsByMonthRow
		if err := rows.Scan(&i.Month, &i.GameType, &i.TotalVouchers); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVouchersByEvent = `-- name: GetVouchersByEvent :many
SELECT v.id, v.qr_code, v.type, v.status, v.expires_at
FROM vouchers v
WHERE v.event_id = $1
`

type GetVouchersByEventRow struct {
	ID        int64       `json:"id"`
	QrCode    pgtype.Text `json:"qr_code"`
	Type      string      `json:"type"`
	Status    string      `json:"status"`
	ExpiresAt time.Time   `json:"expires_at"`
}

func (q *Queries) GetVouchersByEvent(ctx context.Context, eventID int64) ([]GetVouchersByEventRow, error) {
	rows, err := q.db.Query(ctx, getVouchersByEvent, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetVouchersByEventRow{}
	for rows.Next() {
		var i GetVouchersByEventRow
		if err := rows.Scan(
			&i.ID,
			&i.QrCode,
			&i.Type,
			&i.Status,
			&i.ExpiresAt,
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
