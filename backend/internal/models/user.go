package models

import "time"

// User represents a user in the system (without profile details)
type User struct {
	UserID    int       `json:"user_id"`
	Phone     *string   `json:"phone_number"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserInput input data for creating a user
type CreateUserInput struct {
	Phone *string `json:"phone_number"`
	Email string  `json:"email"`
}

// UpdateUserInput input data for updating a user
type UpdateUserInput struct {
	UserID int     `json:"user_id"`
	Phone  *string `json:"phone_number"`
	Email  string  `json:"email"`
}
