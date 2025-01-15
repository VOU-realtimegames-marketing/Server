package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/pkg/utils"
	"VOU-Server/proto/gen"
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetPartnerCmsOverview(ctx context.Context, req *gen.GetPartnerCmsOverviewRequest) (*gen.GetPartnerCmsOverviewResponse, error) {
	owner := req.GetOwner()
	log.Print("GetCmsOverview_rpc_cms owner:", owner)

	// Query tổng quan
	cmsData, err := server.store.GetCmsOverview(ctx, owner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get CMS overview: %s", err)
	}

	// Query số lượng user chơi từng game theo ngày
	userPlayData, err := server.store.GetUserPlayStatsByGameAndDate(ctx, owner)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user play data: %s", err)
	}

	// Map
	chartUserPlay := mapUserPlayData(userPlayData)

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
	response := &gen.GetPartnerCmsOverviewResponse{
		TotalStore:             int32(cmsData.TotalStore),
		TotalBranch:            int32(cmsData.TotalBranch),
		TotalEvent:             int32(cmsData.TotalEvent),
		TotalUserPlay:          int32(cmsData.TotalUserPlay),
		LastMonthTotalUserPlay: int32(cmsData.LastMonthTotalUserPlay),
		ChartUserPlay:          chartUserPlay,
		ListRecentUser:         mapRecentUsers(recentUsers),
		ChartVoucher:           mapVoucherStats(voucherStats),
		ChartUserStore:         mapUserStoreStats(userStoreStats),
	}

	return response, nil
}
func mapUserPlayData(userPlayData []db.GetUserPlayStatsByGameAndDateRow) []*gen.UserPlayData {
	// Sử dụng map để gom nhóm dữ liệu theo ngày
	resultMap := make(map[int64]*gen.UserPlayData)

	for _, stat := range userPlayData {
		// Lấy timestamp từ stat.Date
		dateTimestamp := stat.PlayDate.Time.Unix()

		// Nếu ngày chưa tồn tại trong map, khởi tạo entry mới
		if _, exists := resultMap[dateTimestamp]; !exists {
			resultMap[dateTimestamp] = &gen.UserPlayData{
				Date:      dateTimestamp,
				QuizGame:  0,
				ShakeGame: 0,
			}
		}

		// Cộng số lượng user vào game tương ứng
		if stat.GameType == "quizGame" {
			resultMap[dateTimestamp].QuizGame += int32(stat.TotalUsers) // Chuyển đổi sang int32
		} else if stat.GameType == "shakeGame" {
			resultMap[dateTimestamp].ShakeGame += int32(stat.TotalUsers) // Chuyển đổi sang int32
		}
	}

	// Chuyển map thành slice để trả về
	chartData := make([]*gen.UserPlayData, 0, len(resultMap))
	for _, data := range resultMap {
		chartData = append(chartData, data)
	}

	return chartData
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
	result := make(map[string]*gen.VoucherStats) // Dùng map để dễ dàng gộp theo tháng

	for _, v := range data {
		// Kiểm tra nếu tháng đã tồn tại trong kết quả
		if _, exists := result[v.Month]; !exists {
			result[v.Month] = &gen.VoucherStats{
				Month:     v.Month,
				QuizGame:  0,
				ShakeGame: 0,
			}
		}

		// Cập nhật dữ liệu theo loại game
		if v.GameType == "phone-shake" {
			result[v.Month].ShakeGame += int32(v.TotalVouchers)
		} else if v.GameType == "quiz" {
			result[v.Month].QuizGame += int32(v.TotalVouchers)
		}
	}

	// Chuyển kết quả từ map sang slice
	finalResult := make([]*gen.VoucherStats, 0, len(result))
	for _, stats := range result {
		finalResult = append(finalResult, stats)
	}

	return finalResult
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

func (server *Server) GetAdminCmsOverview(ctx context.Context, req *gen.GetAdminCmsOverviewRequest) (*gen.GetAdminCmsOverviewResponse, error) {
	log.Println("Starting GetAdminCmsOverview")

	// Fetch admin stats
	adminStats, err := server.store.GetAdminStats(ctx)
	if err != nil {
		log.Printf("Failed to fetch admin stats: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to fetch admin stats: %v", err)
	}

	log.Println("\n_________adminStats: ", adminStats)

	// Fetch event creation stats (bar chart)
	eventCreatedStats, err := server.store.GetEventCreatedStats(ctx)
	if err != nil {
		log.Printf("Failed to fetch event creation stats: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to fetch event creation stats: %v", err)
	}

	// // Map event creation stats
	chartEventCreated := []*gen.UserPlayData{}
	for _, stat := range eventCreatedStats {
		chartEventCreated = append(chartEventCreated, &gen.UserPlayData{
			Date:      stat.Date.Time.Unix(),
			QuizGame:  int32(stat.QuizGame),
			ShakeGame: int32(stat.ShakeGame),
		})
	}
	log.Println("\n_________chartEventCreated: ", chartEventCreated)

	// Fetch user play stats (area chart)
	userPlayStats, err := server.store.GetUserPlayStats(ctx)
	if err != nil {
		log.Printf("Failed to fetch user play stats: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to fetch user play stats: %v", err)
	}

	// Map user play stats
	chartUserPlayGame := []*gen.VoucherStats{}
	for _, stat := range userPlayStats {
		chartUserPlayGame = append(chartUserPlayGame, &gen.VoucherStats{
			Month:     stat.Month,
			QuizGame:  int32(stat.QuizGame),
			ShakeGame: int32(stat.ShakeGame),
		})
	}
	log.Println("\n_________chartUserPlayGame: ", chartUserPlayGame)

	// Fetch user stats by Partner (pie chart)
	userStoreStats, err := server.store.GetUserPlayCountByPartner(ctx)
	if err != nil {
		log.Printf("Failed to fetch user store stats: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to fetch user store stats: %v", err)
	}
	log.Println("\n_________userStoreStats: ", userStoreStats)

	// Map user stats by Partner
	chartUserPlayGroupByPartner := []*gen.AdminUserPlayGroupByPartnerStats{}
	for _, stat := range userStoreStats {
		chartUserPlayGroupByPartner = append(chartUserPlayGroupByPartner, &gen.AdminUserPlayGroupByPartnerStats{
			Username:      stat.PartnerUsername,
			TotalUserPlay: int32(stat.TotalUserPlay),
		})
	}
	log.Println("\n_________chartUserPlayGroupByPartner: ", chartUserPlayGroupByPartner)
	// Fetch recent partners
	recentPartners, err := server.store.GetRecentPartners(ctx)
	if err != nil {
		log.Printf("Failed to fetch recent partners: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to fetch recent partners: %v", err)
	}

	// Map recent partners
	listRecentPartners := []*gen.RecentUser{}
	for _, partner := range recentPartners {
		listRecentPartners = append(listRecentPartners, &gen.RecentUser{
			Username: partner.Username,
			FullName: partner.FullName,
			Email:    partner.Email,
			Photo:    partner.Photo,
		})
	}

	log.Println("\n_________listRecentPartners: ", listRecentPartners)

	// Construct response
	response := &gen.GetAdminCmsOverviewResponse{
		TotalPartner:                int32(adminStats.TotalPartner),
		TotalPartnerLastMonth:       int32(adminStats.TotalPartnerLastMonth),
		TotalUser:                   int32(adminStats.TotalUser),
		TotalUserLastMonth:          int32(adminStats.TotalUserLastMonth),
		TotalBranch:                 int32(adminStats.TotalBranch),
		TotalBranchLastMonth:        int32(adminStats.TotalPartnerLastMonth),
		TotalEarning:                float64(0.0), // TODO: Xử lý lợi nhuận của VOU
		TotalEarningLastMonth:       float64(0.0), // TODO: Xử lý lợi nhuận của VOU
		ChartEventCreated:           chartEventCreated,
		ChartUserPlayGame:           chartUserPlayGame,
		ChartUserPlayGroupByPartner: chartUserPlayGroupByPartner,
		ListRecentPartners:          listRecentPartners,
	}

	log.Println("GetAdminCmsOverview completed successfully")
	return response, nil
}

func (server *Server) FakeCmsOverview(ctx context.Context, req *gen.FakeCmsOverviewRequest) (*gen.FakeCmsOverviewResponse, error) {
	log.Println("Starting FakeCmsOverview")

	// Create a admin user
	hashedPassword, err := utils.HashPassword("12341234")
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, status.Errorf(codes.Internal, "error hashing password")
	}
	username := "admin_user"
	_, err = server.store.CreateFakeUser(ctx, db.CreateFakeUserParams{
		Username:       username,
		HashedPassword: hashedPassword,
		FullName:       "Vou Admin",
		Email:          "admin_user@gmail.com",
		Role:           "admin",
		Photo:          "https://lh3.googleusercontent.com/a/ACg8ocK_Z9gZipugNC-xdIe1RB6AmUJz6pTQo_jPGE7dEIL-ZwMMn2Ps=s192-c-rg-br100",
		CreatedAt:      time.Now(),
	})
	if err != nil {
		log.Printf("Failed to create partner user: %v", err)
		return nil, status.Errorf(codes.Internal, "error creating partner user")
	}
	log.Println("Admin Created: ")

	// Tạo danh sách Partner và fake data cho Partner đó
	// Danh sách partner
	partners := generatePartners(20)
	for _, partner := range partners {
		hashedPassword, err := utils.HashPassword("12341234")
		if err != nil {
			log.Printf("Failed to hash password: %v", err)
			continue
		}

		_, err = server.store.CreateFakeUser(ctx, db.CreateFakeUserParams{
			Username:       partner.Username,
			HashedPassword: hashedPassword,
			FullName:       partner.FullName,
			Email:          partner.Email,
			Role:           "partner",
			Photo:          partner.Photo,
			CreatedAt:      partner.CreatedAt,
		})
		if err != nil {
			log.Printf("Failed to create partner user %s: %v", partner.Username, err)
			continue
		}
		log.Printf("\n\n\n\n==================== Partner user created: %s", partner)

		server.genAllDataForEachPartner(partner.Username, ctx)
	}

	return &gen.FakeCmsOverviewResponse{}, nil
}

