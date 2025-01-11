-- Drop foreign key constraints
ALTER TABLE IF EXISTS "verify_emails" DROP CONSTRAINT IF EXISTS "verify_emails_username_fkey";
ALTER TABLE IF EXISTS "sessions" DROP CONSTRAINT IF EXISTS "sessions_username_fkey";
ALTER TABLE IF EXISTS "transfer_voucher" DROP CONSTRAINT IF EXISTS "transfer_voucher_from_fkey";
ALTER TABLE IF EXISTS "transfer_voucher" DROP CONSTRAINT IF EXISTS "transfer_voucher_to_fkey";
ALTER TABLE IF EXISTS "transfer_voucher" DROP CONSTRAINT IF EXISTS "transfer_voucher_voucher_id_fkey";
ALTER TABLE IF EXISTS "voucher_owner" DROP CONSTRAINT IF EXISTS "voucher_owner_username_fkey";
ALTER TABLE IF EXISTS "voucher_owner" DROP CONSTRAINT IF EXISTS "voucher_owner_voucher_id_fkey";
ALTER TABLE IF EXISTS "stores" DROP CONSTRAINT IF EXISTS "stores_owner_fkey";
ALTER TABLE IF EXISTS "branchs" DROP CONSTRAINT IF EXISTS "branchs_store_id_fkey";
ALTER TABLE IF EXISTS "events" DROP CONSTRAINT IF EXISTS "events_owner_fkey";
ALTER TABLE IF EXISTS "events" DROP CONSTRAINT IF EXISTS "events_game_id_fkey";
ALTER TABLE IF EXISTS "events" DROP CONSTRAINT IF EXISTS "events_store_id_fkey";
ALTER TABLE IF EXISTS "quizzes" DROP CONSTRAINT IF EXISTS "quizzes_event_id_fkey";
ALTER TABLE IF EXISTS "vouchers" DROP CONSTRAINT IF EXISTS "vouchers_event_id_fkey";

-- Drop check constraints
ALTER TABLE IF EXISTS "vouchers" DROP CONSTRAINT IF EXISTS "status_check";
ALTER TABLE IF EXISTS "games" DROP CONSTRAINT IF EXISTS "game_type";

-- Drop indexes
DROP INDEX IF EXISTS "transfer_voucher_from_idx";
DROP INDEX IF EXISTS "transfer_voucher_to_idx";
DROP INDEX IF EXISTS "transfer_voucher_from_to_idx";
DROP INDEX IF EXISTS "voucher_owner_username_idx";
DROP INDEX IF EXISTS "stores_owner_idx";
DROP INDEX IF EXISTS "branchs_store_id_name_idx";
DROP INDEX IF EXISTS "events_game_id_idx";
DROP INDEX IF EXISTS "events_store_id_idx";

-- Drop tables
DROP TABLE IF EXISTS "vouchers";
DROP TABLE IF EXISTS "quizzes";
DROP TABLE IF EXISTS "events";
DROP TABLE IF EXISTS "games";
DROP TABLE IF EXISTS "branchs";
DROP TABLE IF EXISTS "stores";
DROP TABLE IF EXISTS "voucher_owner";
DROP TABLE IF EXISTS "transfer_voucher";
DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "verify_emails";
DROP TABLE IF EXISTS "users";

-- Drop enum type
DROP TYPE IF EXISTS "events_status";