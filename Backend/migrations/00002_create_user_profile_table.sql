
-- +goose up

-- we store additional user details in user_profile table
-- linked to users table via user_id foreign key

CREATE TABLE IF NOT EXISTS user_profile (
    id SERIAL PRIMARY KEY, -- Unique identifier for each profile
    user_id INTEGER NOT NULL UNIQUE, -- Foreign key to user table
    first_name VARCHAR(100), -- User's first name
    surname VARCHAR(100), -- User's surname
    patronymic VARCHAR(100), -- User's patronymic
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of profile creation
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of last profile update
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE -- Foreign key constraint linking to users table
);

-- +goose down

DROP TABLE IF EXISTS user_profile;