func (server *Server) genAllDataForEachPartner(partnerUsername string, ctx context.Context) {
	log.Println("\n\n======== Start generating data for partner: ", partnerUsername)
	log.Println("======== Start step 2 ========")
	listFakeStores := fakeStores(5)
	log.Println("storeNames: ", listFakeStores)
	storeIDs := make([]int64, 0)
	for _, fakeStore := range listFakeStores {
		// Check if the store already exists
		existingStoreId, err := server.store.CheckStoreExists(ctx, db.CheckStoreExistsParams{
			Name:  fakeStore.Name,
			Owner: partnerUsername,
		})

		log.Println("existingStoreId: ", existingStoreId)
		log.Println("err: ", err)

		if existingStoreId != 0 {
			log.Printf("Store %s for owner %s already exists, skipping.", fakeStore.Name, partnerUsername)
			storeIDs = append(storeIDs, existingStoreId)
			continue
		}

		// Create the store
		storeID, err := server.store.CreateFakeStore(ctx, db.CreateFakeStoreParams{
			Name:         fakeStore.Name,
			Owner:        partnerUsername,
			BusinessType: fakeStore.BusinessType,
		})
		if err != nil {
			log.Printf("Failed to create store %s: %v", fakeStore.Name, err)
			continue
		}
		storeIDs = append(storeIDs, storeID)
		log.Printf("Store created: %s, ID: %d", fakeStore.Name, storeID)
	}

	// Step 3: Create branches for each store
	log.Println("======== Start step 3 ========")
	// branchNames := map[string][]string{
	// 	"Highland Coffee": {"Highland Vincom", "Highland Landmark 81"},
	// 	"Starbucks":       {"Starbucks District 1", "Starbucks District 7"},
	// }
	// branchIDs := make(map[string][]int64)

	// for storeIndex, storeID := range storeIDs {
	// 	storeName := storeNames[storeIndex]
	// 	branches := branchNames[storeName]
	// 	for _, branchName := range branches {
	// 		// Create branch
	// 		branchID, err := server.store.CreateFakeBranch(ctx, db.CreateFakeBranchParams{
	// 			StoreID:  storeID,
	// 			Name:     branchName,
	// 			Position: fmt.Sprintf("Position of %s", branchName),
	// 			CityName: "Ho Chi Minh City",
	// 			Country:  "Vietnam",
	// 			Address:  fmt.Sprintf("%s address", branchName),
	// 			Emoji:    "🏢",
	// 		})
	// 		if err != nil {
	// 			log.Printf("Failed to create branch %s for store %s: %v", branchName, storeName, err)
	// 			continue
	// 		}
	// 		branchIDs[storeName] = append(branchIDs[storeName], branchID)
	// 		log.Printf("Branch created: %s, ID: %d", branchName, branchID)
	// 	}
	// }

	// Step 4: Create games
	log.Println("\n\n======== Start step 4 ========")

	// Danh sách game cần tạo
	gamesToCreate := []struct {
		Name        string
		Type        string
		Photo       string
		PlayGuide   pgtype.Text
		GiftAllowed bool
	}{
		{"Phone Shake", "phone-shake", "default-game.jpg", pgtype.Text{String: "Shake phone to get voucher", Valid: true}, true},
		{"Quiz", "quiz", "default-game.jpg", pgtype.Text{String: "Jump to play the quiz to get voucher", Valid: true}, true},
	}

	gameIDs := make(map[string]int64)

	// Tạo game nếu chưa tồn tại
	for _, game := range gamesToCreate {
		existingGameID, err := server.store.CheckGameExists(ctx, game.Type)
		log.Println("existingGameID: ", existingGameID)

		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("Game %s does not exist, creating it...", game.Name)
			} else {
				log.Printf("Error checking game existence for %s: %v", game.Name, err)
				continue
			}
		} else {
			log.Printf("Game %s already exists with ID: %d", game.Name, existingGameID)
			gameIDs[game.Type] = existingGameID
			continue
		}

		// Tạo game
		newGameID, err := server.store.CreateGame(ctx, db.CreateGameParams{
			Name:        game.Name,
			Type:        game.Type,
			Photo:       game.Photo,
			PlayGuide:   game.PlayGuide,
			GiftAllowed: game.GiftAllowed,
		})
		if err != nil {
			log.Printf("Failed to create game %s: %v", game.Name, err)
			continue
		}
		gameIDs[game.Type] = newGameID
		log.Printf("Game created: %s, ID: %d", game.Name, newGameID)
	}

	log.Printf("Step 4 completed. Created/Existing games: %v", gameIDs)

	// Step 5: Create events
	log.Println("\n\n======== Start step 5 ========")
	// Danh sách sự kiện với thời gian phù hợp
	eventDetails := []struct {
		Name      string
		StartTime time.Time
		EndTime   time.Time
	}{
		{"Trung thu 2024", time.Date(2024, 9, 20, 0, 0, 0, 0, time.UTC), time.Date(2024, 9, 27, 23, 59, 59, 0, time.UTC)},
		{"Black Friday 2024", time.Date(2024, 11, 29, 0, 0, 0, 0, time.UTC), time.Date(2024, 12, 5, 23, 59, 59, 0, time.UTC)},
		{"Giáng sinh 2024", time.Date(2024, 12, 24, 0, 0, 0, 0, time.UTC), time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)},
		{"Tết dương lịch 2025", time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 4, 23, 59, 59, 0, time.UTC)},
		{"Tết âm lịch 2025", time.Date(2025, 1, 29, 0, 0, 0, 0, time.UTC), time.Date(2025, 2, 5, 23, 59, 59, 0, time.UTC)},
	}
	voucherQuantities := []int{10, 20, 50, 100}

	eventIDs := make([]int64, 0)

	for _, storeID := range storeIDs { // Duyệt qua từng store từ Step 2
		for _, eventDetail := range eventDetails {
			for gameType, gameID := range gameIDs { // Gắn mỗi game vào event
				// Kiểm tra xem sự kiện đã tồn tại chưa
				existingEventID, err := server.store.CheckEventExists(ctx, db.CheckEventExistsParams{
					Owner:   partnerUsername,
					GameID:  gameID,
					StoreID: storeID,
					Name:    eventDetail.Name,
				})

				// Log kết quả kiểm tra
				log.Printf("Checking event: StoreID: %d, GameType: %s, EventName: %s", storeID, gameType, eventDetail.Name)
				if existingEventID != 0 {
					log.Printf("Event already exists: %s (ID: %d)", eventDetail.Name, existingEventID)
					eventIDs = append(eventIDs, existingEventID)
					continue
				}

				voucherQuantity := voucherQuantities[rand.Intn(len(voucherQuantities))]

				// Tạo sự kiện
				eventID, err := server.store.CreateFakeEvent(ctx, db.CreateFakeEventParams{
					GameID:          gameID,
					StoreID:         storeID,
					Owner:           partnerUsername,
					Name:            eventDetail.Name,
					Photo:           "default-event.jpg",
					VoucherQuantity: int32(voucherQuantity),
					Status:          "ready",
					StartTime:       eventDetail.StartTime,
					EndTime:         eventDetail.EndTime,
				})
				if err != nil {
					log.Printf("Failed to create event %s for store ID %d: %v", eventDetail.Name, storeID, err)
					continue
				}

				eventIDs = append(eventIDs, eventID)
				log.Printf("Event created: %s for store ID %d, GameType: %s, Event ID: %d", eventDetail.Name, storeID, gameType, eventID)
			}
		}
	}
	log.Println("======== Step 5 Completed ========")

	// Step 6: Create vouchers for each event
	log.Println("\n\n======== Start step 6 ========")
	// Danh sách voucher mẫu
	voucherTypes := []string{"Discount 50%", "Discount 20%", "Discount 10%"}
	voucherIDs := make([]int64, 0)

	for _, eventID := range eventIDs { // Duyệt qua từng event từ Step 5
		for _, voucherType := range voucherTypes {
			// Tạo QR Code giả
			qrCode := fmt.Sprintf("QR-%s-%d", voucherType, eventID)

			qrCodeText := pgtype.Text{
				String: qrCode,
				Valid:  true, // Đảm bảo giá trị này là hợp lệ
			}

			// Kiểm tra voucher đã tồn tại chưa
			existingVoucherID, err := server.store.CheckVoucherExists(ctx, db.CheckVoucherExistsParams{
				EventID: eventID,
				QrCode:  qrCodeText,
				Type:    voucherType,
			})

			log.Println("\nExistingVoucherID: ", existingVoucherID)

			// if err != nil && err != sql.ErrNoRows {
			// 	log.Printf("Error checking voucher existence for event ID %d, QRCode %s: %v", eventID, qrCode, err)
			// 	continue
			// }

			if existingVoucherID != 0 {
				log.Printf("Voucher already exists for event ID %d, QRCode: %s", eventID, qrCode)
				voucherIDs = append(voucherIDs, existingVoucherID)
				continue
			}

			// Tạo thời gian hết hạn voucher
			expiresAt := time.Now().AddDate(0, 1, 0) // Hết hạn sau 1 tháng

			// Tạo voucher
			newVoucherID, err := server.store.CreateFakeVoucher(ctx, db.CreateFakeVoucherParams{
				EventID:   eventID,
				QrCode:    qrCodeText,
				Type:      voucherType,
				Status:    "valid",
				ExpiresAt: expiresAt,
			})
			if err != nil {
				log.Printf("Failed to create voucher %s for event ID %d: %v", voucherType, eventID, err)
				continue
			}
			voucherIDs = append(voucherIDs, newVoucherID)
			log.Printf("Voucher created: %s for event ID %d, Voucher ID: %d", voucherType, eventID, newVoucherID)
		}
	}

	log.Printf("Step 6 completed. Total vouchers created: %d", len(voucherIDs))

	// Step 7: Create fake users
	log.Println("\n\n======== Start step 7 ========")

	// Danh sách users giả
	// fakeUsers := []struct {
	// 	Username string
	// 	FullName string
	// 	Email    string
	// 	Password string
	// }{
	// 	{"fakeUserA", "Fake User A", "fakeA@gmail.com", "12341234"},
	// 	{"fakeUserB", "Fake User B", "fakeB@gmail.com", "12341234"},
	// 	{"fakeUserC", "Fake User C", "fakeC@gmail.com", "12341234"},
	// 	{"nguyenanhA", "Nguyễn Anh A", "nguyenanhA@gmail.com", "12341234"},
	// 	{"lethithaoB", "Lê Thị Thảo B", "lethithaoB@gmail.com", "12341234"},
	// 	{"tranvanC", "Trần Văn C", "tranvanC@gmail.com", "12341234"},
	// 	{"phamminhD", "Phạm Minh D", "phamminhD@gmail.com", "12341234"},
	// 	{"dangthuyE", "Đặng Thúy E", "dangthuyE@gmail.com", "12341234"},
	// 	{"huynhngocF", "Huỳnh Ngọc F", "huynhngocF@gmail.com", "12341234"},
	// 	{"vohuyG", "Võ Huy G", "vohuyG@gmail.com", "12341234"},
	// 	{"doquyenH", "Đỗ Quyên H", "doquyenH@gmail.com", "12341234"},
	// 	{"buiducI", "Bùi Đức I", "buiducI@gmail.com", "12341234"},
	// 	{"hoangmaiJ", "Hoàng Mai J", "hoangmaiJ@gmail.com", "12341234"},
	// }

	fakeUsers := generateVietnameseUsers(50)

	usernames := []string{}

	for _, user := range fakeUsers {
		// Hash password
		hashedPassword, err := utils.HashPassword("12341234")
		if err != nil {
			log.Printf("Failed to hash password for user %s: %v", user.Username, err)
			continue
		}

		// Tạo hoặc cập nhật user
		createdUser, err := server.store.CreateFakeUser(ctx, db.CreateFakeUserParams{
			Username:       user.Username,
			HashedPassword: hashedPassword,
			FullName:       user.FullName,
			Email:          user.Email,
			Role:           "user",
			Photo:          user.Photo,
			CreatedAt:      user.CreatedAt,
		})
		if err != nil {
			log.Printf("Failed to create or update user %s: %v", user.Username, err)
			continue
		}

		usernames = append(usernames, createdUser)
		log.Printf("User created/updated: %s", createdUser)
	}

	log.Println("Step 7 Completed. Users:", usernames)
	// Step 8: Assign vouchers to users (randomly)
	log.Println("\n\n======== Start step 8 ========")

	voucherOwnerIDs := make([]int64, 0)

	// Random number generator
	rand.Seed(time.Now().UnixNano())

	for _, username := range usernames { // Iterate over each user
		// Randomly decide how many vouchers this user will receive (1 to len(voucherIDs))
		numVouchers := rand.Intn(len(voucherIDs)) + 1

		// Shuffle voucherIDs to ensure randomness
		shuffledVoucherIDs := append([]int64{}, voucherIDs...)
		rand.Shuffle(len(shuffledVoucherIDs), func(i, j int) {
			shuffledVoucherIDs[i], shuffledVoucherIDs[j] = shuffledVoucherIDs[j], shuffledVoucherIDs[i]
		})

		// Assign a random subset of vouchers to the user
		for i := 0; i < numVouchers; i++ {
			voucherID := shuffledVoucherIDs[i]

			// Check if the user already owns this voucher
			existingVoucherOwnerID, err := server.store.CheckVoucherOwnerExists(ctx, db.CheckVoucherOwnerExistsParams{
				Username:  username,
				VoucherID: voucherID,
			})

			log.Println("\n-----existingVoucherOwnerID: ", existingVoucherOwnerID)

			// if err != nil && err != sql.ErrNoRows {
			// 	log.Printf("Error checking voucher ownership for User: %s, VoucherID: %d: %v", username, voucherID, err)
			// 	continue
			// }

			if existingVoucherOwnerID != 0 {
				log.Printf("User %s already owns VoucherID: %d, skipping assignment.", username, voucherID)
				continue
			}

			// Generate a random creation time for the voucher assignment
			createdAt := time.Now().Add(-time.Duration(rand.Intn(30)) * 24 * time.Hour) // Randomly assign within the last 30 days

			log.Printf("\n=======Assigning VoucherID: %d to User: %s (CreatedAt: %s)", voucherID, username, createdAt)
			// Assign the voucher to the user
			voucherOwnerID, err := server.store.CreateFakeVoucherOwner(ctx, db.CreateFakeVoucherOwnerParams{
				Username:  username,
				VoucherID: voucherID,
				CreatedAt: createdAt,
			})
			if err != nil {
				log.Printf("Failed to assign VoucherID: %d to User: %s: %v", voucherID, username, err)
				continue
			}

			voucherOwnerIDs = append(voucherOwnerIDs, voucherOwnerID)
			log.Printf("Assigned VoucherID: %d to User: %s (VoucherOwnerID: %d)", voucherID, username, voucherOwnerID)
		}
	}

	log.Printf("Step 8 Completed. Total voucher assignments created: %d \n", len(voucherOwnerIDs))

	// Test step 8
	// Query total voucher_owner count from database
	totalVoucherOwners, err := server.store.CountVoucherOwners(ctx)
	if err != nil {
		log.Printf("Failed to count voucher_owner records: %v", err)
	} else {
		log.Printf("Total voucher_owner records in database: %d", totalVoucherOwners)
	}
}

