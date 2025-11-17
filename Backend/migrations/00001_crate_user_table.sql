-- +goose Up

-- we do not store passwords here, only email and phone for OTP auth
-- other user details are in user_profile table
CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- Unique identifier for each user
    email TEXT UNIQUE, -- User's email address
    phone_number TEXT UNIQUE, -- E.164 format
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- Timestamp of user creation
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP -- Timestamp of last user update
);

-- +goose Down

DROP TABLE IF EXISTS user;