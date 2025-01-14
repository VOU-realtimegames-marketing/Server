-- name: CreateFakeUser :one
INSERT INTO users (username, hashed_password, full_name, email, role, created_at)
VALUES ($1, $2, $3, $4, $5, NOW())
ON CONFLICT (username) DO UPDATE SET 
    hashed_password = EXCLUDED.hashed_password,
    full_name = EXCLUDED.full_name,
    email = EXCLUDED.email,
    role = EXCLUDED.role,
    created_at = EXCLUDED.created_at
RETURNING username;

-- name: CheckStoreExists :one
SELECT COALESCE((SELECT id 
                 FROM stores 
                 WHERE name = $1 AND owner = $2 
                 LIMIT 1), 0) AS id;

-- name: CreateFakeStore :one
INSERT INTO stores (name, owner, business_type)
VALUES ($1, $2, $3)
RETURNING id;

-- name: CreateFakeBranch :one
INSERT INTO branchs (store_id, name, position, city_name, country, address, emoji)
VALUES ($1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (store_id, name) DO NOTHING
RETURNING id;

-- name: CheckGameExists :one
SELECT id 
FROM games
WHERE type = $1;

-- name: CreateGame :one
INSERT INTO games (name, photo, type, play_guide, gift_allowed)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: CheckEventExists :one
SELECT id 
FROM events
WHERE owner = $1 
  AND game_id = $2 
  AND store_id = $3 
  AND name = $4;

-- name: CreateFakeEvent :one
INSERT INTO events (game_id, store_id, owner, name, photo, voucher_quantity, status, start_time, end_time)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;

-- name: CheckVoucherExists :one
SELECT id 
FROM vouchers
WHERE event_id = $1 AND qr_code = $2 AND type = $3;

-- name: CreateVoucher :one
INSERT INTO vouchers (event_id, qr_code, type, status, expires_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: CheckVoucherOwnerExists :one
SELECT id 
FROM voucher_owner
WHERE username = $1 AND voucher_id = $2;

-- name: CreateVoucherOwner :one
INSERT INTO voucher_owner (username, voucher_id, created_at)
VALUES ($1, $2, $3)
RETURNING id;

-- name: CountVoucherOwners :one
SELECT COUNT(*) AS total
FROM voucher_owner;

-- name: GetUserByUsername :one
SELECT username, hashed_password, full_name, email, role, active
FROM users
WHERE username = $1;

-- name: GetStoresByOwner :many
SELECT id, name, business_type
FROM stores
WHERE owner = $1;

-- name: GetBranchesByStore :many
SELECT b.id, b.name, b.position, b.city_name, b.country, b.address, b.emoji
FROM branchs b
WHERE b.store_id = $1;

-- name: GetEventsByStore :many
SELECT e.id, e.name, e.photo, e.voucher_quantity, e.status, e.start_time, e.end_time, g.type AS game_type
FROM events e
JOIN games g ON e.game_id = g.id
WHERE e.store_id = $1;

-- name: GetVouchersByEvent :many
SELECT v.id, v.qr_code, v.type, v.status, v.expires_at
FROM vouchers v
WHERE v.event_id = $1;

-- name: GetVoucherOwnersByVoucher :many
SELECT vo.username, u.full_name, u.email, vo.created_at
FROM voucher_owner vo
JOIN users u ON vo.username = u.username
WHERE vo.voucher_id = $1;

-- name: GetCmsOverview :one
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
       AND e.start_time < DATE_TRUNC('month', NOW())) AS last_month_total_user_play;

-- name: GetUserPlayByDate :many
SELECT 
    DATE(vo.created_at) AS play_date,
    g.type AS game_type,
    COUNT(DISTINCT vo.username) AS total_users
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN events e ON v.event_id = e.id
JOIN games g ON e.game_id = g.id
WHERE e.owner = $1
  AND vo.created_at >= NOW() - INTERVAL '2 MONTHS'
GROUP BY play_date, game_type
ORDER BY play_date, game_type;

-- name: GetRecentVoucherOwners :many
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
LIMIT 5;

-- name: GetVoucherStatsByMonth :many
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
ORDER BY month;

-- name: GetUserStatsByStore :many
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
ORDER BY total_users DESC;

-- name: GetEventCreatedStats :many
SELECT 
    DATE(start_time) AS date,
    COUNT(CASE WHEN g.type = 'quiz' THEN 1 END) AS quiz_game,
    COUNT(CASE WHEN g.type = 'phone-shake' THEN 1 END) AS shake_game
FROM events e
JOIN games g ON e.game_id = g.id
WHERE start_time >= NOW() - INTERVAL '2 MONTHS'
GROUP BY DATE(start_time)
ORDER BY DATE(start_time);

-- name: GetUserStoreStats :many
SELECT 
    s.owner AS username,
    COUNT(DISTINCT vo.username) AS total_user_play
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN events e ON v.event_id = e.id
JOIN stores s ON e.store_id = s.id
GROUP BY s.owner;

-- name: GetUserPlayStats :many
SELECT 
    TO_CHAR(DATE_TRUNC('month', vo.created_at), 'Month') AS month,
    COUNT(CASE WHEN g.type = 'quiz' THEN 1 END) AS quiz_game,
    COUNT(CASE WHEN g.type = 'phone-shake' THEN 1 END) AS shake_game
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN events e ON v.event_id = e.id
JOIN games g ON e.game_id = g.id
GROUP BY DATE_TRUNC('month', vo.created_at)
ORDER BY DATE_TRUNC('month', vo.created_at);

-- name: GetRecentUsers :many
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
LIMIT 5;

-- name: GetAdminStats :one
SELECT 
    (SELECT COUNT(*) FROM users WHERE role = 'partner') AS total_partner,
    (SELECT COUNT(*) FROM users WHERE role = 'partner' AND created_at >= NOW() - INTERVAL '1 MONTH') AS total_partner_last_month,
    (SELECT COUNT(*) FROM users WHERE role = 'user') AS total_user,
    (SELECT COUNT(*) FROM users WHERE role = 'user' AND created_at >= NOW() - INTERVAL '1 MONTH') AS total_user_last_month,
    (SELECT COUNT(*) FROM branchs) AS total_branch,
    (SELECT COUNT(*) FROM branchs WHERE created_at >= NOW() - INTERVAL '1 MONTH') AS total_branch_last_month,
    (SELECT SUM(v.voucher_quantity * 10) FROM events e JOIN vouchers v ON e.id = v.event_id) AS total_earning,
    (SELECT SUM(v.voucher_quantity * 10) FROM events e JOIN vouchers v ON e.id = v.event_id WHERE e.start_time >= NOW() - INTERVAL '1 MONTH') AS total_earning_last_month;