func generateVietnameseUsers(count int) []struct {
	Username  string
	FullName  string
	Email     string
	Photo     string
	CreatedAt time.Time
} {
	hoList := []string{"Nguyễn", "Trần", "Lê", "Phạm", "Hoàng", "Vũ", "Võ", "Đặng", "Bùi", "Đỗ"}
	tenDemList := []string{"Văn", "Thị", "Minh", "Ngọc", "Quốc", "Thuỳ", "Khánh", "Hải", "Hồng", "Thái"}
	tenList := []string{"Anh", "Bảo", "Chi", "Dũng", "Dung", "Giang", "Hà", "Hùng", "Lan", "Linh", "Phong", "Quân", "Thảo", "Trang", "Tuấn", "Vy"}

	rand.Seed(time.Now().UnixNano())
	users := make([]struct {
		Username  string
		FullName  string
		Email     string
		Photo     string
		CreatedAt time.Time
	}, count)

	for i := 1; i <= count; i++ {
		ho := hoList[rand.Intn(len(hoList))]
		tenDem := tenDemList[rand.Intn(len(tenDemList))]
		ten := tenList[rand.Intn(len(tenList))]
		fullName := fmt.Sprintf("%s %s %s", ho, tenDem, ten)

		emailPrefix := removeDiacritics(ho + ten)
		email := fmt.Sprintf("%s%d@gmail.com", emailPrefix, i)
		createdAt := time.Now().Add(-time.Duration(rand.Intn(180)) * 24 * time.Hour)

		// Định dạng photo URL
		photoBase := "https://xsgames.co/randomusers/assets/avatars"
		photoType := "male"
		if tenDem == "Thị" || ten == "Lan" || ten == "Vy" || ten == "Chi" {
			photoType = "female"
		}
		photo := fmt.Sprintf("%s/%s/%d.jpg", photoBase, photoType, i)

		// Generate username
		username := removeDiacritics(strings.ToLower(strings.ReplaceAll(fullName, " ", "")))

		users[i-1] = struct {
			Username  string
			FullName  string
			Email     string
			Photo     string
			CreatedAt time.Time
		}{
			Username:  username,
			FullName:  fullName,
			Email:     email,
			Photo:     photo,
			CreatedAt: createdAt,
		}
	}

	return users
}

