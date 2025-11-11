package store

import (
	"database/sql"

	"rentor/internal/repository"
	"rentor/internal/service"
)

// Store includes all layers of the application.
// This is the central place for initializing all layers of the application
type Store struct {
	// Repositories (working with DB)
	User          repository.UserRepository          // User repository
	UserProfile   repository.UserProfileRepository   // User profile repository
	Advertisement repository.AdvertisementRepository // Advertisement repository

	// Services (business logic)
	UserService          service.UserService          // User service
	UserProfileService   service.UserProfileService   // User profile service
	AdvertisementService service.AdvertisementService // Advertisement service
}

// NewStore creates a new store with initialized layers
func NewStore(db *sql.DB) *Store {
	// Create repositories
	userRepo := repository.NewUserRepository(db)
	userProfileRepo := repository.NewUserProfileRepository(db)
	advertisementRepo := repository.NewAdvertisementRepository(db)

	// Create services, passing repositories to them
	userService := service.NewUserService(userRepo)
	userProfileService := service.NewUserProfileService(userProfileRepo)
	advertisementService := service.NewAdvertisementService(advertisementRepo)

	return &Store{
		User:                 userRepo,
		UserProfile:          userProfileRepo,
		Advertisement:        advertisementRepo,
		UserService:          userService,
		UserProfileService:   userProfileService,
		AdvertisementService: advertisementService,
	}
}
