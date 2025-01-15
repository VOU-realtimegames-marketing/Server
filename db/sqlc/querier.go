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
	CheckStoreExists(ctx context.Context, arg CheckStoreExistsParams) (int64, error)
	CheckVoucherExists(ctx context.Context, arg CheckVoucherExistsParams) (int64, error)
	CheckVoucherOwnerExists(ctx context.Context, arg CheckVoucherOwnerExistsParams) (int64, error)
	CountVoucherOwners(ctx context.Context) (int64, error)
	CreateBranch(ctx context.Context, arg CreateBranchParams) (Branch, error)
	CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error)
	CreateFakeBranch(ctx context.Context, arg CreateFakeBranchParams) (int64, error)
	CreateFakeEvent(ctx context.Context, arg CreateFakeEventParams) (int64, error)
	CreateFakeStore(ctx context.Context, arg CreateFakeStoreParams) (int64, error)
	CreateFakeUser(ctx context.Context, arg CreateFakeUserParams) (string, error)
	CreateFakeVoucher(ctx context.Context, arg CreateFakeVoucherParams) (int64, error)
	CreateFakeVoucherOwner(ctx context.Context, arg CreateFakeVoucherOwnerParams) (int64, error)
	CreateGame(ctx context.Context, arg CreateGameParams) (int64, error)
	CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserAnswer(ctx context.Context, arg CreateUserAnswerParams) (UserAnswer, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	CreateVoucher(ctx context.Context, arg CreateVoucherParams) (Voucher, error)
	CreateVoucherOwner(ctx context.Context, arg CreateVoucherOwnerParams) (VoucherOwner, error)
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
	// Description:
	// This query retrieves the 5 most recently created partners from the "users" table.
	// Only users with the role "partner" are considered.
	// The output includes the username, full name, email, and photo of each partner.
	GetRecentPartners(ctx context.Context) ([]GetRecentPartnersRow, error)
	GetRecentUsers(ctx context.Context) ([]GetRecentUsersRow, error)
	GetRecentVoucherOwners(ctx context.Context, owner string) ([]GetRecentVoucherOwnersRow, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetStoreByIdAndOwner(ctx context.Context, arg GetStoreByIdAndOwnerParams) (Store, error)
	GetStoresByOwner(ctx context.Context, owner string) ([]GetStoresByOwnerRow, error)
	GetUser(ctx context.Context, arg GetUserParams) (User, error)
	GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error)
	// todo: thêm created_at vào mỗi Table (bắt buộc)
	// Description:
	// Đầu tiên lấy tất cả user có role là "partner" trong table "users", ví dụ được 3 partner có user name là: ["partner_01", "partner_02", "partner_03"]
	// Tiếp theo với từng Partner ở trên, ví dụ partner_01, sẽ tìm tất cả events có event.owner = username = "partner_01"
	// Từ các events này tiếp tục query tất cả "vouchers" có voucher.event_id bằng event id trên
	// Từ các vouchers này tiếp tục query "voucher_owner" tất cả có voucher_owner.voucher_id bằng voucher id trên
	// Tổng hợp list voucher_owner của các voucher từ tất cả events trên, ta distinct theo voucher_owner.username, sẽ ra được số lượng user chơi game nhận được voucher groupby Partner.
	GetUserPlayCountByPartner(ctx context.Context) ([]GetUserPlayCountByPartnerRow, error)
	GetUserPlayStats(ctx context.Context) ([]GetUserPlayStatsRow, error)
	// Description: Count số lượng user chơi game quizGame và shakeGame theo từng ngày của 1 Partner.
	// input: partner username, var owner = "partner_01"
	// flow:
	//    Từ owner, tìm tất cả events có event.owner = owner = "partner_01", groupby "event.game_id" (game_id 1 là quizGame, game_id 2 là shakeGame)
	//    Với mỗi list events group by game_id, ví dụ từ list events có game_id = 1 (quizGame), tiếp tục query tất cả "vouchers" có voucher.event_id nằm trong list events id trên, kết quả sẽ ra được list vouchers của game_id =1 thuộc owner="partner_01). Tương tự list events của các game_id còn lại
	//    Từ list vouchers đã query (của từng game_id) tiếp tục query table "voucher_owner" với voucher_owner.voucher_id nằm trong list vouchers id trên, group by từng ngày bởi field "voucher_owner.created_at", định dạng "YYYY-MM-dd". Kết quả sẽ ra được list voucher_owner thuộc từng thể loại game_id theo từng ngày
	//    Tổng hợp tất cả  voucher_owner từ các vouchers trên, distinct by username, sẽ ra được số lượng user chơi game_id đó theo từng ngày
	GetUserPlayStatsByGameAndDate(ctx context.Context, owner string) ([]GetUserPlayStatsByGameAndDateRow, error)
	GetUserStatsByStore(ctx context.Context, owner string) ([]GetUserStatsByStoreRow, error)
	GetUserStoreStats(ctx context.Context) ([]GetUserStoreStatsRow, error)
	GetVoucherOwnersByVoucher(ctx context.Context, voucherID int64) ([]GetVoucherOwnersByVoucherRow, error)
	GetVoucherStatsByMonth(ctx context.Context, owner string) ([]GetVoucherStatsByMonthRow, error)
	GetVouchersByEvent(ctx context.Context, eventID int64) ([]GetVouchersByEventRow, error)
	ListAllOtherUsers(ctx context.Context, username string) ([]User, error)
	ListBranchs(ctx context.Context, storeID int64) ([]Branch, error)
	ListEvents(ctx context.Context) ([]ListEventsRow, error)
	ListEventsOfOwner(ctx context.Context, owner string) ([]ListEventsOfOwnerRow, error)
	ListMyVouchers(ctx context.Context, username string) ([]Voucher, error)
	ListStoresOfOwner(ctx context.Context, owner string) ([]Store, error)
	UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error)
	UpdateStore(ctx context.Context, arg UpdateStoreParams) (Store, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserAnswer(ctx context.Context, arg UpdateUserAnswerParams) (UserAnswer, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
