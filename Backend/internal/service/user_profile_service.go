package service

import (
	"errors"
	"rentor/internal/models"
	"rentor/internal/repository"
)

// userProfileService implements UserProfileService
type userProfileService struct {
	userProfileRepo repository.UserProfileRepository
	userRepo        repository.UserRepository
}

// NewUserProfileService creates a new user profile service
func NewUserProfileService(userRepo repository.UserRepository, userProfileRepo repository.UserProfileRepository) UserProfileService {
	return &userProfileService{
		userProfileRepo: userProfileRepo,
		userRepo:        userRepo,
	}
}

// GetUserProfile retrieves user profile
func (s *userProfileService) GetUserProfile(userID int) (*models.UserProfile, error) {
	profile, err := s.userProfileRepo.GetUserProfileByUserID(userID)
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
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	profile, err := s.userProfileRepo.GetUserProfileByUserID(userID)
	if err != nil {
		return err
	}
	if profile == nil {
		return errors.New("user profile not found")
	}

	user.Phone = input.Phone
	profile.FirstName = input.FirstName
	profile.Surname = input.Surname
	profile.Patronymic = input.Patronymic

	err = s.userRepo.UpdateUser(user.UserID, user)
	if err != nil {
		return err
	}
	err = s.userProfileRepo.UpdateUserProfile(profile.ID, profile)
	if err != nil {
		return err
	}

	return nil
}

// CreateDefaultUserProfile creates a default user profile
func (s *userProfileService) CreateDefaultUserProfile(userID int) error {
	profile := &models.UserProfile{
		UserID:     userID,
		FirstName:  "",
		Surname:    "",
		Patronymic: "",
	}
	_, err := s.userProfileRepo.CreateUserProfile(profile)
	return err
}
