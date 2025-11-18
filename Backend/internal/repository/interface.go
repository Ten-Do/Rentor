package repository

import (
	"rentor/internal/models"
	"time"
)

// UserRepository interface for working with users in the DB
type UserRepository interface {
	CreateUser(phone string, email string) (int, error)     // creates a new user with phone and email (one or both required)
	GetUserByID(id int) (*models.User, error)               // retrieves a user by their ID
	GetUserByEmail(email string) (*models.User, error)      // retrieves a user by their email
	GetUserByPhone(phone string) (*models.User, error)      // retrieves a user by their phone
	GetAllUsers() ([]*models.User, error)                   // retrieves all users
	GetPageUsers(offset, limit int) ([]*models.User, error) // retrieves users with pagination
	UpdateUser(id int, user *models.User) error             // updates user details
	DeleteUserByID(id int) error                            // deletes a user by their ID
	DeleteUserByPhone(phone string) error                   // deletes a user by their phone
	DeleteUserByEmail(email string) error                   // deletes a user by their email
}

// UserProfileRepository interface for working with user profiles in the DB
type UserProfileRepository interface {
	CreateUserProfile(profile *models.UserProfile) (int, error)           // creates a new user profile
	GetUserProfileByID(id int) (*models.UserProfile, error)               // retrieves a user profile by its ID
	GetUserProfileByUserID(userID int) (*models.UserProfile, error)       // retrieves a user profile by the associated user ID
	GetAllUserProfiles() ([]*models.UserProfile, error)                   // retrieves all user profiles
	GetPageUserProfiles(offset, limit int) ([]*models.UserProfile, error) // retrieves user profiles with pagination
	UpdateUserProfile(id int, profile *models.UserProfile) error          // updates user profile details
	DeleteUserProfileByID(id int) error                                   // deletes a user profile by its ID
}

// OTPRepository interface for working with OTP codes in the DB
type OTPRepository interface {
	CreateOTP(userID int, email string, codeHash string, maxAttempts int, expiresAt time.Time) error
	GetOTPByEmail(email string) (*models.OTPCode, error)
	GetOTPByID(id int) (*models.OTPCode, error)
	UpdateOTPAttempts(id int, attempts int) error
	DeleteOTPByID(id int) error
	DeleteOTPByEmail(email string) error
	DeleteExpiredOTPs(now time.Time) error
}

// AdvertisementRepository interface for working with advertisements in the DB
type AdvertisementRepository interface {
	CreateAdvertisement(ad *models.Advertisement) (int, error)                            // creates a new advertisement
	GetAdvertisementByID(id int) (*models.Advertisement, error)                           // retrieves an advertisement by its ID
	GetAllAdvertisements() ([]*models.Advertisement, error)                               // retrieves all advertisements
	GetPageAdvertisements(offset, limit int) ([]*models.Advertisement, error)             // retrieves advertisements with pagination
	GetAllUserAdvertisements(userID int) ([]*models.Advertisement, error)                 // retrieves all advertisements for a specific user
	GetPageUserAdvertisements(userID, offset, limit int) ([]*models.Advertisement, error) // retrieves advertisements for a specific user with pagination
	UpdateAdvertisement(id int, ad *models.Advertisement) error                           // updates advertisement details
	DeleteAdvertisementByID(id int) error                                                 // deletes an advertisement by its ID
	GetImagePath(adID, imageID int) (string, error)
}
