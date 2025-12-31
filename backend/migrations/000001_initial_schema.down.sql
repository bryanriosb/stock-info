-- Rollback: 000001_initial_schema
-- Description: Drop all initial tables

DROP INDEX IF EXISTS idx_refresh_tokens_token;
DROP INDEX IF EXISTS idx_refresh_tokens_user_id;
DROP TABLE IF EXISTS refresh_tokens;

DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_username;
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_stocks_symbol;
DROP TABLE IF EXISTS stocks;