// generatePartners generates a list of partners based on the count.
func generatePartners(count int) []struct {
	Username  string
	FullName  string
	Email     string
	Photo     string
	CreatedAt time.Time
} {
	// Common name components for partners
	firstNames := []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}
	lastNames := []string{"Tech", "Biz", "Solutions", "Systems", "Corp"}

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Ensure we have enough combinations
	if count > len(firstNames)*len(lastNames) {
		panic("Not enough unique combinations of first and last names for the count provided!")
	}

	partners := make([]struct {
		Username  string
		FullName  string
		Email     string
		Photo     string
		CreatedAt time.Time
	}, count)

	// Generate partners
	id := 1
	now := time.Now()
	for i := 0; i < count; i++ {
		// Cycle through names to generate unique combinations
		first := firstNames[(id-1)%len(firstNames)]
		last := lastNames[(id-1)/len(firstNames)%len(lastNames)]

		fullName := fmt.Sprintf("%s %s", first, last)
		username := strings.ToLower(strings.ReplaceAll(fmt.Sprintf("%s_%s_%d", first, last, id), " ", ""))
		email := fmt.Sprintf("%s@gmail.com", username)
		photo := fmt.Sprintf("https://xsgames.co/randomusers/assets/avatars/male/%d.jpg", id) // Male avatar

		// Random created_at within the last 4 months
		randomDays := rand.Intn(4 * 30) // Random up to ~120 days
		createdAt := now.AddDate(0, 0, -randomDays)

		partners[i] = struct {
			Username  string
			FullName  string
			Email     string
			Photo     string
			CreatedAt time.Time
		}{
			Username:  username,
			FullName:  fullName,
			Email:     email,
			Photo:     photo,
			CreatedAt: createdAt,
		}
		id++
	}

	return partners
}

