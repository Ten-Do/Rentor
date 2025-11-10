-- +goose Up

-- we store additional user details in user_profile table
-- linked to users table via user_id foreign key

CREATE TABLE IF NOT EXISTS user_profile (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- Unique identifier for each profile
    user_id INTEGER NOT NULL UNIQUE, -- Foreign key to users table
    first_name TEXT, -- User's first name
    surname TEXT, -- User's surname
    patronymic TEXT, -- User's patronymic
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- Timestamp of profile creation
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- Timestamp of last profile update
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE IF EXISTS user_profile;