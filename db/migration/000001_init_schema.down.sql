-- Drop foreign key constraints
ALTER TABLE "verify_emails" DROP CONSTRAINT "verify_emails_username_fkey";
ALTER TABLE "sessions" DROP CONSTRAINT "sessions_username_fkey";
ALTER TABLE "transfer_voucher" DROP CONSTRAINT "transfer_voucher_from_fkey";
ALTER TABLE "transfer_voucher" DROP CONSTRAINT "transfer_voucher_to_fkey";
ALTER TABLE "transfer_voucher" DROP CONSTRAINT "transfer_voucher_voucher_id_fkey";
ALTER TABLE "voucher_owner" DROP CONSTRAINT "voucher_owner_username_fkey";
ALTER TABLE "voucher_owner" DROP CONSTRAINT "voucher_owner_voucher_id_fkey";
ALTER TABLE "stores" DROP CONSTRAINT "stores_owner_fkey";
ALTER TABLE "branchs" DROP CONSTRAINT "branchs_store_id_fkey";
ALTER TABLE "events" DROP CONSTRAINT "events_game_id_fkey";
ALTER TABLE "vouchers" DROP CONSTRAINT "vouchers_event_id_fkey";

-- Drop check constraints
ALTER TABLE "vouchers" DROP CONSTRAINT "status_check";
ALTER TABLE "games" DROP CONSTRAINT "game_type";

-- Drop indexes
DROP INDEX IF EXISTS "transfer_voucher_from_idx";
DROP INDEX IF EXISTS "transfer_voucher_to_idx";
DROP INDEX IF EXISTS "transfer_voucher_from_to_idx";
DROP INDEX IF EXISTS "voucher_owner_username_idx";
DROP INDEX IF EXISTS "stores_owner_idx";

-- Drop tables
DROP TABLE IF EXISTS "vouchers";
DROP TABLE IF EXISTS "events";
DROP TABLE IF EXISTS "games";
DROP TABLE IF EXISTS "branchs";
DROP TABLE IF EXISTS "stores";
DROP TABLE IF EXISTS "voucher_owner";
DROP TABLE IF EXISTS "transfer_voucher";
DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "verify_emails";
DROP TABLE IF EXISTS "users";
