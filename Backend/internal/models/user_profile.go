package models

import "time"

// UserProfile represents detailed profile information for a user
type UserProfile struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	FirstName  string    `json:"first_name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateUserProfileInput input data for creating a user profile
type CreateUserProfileInput struct {
	UserID int `json:"user_id"`
}

// UpdateUserProfileInput input data for updating a user profile
type UpdateUserProfileInput struct {
	FirstName  string `json:"first_name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Phone      string `json:"phone_number"`
}

type GetUserProfileOutput struct {
	UserID     int       `json:"user_id"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone_number"`
	FirstName  string    `json:"first_name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	CreatedAt  time.Time `json:"created_at"`
}
