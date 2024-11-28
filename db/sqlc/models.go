// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Branch struct {
	ID       int64  `json:"id"`
	StoreID  int64  `json:"store_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Event struct {
	ID              int64     `json:"id"`
	GameID          int64     `json:"game_id"`
	Name            string    `json:"name"`
	Photo           string    `json:"photo"`
	VoucherQuantity int32     `json:"voucher_quantity"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
}

type Game struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Photo       string      `json:"photo"`
	Type        string      `json:"type"`
	PlayGuide   pgtype.Text `json:"play_guide"`
	GiftAllowed bool        `json:"gift_allowed"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Store struct {
	ID           int64  `json:"id"`
	Owner        string `json:"owner"`
	Name         string `json:"name"`
	BusinessType string `json:"business_type"`
}

type TransferVoucher struct {
	ID        int64     `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	CreatedAt time.Time `json:"created_at"`
	VoucherID int64     `json:"voucher_id"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	Role              string    `json:"role"`
	Photo             string    `json:"photo"`
	Active            bool      `json:"active"`
	IsEmailVerified   bool      `json:"is_email_verified"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type VerifyEmail struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}

type Voucher struct {
	ID        int64       `json:"id"`
	EventID   int64       `json:"event_id"`
	QrCode    pgtype.Text `json:"qr_code"`
	Type      string      `json:"type"`
	Status    string      `json:"status"`
	ExpiresAt time.Time   `json:"expires_at"`
}

type VoucherOwner struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	VoucherID int64     `json:"voucher_id"`
	CreatedAt time.Time `json:"created_at"`
}
