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
	"time"

	"github.com/jackc/pgx/v5/pgtype"
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

func (server *Server) FakeCmsOverview(ctx context.Context, req *gen.FakeCmsOverviewRequest) (*gen.FakeCmsOverviewResponse, error) {
	log.Println("Starting FakeCmsOverview")

	// Step 1: Create a partner user
	hashedPassword, err := utils.HashPassword("12341234")
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, status.Errorf(codes.Internal, "error hashing password")
	}
	partnerUsername := "partner_user"
	_, err = server.store.CreateFakeUser(ctx, db.CreateFakeUserParams{
		Username:       partnerUsername,
		HashedPassword: hashedPassword,
		FullName:       "Partner User",
		Email:          "partner_user@example.com",
		Role:           "partner",
	})
	if err != nil {
		log.Printf("Failed to create partner user: %v", err)
		return nil, status.Errorf(codes.Internal, "error creating partner user")
	}
	log.Println("Step 1: Partner user created")

	log.Println("======== Start step 2 ========")
	storeNames := []string{"Highland Coffee", "Starbucks"}
	storeIDs := make([]int64, 0)
	for _, name := range storeNames {
		// Check if the store already exists
		existingStoreId, err := server.store.CheckStoreExists(ctx, db.CheckStoreExistsParams{
			Name:  name,
			Owner: partnerUsername,
		})

		log.Println("existingStoreId: ", existingStoreId)
		log.Println("err: ", err)

		if existingStoreId != 0 {
			log.Printf("Store %s for owner %s already exists, skipping.", name, partnerUsername)
			storeIDs = append(storeIDs, existingStoreId.(int64))
			continue
		}

		// Create the store
		storeID, err := server.store.CreateFakeStore(ctx, db.CreateFakeStoreParams{
			Name:         name,
			Owner:        partnerUsername,
			BusinessType: "Coffee Shop",
		})
		if err != nil {
			log.Printf("Failed to create store %s: %v", name, err)
			continue
		}
		storeIDs = append(storeIDs, storeID)
		log.Printf("Store created: %s, ID: %d", name, storeID)
	}

	// Step 3: Create branches for each store
	log.Println("======== Start step 3 ========")
	branchNames := map[string][]string{
		"Highland Coffee": {"Highland Vincom", "Highland Landmark 81"},
		"Starbucks":       {"Starbucks District 1", "Starbucks District 7"},
	}
	branchIDs := make(map[string][]int64)

	for storeIndex, storeID := range storeIDs {
		storeName := storeNames[storeIndex]
		branches := branchNames[storeName]
		for _, branchName := range branches {
			// Create branch
			branchID, err := server.store.CreateFakeBranch(ctx, db.CreateFakeBranchParams{
				StoreID:  storeID,
				Name:     branchName,
				Position: fmt.Sprintf("Position of %s", branchName),
				CityName: "Ho Chi Minh City",
				Country:  "Vietnam",
				Address:  fmt.Sprintf("%s address", branchName),
				Emoji:    "🏢",
			})
			if err != nil {
				log.Printf("Failed to create branch %s for store %s: %v", branchName, storeName, err)
				continue
			}
			branchIDs[storeName] = append(branchIDs[storeName], branchID)
			log.Printf("Branch created: %s, ID: %d", branchName, branchID)
		}
	}

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
	voucherTypes := []string{"Discount 50%", "Discount 20%", "Free Coffee", "Buy 1 Get 1 Free"}
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
			newVoucherID, err := server.store.CreateVoucher(ctx, db.CreateVoucherParams{
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
	fakeUsers := []struct {
		Username string
		FullName string
		Email    string
		Password string
	}{
		{"fakeUserA", "Fake User A", "fakeA@gmail.com", "12341234"},
		{"fakeUserB", "Fake User B", "fakeB@gmail.com", "12341234"},
		{"fakeUserC", "Fake User C", "fakeC@gmail.com", "12341234"},
		{"nguyenanhA", "Nguyễn Anh A", "nguyenanhA@gmail.com", "12341234"},
		{"lethithaoB", "Lê Thị Thảo B", "lethithaoB@gmail.com", "12341234"},
		{"tranvanC", "Trần Văn C", "tranvanC@gmail.com", "12341234"},
		{"phamminhD", "Phạm Minh D", "phamminhD@gmail.com", "12341234"},
		{"dangthuyE", "Đặng Thúy E", "dangthuyE@gmail.com", "12341234"},
		{"huynhngocF", "Huỳnh Ngọc F", "huynhngocF@gmail.com", "12341234"},
		{"vohuyG", "Võ Huy G", "vohuyG@gmail.com", "12341234"},
		{"doquyenH", "Đỗ Quyên H", "doquyenH@gmail.com", "12341234"},
		{"buiducI", "Bùi Đức I", "buiducI@gmail.com", "12341234"},
		{"hoangmaiJ", "Hoàng Mai J", "hoangmaiJ@gmail.com", "12341234"},
	}

	usernames := []string{}

	for _, user := range fakeUsers {
		// Hash password
		hashedPassword, err := utils.HashPassword(user.Password)
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
			voucherOwnerID, err := server.store.CreateVoucherOwner(ctx, db.CreateVoucherOwnerParams{
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

	log.Println("FakeCmsOverview completed successfully")
	return &gen.FakeCmsOverviewResponse{}, nil
}

// func (server *Server) MockCmsData(ctx context.Context) error {
// 	// 1. Tạo user partner "partner_user"
// 	passwordHash, _ := utils.HashPassword("12341234")
// 	_, err := server.store.CreateUserIfNotExists(ctx, db.CreateUserIfNotExistsParams{
// 		Username:       "partner_user",
// 		HashedPassword: passwordHash,
// 		FullName:       "Partner User",
// 		Email:          "partner_user@example.com",
// 		Role:           "partner",
// 	})
// 	if err != nil {
// 		log.Printf("Failed to create user: %v", err)
// 		return err
// 	}
// 	log.Println("User `partner_user` created or already exists.")

// 	// 2. Tạo stores
// 	stores := []string{"Highland Coffee", "Starbucks"}
// 	storeIDs := []int64{}
// 	for _, storeName := range stores {
// 		store, err := server.store.CreateStoreIfNotExists(ctx, db.CreateStoreIfNotExistsParams{
// 			Name:         storeName,
// 			Owner:        "partner_user",
// 			BusinessType: "cafe",
// 		})
// 		if err != nil {
// 			log.Printf("Failed to create store: %v", err)
// 			return err
// 		}
// 		storeIDs = append(storeIDs, store.ID)
// 		log.Printf("Store `%s` created or already exists.", storeName)
// 	}

// 	// 3. Tạo branches cho mỗi store
// 	branches := map[string][]string{
// 		"Highland Coffee": {"Highland Coffee Quận 1", "Highland Coffee Quận 3"},
// 		"Starbucks":       {"Starbucks Hàn Thuyên", "Starbucks VivoCity"},
// 	}
// 	for i, storeID := range storeIDs {
// 		storeName := stores[i]
// 		for _, branchName := range branches[storeName] {
// 			_, err := server.store.CreateBranchIfNotExists(ctx, db.CreateBranchIfNotExistsParams{
// 				StoreID:  storeID,
// 				Name:     branchName,
// 				Position: "",
// 				CityName: "Hồ Chí Minh",
// 				Country:  "Việt Nam",
// 				Address:  branchName,
// 				Emoji:    "☕",
// 			})
// 			if err != nil {
// 				log.Printf("Failed to create branch: %v", err)
// 				return err
// 			}
// 			log.Printf("Branch `%s` created or already exists.", branchName)
// 		}
// 	}

// 	// 4. Tạo games
// 	games := []struct {
// 		Name string
// 		Type string
// 	}{
// 		{"phone-shake", "phone-shake"},
// 		{"quiz", "quiz"},
// 	}
// 	gameIDs := []int64{}
// 	for _, game := range games {
// 		g, err := server.store.CreateGameIfNotExists(ctx, db.CreateGameIfNotExistsParams{
// 			Name:        game.Name,
// 			Type:        game.Type,
// 			Photo:       "default-game.jpg",
// 			PlayGuide:   "Guide here",
// 			GiftAllowed: true,
// 		})
// 		if err != nil {
// 			log.Printf("Failed to create game: %v", err)
// 			return err
// 		}
// 		gameIDs = append(gameIDs, g.ID)
// 		log.Printf("Game `%s` created or already exists.", game.Name)
// 	}

// 	// 5. Tạo events
// 	events := []string{"Trung thu 2024", "Black Friday 2024", "Giáng sinh 2024", "Tết dương lịch 2025", "Tết âm lịch 2025"}
// 	eventIDs := []int64{}
// 	for _, storeID := range storeIDs {
// 		for _, eventName := range events {
// 			for _, gameID := range gameIDs {
// 				event, err := server.store.CreateEvent(ctx, db.CreateEventParams{
// 					Name:      eventName,
// 					Photo:     "default-event.jpg",
// 					GameID:    gameID,
// 					StoreID:   storeID,
// 					Owner:     "partner_user",
// 					StartTime: time.Now().Add(24 * time.Hour),
// 					EndTime:   time.Now().Add(48 * time.Hour),
// 					Status:    "ready",
// 				})
// 				if err != nil {
// 					log.Printf("Failed to create event: %v", err)
// 					return err
// 				}
// 				eventIDs = append(eventIDs, event.ID)
// 				log.Printf("Event `%s` created.", eventName)
// 			}
// 		}
// 	}

// 	// 6. Tạo vouchers cho mỗi event
// 	vouchers := []string{"Discount 50%", "Discount 20%", "Free drink", "Buy 1 Get 1 Free"}
// 	for _, eventID := range eventIDs {
// 		for _, voucherName := range vouchers {
// 			_, err := server.store.CreateVoucher(ctx, db.CreateVoucherParams{
// 				EventID:   eventID,
// 				Type:      "discount",
// 				Status:    "valid",
// 				ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
// 			})
// 			if err != nil {
// 				log.Printf("Failed to create voucher: %v", err)
// 				return err
// 			}
// 			log.Printf("Voucher `%s` created for event ID `%d`.", voucherName, eventID)
// 		}
// 	}

// 	// 7. Tạo users thường
// 	users := []string{"user1", "user2", "user3", "user4", "user5"}
// 	for _, username := range users {
// 		_, err := server.store.CreateUserIfNotExists(ctx, db.CreateUserIfNotExistsParams{
// 			Username: username,
// 			FullName: "User " + username,
// 			Email:    username + "@example.com",
// 			Role:     "user",
// 			Password: "password123",
// 		})
// 		if err != nil {
// 			log.Printf("Failed to create user: %v", err)
// 			return err
// 		}
// 		log.Printf("User `%s` created or already exists.", username)
// 	}

// 	// 8. Tạo voucher_owner
// 	for _, eventID := range eventIDs {
// 		vouchers, err := server.store.GetVouchersByEvent(ctx, eventID)
// 		if err != nil {
// 			log.Printf("Failed to get vouchers for event ID `%d`: %v", eventID, err)
// 			return err
// 		}
// 		for _, voucher := range vouchers {
// 			for _, username := range users {
// 				_, err := server.store.CreateVoucherOwner(ctx, db.CreateVoucherOwnerParams{
// 					Username:  username,
// 					VoucherID: voucher.ID,
// 					CreatedAt: time.Now(),
// 				})
// 				if err != nil {
// 					log.Printf("Failed to create voucher owner: %v", err)
// 					return err
// 				}
// 				log.Printf("User `%s` received voucher ID `%d`.", username, voucher.ID)
// 			}
// 		}
// 	}

// 	return nil
// }