// removeDiacritics removes diacritics (dấu) from Vietnamese strings.
func removeDiacritics(s string) string {
	var result strings.Builder
	result.Grow(len(s))
	for _, r := range s {
		switch r {
		case 'á', 'à', 'ả', 'ã', 'ạ', 'ă', 'ắ', 'ằ', 'ẳ', 'ẵ', 'ặ', 'â', 'ấ', 'ầ', 'ẩ', 'ẫ', 'ậ':
			r = 'a'
		case 'Á', 'À', 'Ả', 'Ã', 'Ạ', 'Ă', 'Ắ', 'Ằ', 'Ẳ', 'Ẵ', 'Ặ', 'Â', 'Ấ', 'Ầ', 'Ẩ', 'Ẫ', 'Ậ':
			r = 'A'
		case 'é', 'è', 'ẻ', 'ẽ', 'ẹ', 'ê', 'ế', 'ề', 'ể', 'ễ', 'ệ':
			r = 'e'
		case 'É', 'È', 'Ẻ', 'Ẽ', 'Ẹ', 'Ê', 'Ế', 'Ề', 'Ể', 'Ễ', 'Ệ':
			r = 'E'
		case 'í', 'ì', 'ỉ', 'ĩ', 'ị':
			r = 'i'
		case 'Í', 'Ì', 'Ỉ', 'Ĩ', 'Ị':
			r = 'I'
		case 'ó', 'ò', 'ỏ', 'õ', 'ọ', 'ô', 'ố', 'ồ', 'ổ', 'ỗ', 'ộ', 'ơ', 'ớ', 'ờ', 'ở', 'ỡ', 'ợ':
			r = 'o'
		case 'Ó', 'Ò', 'Ỏ', 'Õ', 'Ọ', 'Ô', 'Ố', 'Ồ', 'Ổ', 'Ỗ', 'Ộ', 'Ơ', 'Ớ', 'Ờ', 'Ở', 'Ỡ', 'Ợ':
			r = 'O'
		case 'ú', 'ù', 'ủ', 'ũ', 'ụ', 'ư', 'ứ', 'ừ', 'ử', 'ữ', 'ự':
			r = 'u'
		case 'Ú', 'Ù', 'Ủ', 'Ũ', 'Ụ', 'Ư', 'Ứ', 'Ừ', 'Ử', 'Ữ', 'Ự':
			r = 'U'
		case 'ý', 'ỳ', 'ỷ', 'ỹ', 'ỵ':
			r = 'y'
		case 'Ý', 'Ỳ', 'Ỷ', 'Ỹ', 'Ỵ':
			r = 'Y'
		case 'đ':
			r = 'd'
		case 'Đ':
			r = 'D'
		}
		result.WriteRune(r)
	}
	return result.String()
}

