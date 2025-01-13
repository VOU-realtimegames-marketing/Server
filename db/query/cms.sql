-- name: GetStoresByOwner :many
SELECT id, name, business_type
FROM stores
WHERE owner = $1;

-- name: GetBranchesByOwner :many
SELECT b.id, b.name, b.position, b.city_name, b.country, b.address, b.emoji
FROM branchs b
JOIN stores s ON b.store_id = s.id
WHERE s.owner = $1;

-- name: GetEventsByOwner :many
SELECT e.id, e.name, e.photo, e.voucher_quantity, e.status, e.start_time, e.end_time, g.type AS game_type
FROM events e
JOIN games g ON e.game_id = g.id
WHERE e.owner = $1;

-- name: GetVouchersByEvent :many
SELECT v.id, v.qr_code, v.type, v.status, v.expires_at
FROM vouchers v
WHERE v.event_id = $1;


-- name: GetVoucherOwnersByEvent :many
SELECT vo.username, u.full_name, u.email, v.id AS voucher_id, v.type AS voucher_type, v.status AS voucher_status
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN users u ON vo.username = u.username
WHERE v.event_id = $1;

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
WHERE s.owner = $1 -- Chỉ lấy các store của owner
  AND vo.created_at >= NOW() - INTERVAL '1 MONTH'
GROUP BY vo.username, u.full_name, u.email, u.photo
ORDER BY MAX(vo.created_at) DESC
LIMIT 5;

-- name: GetVoucherStatsByMonth :many
SELECT 
    TO_CHAR(DATE_TRUNC('month', vo.created_at), 'YYYY-MM') AS month, -- Trả về chuỗi định dạng YYYY-MM
    g.type AS game_type,
    COUNT(*) AS total_vouchers
FROM voucher_owner vo
JOIN vouchers v ON vo.voucher_id = v.id
JOIN events e ON v.event_id = e.id
JOIN games g ON e.game_id = g.id
WHERE e.owner = $1 -- Chỉ lấy các sự kiện của owner
  AND vo.created_at >= NOW() - INTERVAL '6 MONTHS' -- Chỉ lấy dữ liệu trong 6 tháng gần đây
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

