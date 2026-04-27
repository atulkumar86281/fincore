-- Drop foreign key constraints first
ALTER TABLE "transfer" DROP CONSTRAINT IF EXISTS transfer_from_account_id_fkey;
ALTER TABLE "transfer" DROP CONSTRAINT IF EXISTS transfer_to_account_id_fkey;
ALTER TABLE "entry" DROP CONSTRAINT IF EXISTS entry_account_id_fkey;

-- Drop indexes
DROP INDEX IF EXISTS transfer_from_account_id_to_account_id_idx;
DROP INDEX IF EXISTS transfer_to_account_id_idx;
DROP INDEX IF EXISTS transfer_from_account_id_idx;
DROP INDEX IF EXISTS entry_account_id_idx;
DROP INDEX IF EXISTS account_owner_idx;

-- Drop tables (order matters due to dependencies)
DROP TABLE IF EXISTS "transfer";
DROP TABLE IF EXISTS "entry";
DROP TABLE IF EXISTS "account";