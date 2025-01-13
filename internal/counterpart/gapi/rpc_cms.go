package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetCmsOverview(ctx context.Context, req *gen.GetCmsOverviewRequest) (*gen.GetCmsOverviewResponse, error) {
	owner := req.GetOwner()
	log.Print("GetCmsOverview_rpc_cms owner:", owner)

	// Query tổng quan
	cmsData, err := server.store.GetCmsOverview(ctx, owner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get CMS overview: %s", err)
	}

	// Query chi tiết user chơi game
	userPlayData, err := server.store.GetUserPlayByDate(ctx, owner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user play data: %s", err)
	}

	// Query danh sách user mới nhận voucher
	recentUsers, err := server.store.GetRecentVoucherOwners(ctx, owner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get recent users: %s", err)
	}

	// Query thống kê voucher theo tháng
	voucherStats, err := server.store.GetVoucherStatsByMonth(ctx, owner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get voucher stats: %s", err)
	}

	// Query user theo store
	userStoreStats, err := server.store.GetUserStatsByStore(ctx, owner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user store stats: %s", err)
	}

	// Map dữ liệu trả về FE
	response := &gen.GetCmsOverviewResponse{
		TotalStore:             int32(cmsData.TotalStore),
		TotalBranch:            int32(cmsData.TotalBranch),
		TotalEvent:             int32(cmsData.TotalEvent),
		TotalUserPlay:          int32(cmsData.TotalUserPlay),
		LastMonthTotalUserPlay: int32(cmsData.LastMonthTotalUserPlay),
		ChartUserPlay:          mapUserPlayData(userPlayData),
		ListRecentUser:         mapRecentUsers(recentUsers),
		ChartVoucher:           mapVoucherStats(voucherStats),
		ChartUserStore:         mapUserStoreStats(userStoreStats),
	}

	return response, nil
}

func mapUserPlayData(data []db.GetUserPlayByDateRow) []*gen.UserPlayData {
	result := []*gen.UserPlayData{}
	for _, d := range data {
		// Kiểm tra giá trị hợp lệ trong PlayDate
		if !d.PlayDate.Valid {
			log.Printf("Invalid PlayDate for record: %+v", d)
			continue
		}

		// Lấy timestamp từ PlayDate
		timestamp := d.PlayDate.Time.Unix()

		// Phân loại game_type thành quizGame và shakeGame
		if d.GameType == "quizGame" {
			result = append(result, &gen.UserPlayData{
				Date:      timestamp,
				QuizGame:  int32(d.TotalUsers),
				ShakeGame: 0,
			})
		} else if d.GameType == "shakeGame" {
			result = append(result, &gen.UserPlayData{
				Date:      timestamp,
				QuizGame:  0,
				ShakeGame: int32(d.TotalUsers),
			})
		}
	}
	return result
}

func mapRecentUsers(data []db.GetRecentVoucherOwnersRow) []*gen.RecentUser {
	result := []*gen.RecentUser{}
	for _, u := range data {
		result = append(result, &gen.RecentUser{
			Username: u.Username,
			FullName: u.FullName,
			Email:    u.Email,
			Photo:    u.Photo,
			Vouchers: int32(u.VouchersReceived),
		})
	}
	return result
}

func mapVoucherStats(data []db.GetVoucherStatsByMonthRow) []*gen.VoucherStats {
	result := []*gen.VoucherStats{}
	for _, v := range data {
		// Tìm tháng hiện tại trong result
		found := false
		for _, existing := range result {
			if existing.Month == v.Month {
				// Gộp dữ liệu theo game_type
				if v.GameType == "shakeGame" {
					existing.ShakeGame += int32(v.TotalVouchers)
				} else if v.GameType == "quizGame" {
					existing.QuizGame += int32(v.TotalVouchers)
				}
				found = true
				break
			}
		}

		// Nếu tháng chưa tồn tại, thêm mới
		if !found {
			result = append(result, &gen.VoucherStats{
				Month:     v.Month,
				QuizGame:  0,
				ShakeGame: 0,
			})

			// Thêm dữ liệu mới
			if v.GameType == "shakeGame" {
				result[len(result)-1].ShakeGame = int32(v.TotalVouchers)
			} else if v.GameType == "quizGame" {
				result[len(result)-1].QuizGame = int32(v.TotalVouchers)
			}
		}
	}
	return result
}

func mapUserStoreStats(data []db.GetUserStatsByStoreRow) []*gen.UserStoreStats {
	result := []*gen.UserStoreStats{}
	for _, s := range data {
		result = append(result, &gen.UserStoreStats{
			StoreId:       s.StoreID,
			StoreName:     s.StoreName,
			TotalUserPlay: int32(s.TotalUsers),
		})
	}
	return result
}
