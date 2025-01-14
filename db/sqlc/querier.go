// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CheckEventExists(ctx context.Context, arg CheckEventExistsParams) (int64, error)
	CheckGameExists(ctx context.Context, type_ string) (int64, error)
	CheckStoreExists(ctx context.Context, arg CheckStoreExistsParams) (interface{}, error)
	CheckVoucherExists(ctx context.Context, arg CheckVoucherExistsParams) (int64, error)
	CheckVoucherOwnerExists(ctx context.Context, arg CheckVoucherOwnerExistsParams) (int64, error)
	CountVoucherOwners(ctx context.Context) (int64, error)
	CreateBranch(ctx context.Context, arg CreateBranchParams) (Branch, error)
	CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error)
	CreateFakeBranch(ctx context.Context, arg CreateFakeBranchParams) (int64, error)
	CreateFakeEvent(ctx context.Context, arg CreateFakeEventParams) (int64, error)
	CreateFakeStore(ctx context.Context, arg CreateFakeStoreParams) (int64, error)
	CreateFakeUser(ctx context.Context, arg CreateFakeUserParams) (string, error)
	CreateGame(ctx context.Context, arg CreateGameParams) (int64, error)
	CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	CreateVoucher(ctx context.Context, arg CreateVoucherParams) (int64, error)
	CreateVoucherOwner(ctx context.Context, arg CreateVoucherOwnerParams) (int64, error)
	DeleteBranch(ctx context.Context, id int64) error
	DeleteStore(ctx context.Context, arg DeleteStoreParams) error
	GetAdminStats(ctx context.Context) (GetAdminStatsRow, error)
	GetBranchesByStore(ctx context.Context, storeID int64) ([]GetBranchesByStoreRow, error)
	GetCmsOverview(ctx context.Context, owner string) (GetCmsOverviewRow, error)
	GetEventById(ctx context.Context, id int64) (Event, error)
	GetEventByIdAndOwner(ctx context.Context, arg GetEventByIdAndOwnerParams) (Event, error)
	GetEventCreatedStats(ctx context.Context) ([]GetEventCreatedStatsRow, error)
	GetEventsByStore(ctx context.Context, storeID int64) ([]GetEventsByStoreRow, error)
	GetQuizzesByEventId(ctx context.Context, eventID int64) (Quiz, error)
	GetRecentUsers(ctx context.Context) ([]GetRecentUsersRow, error)
	GetRecentVoucherOwners(ctx context.Context, owner string) ([]GetRecentVoucherOwnersRow, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetStoreByIdAndOwner(ctx context.Context, arg GetStoreByIdAndOwnerParams) (Store, error)
	GetStoresByOwner(ctx context.Context, owner string) ([]GetStoresByOwnerRow, error)
	GetUser(ctx context.Context, arg GetUserParams) (User, error)
	GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error)
	GetUserPlayByDate(ctx context.Context, owner string) ([]GetUserPlayByDateRow, error)
	GetUserPlayStats(ctx context.Context) ([]GetUserPlayStatsRow, error)
	GetUserStatsByStore(ctx context.Context, owner string) ([]GetUserStatsByStoreRow, error)
	GetUserStoreStats(ctx context.Context) ([]GetUserStoreStatsRow, error)
	GetVoucherOwnersByVoucher(ctx context.Context, voucherID int64) ([]GetVoucherOwnersByVoucherRow, error)
	GetVoucherStatsByMonth(ctx context.Context, owner string) ([]GetVoucherStatsByMonthRow, error)
	GetVouchersByEvent(ctx context.Context, eventID int64) ([]GetVouchersByEventRow, error)
	ListBranchs(ctx context.Context, storeID int64) ([]Branch, error)
	ListEvents(ctx context.Context) ([]ListEventsRow, error)
	ListEventsOfOwner(ctx context.Context, owner string) ([]ListEventsOfOwnerRow, error)
	ListStoresOfOwner(ctx context.Context, owner string) ([]Store, error)
	UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error)
	UpdateStore(ctx context.Context, arg UpdateStoreParams) (Store, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
