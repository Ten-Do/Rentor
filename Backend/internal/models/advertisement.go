package models

import "time"

// Advertisement represents an advertisement in the system
type Advertisement struct {
	ID            int       `json:"id"`
	UserProfileID int       `json:"user_profile_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	Type          string    `json:"type"`
	Rooms         string    `json:"rooms"`
	City          string    `json:"city"`
	Address       string    `json:"address"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	Square        float64   `json:"square"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CreateAdvertisementInput input data for creating an advertisement
type CreateAdvertisementInput struct {
	UserProfileID int     `json:"user_profile_id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Type          string  `json:"type"`
	Rooms         string  `json:"rooms"`
	City          string  `json:"city"`
	Address       string  `json:"address"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Square        float64 `json:"square"`
}

// UpdateAdvertisementInput input data for updating an advertisement
type UpdateAdvertisementInput struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Type        string  `json:"type"`
	Rooms       string  `json:"rooms"`
	City        string  `json:"city"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Square      float64 `json:"square"`
	Status      string  `json:"status"`
}
