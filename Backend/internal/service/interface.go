package service

import "rentor/internal/models"

// UserService interface for user business logic
type UserService interface {
	RegisterUser(input *models.CreateUserInput) (int, error)
	GetUser(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByPhone(phone string) (*models.User, error)
	FindOrCreateUserByEmail(email string) (*models.User, error)
}

// UserProfileService interface for user profile business logic
type UserProfileService interface {
	GetUserProfile(userID int) (*models.UserProfile, error)
	UpdateUserProfile(userID int, input *models.UpdateUserProfileInput) error
	CreateDefaultUserProfile(userID int) error
}

type AdvertisementService interface {
}
