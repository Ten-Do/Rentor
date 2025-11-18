-- +goose Up

-- Note: SQLite doesn't support CREATE TYPE/ENUM.
-- We'll store enum-like fields as TEXT (optionally guarded by application-level checks).

-- we store advertisement linked to user profiles
CREATE TABLE IF NOT EXISTS advertisement (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL, -- Foreign key to user table
    title TEXT NOT NULL, -- Title of the advertisement
    description TEXT, -- Detailed description of the advertisement
    price NUMERIC NOT NULL, -- Price of the placement
    type TEXT NOT NULL, -- Type of placement (apartment|house|room)
    rooms TEXT NOT NULL, -- Number of rooms (studio|1|2|3|4|5|6+)
    city TEXT NOT NULL, -- City where the placement is located
    address TEXT NOT NULL, -- Full address of the placement
    latitude REAL, -- Latitude for geolocation
    longitude REAL, -- Longitude for geolocation
    square REAL NOT NULL, -- Square footage of the placement
    status TEXT NOT NULL, -- Status of the advertisement (active|paused)
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- Timestamp of advertisement creation
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP, -- Timestamp of last advertisement update
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

-- table to store photos related to advertisements (images stored as URLs that point to external storage)
CREATE TABLE IF NOT EXISTS advertisement_photos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    advertisement_id INTEGER NOT NULL, -- Foreign key to advertisement table
    photo_url TEXT NOT NULL, -- URL of the photo
    FOREIGN KEY (advertisement_id) REFERENCES advertisement(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE IF EXISTS advertisement_photos;
DROP TABLE IF EXISTS advertisements;