type FakeStore struct {
	Name         string
	BusinessType string
}

func fakeStores(count int) []FakeStore {
	storeData := []FakeStore{
		{"Apple", "Technology"},
		{"Nike", "Sportswear"},
		{"Adidas", "Sportswear"},
		{"Starbucks", "Coffee Shop"},
		{"McDonald's", "Fast Food"},
		{"Coca-Cola", "Beverages"},
		{"Samsung", "Technology"},
		{"Amazon", "E-commerce"},
		{"Google", "Technology"},
		{"Microsoft", "Technology"},
		{"IKEA", "Furniture"},
		{"Louis Vuitton", "Luxury Goods"},
		{"Gucci", "Luxury Goods"},
		{"Prada", "Luxury Goods"},
		{"Zara", "Clothing"},
		{"H&M", "Clothing"},
		{"Uniqlo", "Clothing"},
		{"Huawei", "Technology"},
		{"Sony", "Electronics"},
		{"BMW", "Automotive"},
		{"Mercedes-Benz", "Automotive"},
		{"Tesla", "Automotive"},
		{"Toyota", "Automotive"},
		{"Ford", "Automotive"},
		{"Pepsi", "Beverages"},
		{"KFC", "Fast Food"},
		{"Burger King", "Fast Food"},
		{"Dior", "Luxury Goods"},
		{"Chanel", "Luxury Goods"},
		{"Rolex", "Luxury Goods"},
		{"Levi's", "Clothing"},
		{"Puma", "Sportswear"},
		{"Under Armour", "Sportswear"},
		{"Netflix", "Entertainment"},
		{"Disney", "Entertainment"},
		{"L'Oréal", "Cosmetics"},
		{"Sephora", "Cosmetics"},
		{"Domino's Pizza", "Fast Food"},
		{"Pizza Hut", "Fast Food"},
		{"Patagonia", "Outdoor Gear"},
		{"The North Face", "Outdoor Gear"},
		{"Balenciaga", "Luxury Goods"},
		{"Versace", "Luxury Goods"},
		{"New Balance", "Sportswear"},
		{"Fossil", "Watches"},
		{"Ray-Ban", "Eyewear"},
		{"Lacoste", "Clothing"},
		{"Tiffany & Co.", "Jewelry"},
		{"Pandora", "Jewelry"},
		{"Hermès", "Luxury Goods"},
		{"Cartier", "Luxury Goods"},
		{"Vans", "Footwear"},
		{"Converse", "Footwear"},
		{"Columbia Sportswear", "Outdoor Gear"},
		{"Calvin Klein", "Clothing"},
		{"Tommy Hilfiger", "Clothing"},
		{"Hollister", "Clothing"},
		{"Abercrombie & Fitch", "Clothing"},
		{"Gap", "Clothing"},
		{"Banana Republic", "Clothing"},
		{"Reebok", "Sportswear"},
		{"ASICS", "Sportswear"},
		{"Oakley", "Eyewear"},
		{"Fila", "Sportswear"},
		{"Subway", "Fast Food"},
		{"Dunkin' Donuts", "Beverages"},
		{"Taco Bell", "Fast Food"},
		{"Chick-fil-A", "Fast Food"},
		{"Wendy's", "Fast Food"},
		{"Papa John's", "Fast Food"},
		{"Arby's", "Fast Food"},
		{"Panera Bread", "Fast Food"},
		{"Chipotle", "Fast Food"},
		{"Five Guys", "Fast Food"},
		{"Baskin-Robbins", "Ice Cream"},
		{"Krispy Kreme", "Beverages"},
		{"In-N-Out Burger", "Fast Food"},
		{"Jollibee", "Fast Food"},
		{"Shake Shack", "Fast Food"},
		{"Wingstop", "Fast Food"},
		{"Pret A Manger", "Fast Food"},
		{"Ben & Jerry's", "Ice Cream"},
		{"Tim Hortons", "Coffee Shop"},
		{"Costa Coffee", "Coffee Shop"},
		{"Cinnabon", "Bakery"},
		{"PizzaExpress", "Fast Food"},
		{"Nando's", "Fast Food"},
		{"Cheesecake Factory", "Bakery"},
		{"Caribou Coffee", "Coffee Shop"},
		{"Auntie Anne's", "Bakery"},
		{"Peet's Coffee", "Coffee Shop"},
		{"Café Amazon", "Coffee Shop"},
	}

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Shuffle storeData to randomize the selection
	rand.Shuffle(len(storeData), func(i, j int) {
		storeData[i], storeData[j] = storeData[j], storeData[i]
	})

	// Limit the count if it exceeds the available stores
	if count > len(storeData) {
		count = len(storeData)
	}

	// Return a slice of the requested number of stores
	return storeData[:count]
}
