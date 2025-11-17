package service

import (
	"errors"
	"rentor/internal/models"
	"rentor/internal/repository"
)

// userProfileService implements UserProfileService
type userProfileService struct {
	repo repository.UserProfileRepository
}

// NewUserProfileService creates a new user profile service
func NewUserProfileService(repo repository.UserProfileRepository) UserProfileService {
	return &userProfileService{
		repo: repo,
	}
}

// GetUserProfile retrieves user profile
func (s *userProfileService) GetUserProfile(userID int) (*models.UserProfile, error) {
	profile, err := s.repo.GetUserProfileByUserID(userID)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, errors.New("user profile not found")
	}
	return profile, nil
}

// UpdateUserProfile updates user profile
func (s *userProfileService) UpdateUserProfile(userID int, input *models.UpdateUserProfileInput) error {
	profile, err := s.repo.GetUserProfileByUserID(userID)
	if err != nil {
		return err
	}
	if profile == nil {
		return errors.New("user profile not found")
	}

	profile.FirstName = input.FirstName
	profile.Surname = input.Surname
	profile.Patronymic = input.Patronymic

	return s.repo.UpdateUserProfile(profile.ID, profile)
}

// CreateDefaultUserProfile creates a default user profile
func (s *userProfileService) CreateDefaultUserProfile(userID int) error {
	profile := &models.UserProfile{
		UserID:     userID,
		FirstName:  "",
		Surname:    "",
		Patronymic: "",
	}
	_, err := s.repo.CreateUserProfile(profile)
	return err
}
