package service

import "rentor/internal/models"

// UserService interface for user business logic
type UserService interface {
	RegisterUser(input *models.CreateUserInput) (int, error)
	GetUser(id int) (*models.User, error)
}

// UserProfileService interface for user profile business logic
type UserProfileService interface {
	GetUserProfile(id int) (*models.UserProfile, error)
	UpdateUserProfile(id int, input *models.UpdateUserProfileInput) error
}

type AdvertisementService interface {
}
