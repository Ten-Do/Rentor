
-- +goose up

-- type_of_placement_enum type to represent different types of placements
CREATE TYPE type_of_placement_enum IF NOT EXISTS AS ENUM ('apartment', 'house', 'room');

-- room_enum type to represent number of rooms in an advertisement placement
 CREATE TYPE room_enum IF NOT EXISTS AS ENUM ('studio', '1', '2', '3', '4', '5', '6+');

CREATE TYPE status_enum IF NOT EXISTS AS ENUM ('active', 'paused');

-- we store advertisements linked to user profiles
CREATE TABLE IF NOT EXISTS advertisements (
    id SERIAL PRIMARY KEY,
    user_profile_id INTEGER NOT NULL, -- Foreign key to user_profile table (not users because ads are linked to profiles, its just for simplicity)
    title VARCHAR(255), -- Title of the advertisement
    description TEXT, -- Detailed description of the advertisement
    price DECIMAL(20, 2), -- Price of the placement (rent price, for example 50000.50 Tenge)
    type_of_placement type_of_placement_enum, -- Type of placement using enum type
    rooms room_enum, -- Number of rooms using room_enum type
    city VARCHAR(255), -- City where the placement is located
    address VARCHAR(500), -- Full address of the placement
    latitude FLOAT, -- Latitude for geolocation
    longitude FLOAT, -- Longitude for geolocation
    square FLOAT, -- Square footage of the placement
    status status_enum, -- Status of the advertisement using status_enum type
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of advertisement creation
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of last advertisement update
    FOREIGN KEY (user_profile_id) REFERENCES user_profile(id) ON DELETE CASCADE -- Foreign key constraint linking to user_profile table
);

-- table to store photos related to advertisements (images stored as URLs that point to external storage)
CREATE TABLE IF NOT EXISTS advertisement_photos (
    id SERIAL PRIMARY KEY,
    advertisement_id INTEGER NOT NULL, -- Foreign key to advertisements table
    photo_url VARCHAR(500) NOT NULL, -- URL of the photo
    FOREIGN KEY (advertisement_id) REFERENCES advertisements(id) ON DELETE CASCADE -- Foreign key constraint linking to advertisements table
);

-- +goose down

DROP TABLE IF EXISTS advertisements;
DROP TABLE IF EXISTS advertisement_photos;
DROP TYPE IF EXISTS type_of_placement_enum;
DROP TYPE IF EXISTS room_enum;
DROP TYPE IF EXISTS status_enum;