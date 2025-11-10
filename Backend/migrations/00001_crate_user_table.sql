
-- + goose up

-- we not store passwords, only email and phone for OTP auth
-- other user details are in user_profile table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, -- Unique identifier for each user
    email VARCHAR(100) UNIQUE, -- User's email address
    phone_number VARCHAR(15) UNIQUE, -- E.164 format
);

-- - goose down

DROP TABLE IF EXISTS users;