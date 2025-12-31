-- Migration: 000001_initial_schema
-- Description: Create initial tables for stocks, users, and refresh_tokens

-- Stocks table
CREATE TABLE IF NOT EXISTS stocks (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    symbol VARCHAR(20) NOT NULL,
    name VARCHAR(255) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    exchange VARCHAR(100) NOT NULL,
    mic_code VARCHAR(20) NOT NULL,
    country VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_stocks_symbol ON stocks (symbol);

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_username ON users (username);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users (email);

-- Refresh tokens table
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    user_id INT8 NOT NULL,
    token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    revoked BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user_id ON refresh_tokens (user_id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_refresh_tokens_token ON refresh_tokens (token);
