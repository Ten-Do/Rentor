package models

import "time"

// Advertisement represents an advertisement in the system
type Advertisement struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Price       float64   `json:"price"`
	Type        string    `json:"type"`
	Rooms       string    `json:"rooms"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	Latitude    *float64  `json:"latitude"`
	Longitude   *float64  `json:"longitude"`
	Square      float64   `json:"square"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateAdvertisementInput input data for creating an advertisement
type CreateAdvertisementInput struct {
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	Price       float64  `json:"price"`
	Type        string   `json:"type"`
	Rooms       string   `json:"rooms"`
	City        string   `json:"city"`
	Address     string   `json:"address"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	Square      float64  `json:"square"`
}

// UpdateAdvertisementInput input data for updating an advertisement
type UpdateAdvertisementInput struct {
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	Price       float64  `json:"price"`
	Type        string   `json:"type"`
	Rooms       string   `json:"rooms"`
	City        string   `json:"city"`
	Address     string   `json:"address"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	Square      float64  `json:"square"`
	Status      string   `json:"status"`
}

type GetAd struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	Price       float64  `json:"price"`
	Type        string   `json:"type"`
	Rooms       string   `json:"rooms"`
	City        string   `json:"city"`
	Address     string   `json:"address"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
	Square      float64  `json:"square"`
	Status      string   `json:"status"`

	LandlordName  *string `json:"landlordName"`
	LandlordEmail string  `json:"landlordEmail"`
	LandlordPhone *string `json:"landlordPhone"`

	ImageUrls []string `json:"imageUrls"`
}

type AdPreview struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	City     string  `json:"city"`
	Price    float64 `json:"price"`
	Type     string  `json:"type"`
	Rooms    string  `json:"rooms"`
	Square   float64 `json:"square"`
	ImageUrl *string `json:"imageUrl"` // первое фото
}

type GetAdPreviewsList struct {
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Items []AdPreview `json:"items"`
}

type AdFilters struct {
	Page     int      `json:"page"`
	Limit    int      `json:"limit"`
	MinPrice *float64 `json:"minPrice,omitempty"`
	MaxPrice *float64 `json:"maxPrice,omitempty"`
	Type     *string  `json:"type,omitempty"`
	Rooms    *string  `json:"rooms,omitempty"`
	City     *string  `json:"city,omitempty"`
	Keywords *string  `json:"keywords,omitempty"`
	UserID   *int     `json:"userId,omitempty"` // нужно для /advertisements/my
}

type ImagesUploadResponse struct {
	Uploaded []string `json:"uploaded"`
	Count    int      `json:"count"`
}
