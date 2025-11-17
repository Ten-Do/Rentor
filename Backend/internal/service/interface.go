package service

import (
	"mime/multipart"
	"rentor/internal/models"
	"time"
)

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

// JWTService handles JWT token operations
type JWTService interface {
	GenerateAccessToken(userID int, email string) (string, error)
	GenerateRefreshToken(userID int, email string) (string, error)
	ValidateAccessToken(tokenString string) (*JWTClaims, error)
	ValidateRefreshToken(tokenString string) (*JWTClaims, error)
	RefreshAccessToken(refreshToken string) (string, error)
	GetRefreshTokenTTL() time.Duration
	GetAccessTokenTTL() time.Duration
}

// OTPService handles OTP generation, storage, and verification
type OTPService interface {
	GenerateAndStoreOTP(userID int, email string, otpLength int, expirationMinutes int, maxAttempts int) error
	VerifyOTP(email string, otpCode string, maxAttempts int) (int, error) // returns userID
	CleanupExpiredOTPs() error
}

type EmailService interface {
	SendEmail(to, subject, body string) error
}

type AdvertisementService interface {
	CreateAdvertisement(userID int, input *models.CreateAdvertisementInput) (*models.GetAd, error)
	AddImages(adID int, urls []string, userID int) error
	GetAdvertisement(id int) (*models.GetAd, error)
	GetAdvertisements(filters *models.AdFilters) (*models.GetAdPreviewsList, error)
	GetMyAdvertisements(userID, page, limit int) (*models.GetAdPreviewsList, error)
	UpdateAdvertisement(id int, userID int, input *models.UpdateAdvertisementInput) error
	DeleteAdvertisement(id int, userID int) error
	DeleteImage(adID, imageID, userID int) error
}

// ImageService интерфейс для работы с изображениями
type ImageServiceInterface interface {
	SaveAdvertisementImages(adID int, files []*multipart.FileHeader) ([]string, error)
}
