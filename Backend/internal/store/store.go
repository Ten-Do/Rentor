package store

import (
	"database/sql"

	"rentor/internal/config"
	"rentor/internal/repository"
	"rentor/internal/service"
)

// Store includes all layers of the application.
// This is the central place for initializing all layers of the application
type Store struct {
	// Repositories (working with DB)
	User        repository.UserRepository
	UserProfile repository.UserProfileRepository
	OTP         repository.OTPRepository

	// Services (business logic)
	UserService        service.UserService
	UserProfileService service.UserProfileService
	OTPService         service.OTPService
	JWTService         service.JWTService
	EmailService       service.EmailService
}

// NewStore creates a new store with initialized layers
func NewStore(db *sql.DB, cfg *config.Config) *Store {
	// Create repositories
	userRepo := repository.NewUserRepository(db)
	userProfileRepo := repository.NewUserProfileRepository(db)
	otpRepo := repository.NewOTPRepository(db)

	// Create services, passing repositories to them
	userService := service.NewUserService(userRepo, userProfileRepo)
	userProfileService := service.NewUserProfileService(userRepo, userProfileRepo)
	jwtService := service.NewJWTService(
		cfg.Auth.JWTSecret,
		cfg.Auth.AccessTokenTTL,
		cfg.Auth.RefreshTokenTTL,
	)
	emailService := service.NewEmailService(cfg.SMTP.SMTPFrom, cfg.SMTP.SMTPPassWord, cfg.SMTP.SMTPHost, cfg.SMTP.SMTPPort)
	otpService := service.NewOTPService(otpRepo, emailService)

	return &Store{
		User:               userRepo,
		UserProfile:        userProfileRepo,
		OTP:                otpRepo,
		UserService:        userService,
		UserProfileService: userProfileService,
		OTPService:         otpService,
		JWTService:         jwtService,
		EmailService:       emailService,
	}
}
