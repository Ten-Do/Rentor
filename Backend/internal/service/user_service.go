package service

import (
	"errors"
	"rentor/internal/logger"
	"rentor/internal/models"
	"rentor/internal/repository"
)

// userService implements UserService
type userService struct {
	repo        repository.UserRepository
	profileRepo repository.UserProfileRepository
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository, profileRepo repository.UserProfileRepository) UserService {
	return &userService{
		repo:        repo,
		profileRepo: profileRepo,
	}
}

// RegisterUser creates a new user (with validation)
func (s *userService) RegisterUser(input *models.CreateUserInput) (int, error) {
	if input.Email == "" && input.Phone == "" {
		return 0, errors.New("email or phone required")
	}

	// Check if user already exists
	if input.Email != "" {
		existing, _ := s.repo.GetUserByEmail(input.Email)
		if existing != nil {
			return 0, errors.New("user with this email already exists")
		}
	}

	if input.Phone != "" {
		existing, _ := s.repo.GetUserByPhone(input.Phone)
		if existing != nil {
			return 0, errors.New("user with this phone already exists")
		}
	}

	userID, err := s.repo.CreateUser(input.Phone, input.Email)
	if err != nil {
		return 0, err
	}

	// Create default profile
	profile := &models.UserProfile{
		UserID:     userID,
		FirstName:  "",
		Surname:    "",
		Patronymic: "",
	}
	_, err = s.profileRepo.CreateUserProfile(profile)
	if err != nil {
		// Log error but don't fail (user is already created)
		logger.Error("failed to create user profile", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		return userID, nil
	}

	return userID, nil
}

// GetUser retrieves a user by ID
func (s *userService) GetUser(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

// GetUserByEmail retrieves a user by email
func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}

// GetUserByPhone retrieves a user by phone
func (s *userService) GetUserByPhone(phone string) (*models.User, error) {
	return s.repo.GetUserByPhone(phone)
}

// FindOrCreateUserByEmail finds existing user or creates new one
func (s *userService) FindOrCreateUserByEmail(email string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// User exists
	if user != nil {
		return user, nil
	}

	// Create new user
	userID, err := s.repo.CreateUser("", email)
	if err != nil {
		return nil, err
	}

	// Create default profile
	profile := &models.UserProfile{
		UserID:     userID,
		FirstName:  "",
		Surname:    "",
		Patronymic: "",
	}
	_, _ = s.profileRepo.CreateUserProfile(profile)

	// Return the newly created user
	return s.repo.GetUserByID(userID)
}
