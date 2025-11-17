package service

import (
	"crypto/rand"
	"math/big"
	"errors"
	"fmt"
	"time"

	"rentor/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// OTPService handles OTP generation, storage, and verification
type OTPService interface {
	GenerateAndStoreOTP(userID int, email string, otpLength int, expirationMinutes int, maxAttempts int) error
	VerifyOTP(email string, otpCode string, maxAttempts int) (int, error) // returns userID
	CleanupExpiredOTPs() error
}

type otpService struct {
	repo repository.OTPRepository
}

// NewOTPService creates a new OTP service
func NewOTPService(repo repository.OTPRepository) OTPService {
	return &otpService{
		repo: repo,
	}
}

// GenerateOTP generates a random OTP string
func (s *otpService) generateOTP(length int) (string, error) {
	const digits = "0123456789"

	otp := make([]byte, length)
    for i := 0; i < length; i++ {
        n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
        if err != nil {
            return "", err
        }
        otp[i] = digits[n.Int64()]
    }
    return string(otp), nil
}

// GenerateAndStoreOTP creates a new OTP, hashes it, and stores in DB
func (s *otpService) GenerateAndStoreOTP(userID int, email string, otpLength int, expirationMinutes int, maxAttempts int) error {
	// Generate OTP
	otpCode, err := s.generateOTP(otpLength)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	// Hash OTP using bcrypt
	hashedCode, err := bcrypt.GenerateFromPassword([]byte(otpCode), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash OTP: %w", err)
	}

	// Delete any existing OTP for this email
	_ = s.repo.DeleteOTPByEmail(email)

	// Store in database
	expiresAt := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)
	err = s.repo.CreateOTP(userID, email, string(hashedCode), maxAttempts, expiresAt)
	if err != nil {
		return fmt.Errorf("failed to store OTP: %w", err)
	}

	// TODO: In production, send OTP via email
	// For now, just log it (NEVER do this in production)
	fmt.Printf("[DEV] OTP for %s: %s (expires at %v)\n", email, otpCode, expiresAt)

	return nil
}

// VerifyOTP verifies the provided OTP code
func (s *otpService) VerifyOTP(email string, otpCode string, maxAttempts int) (int, error) {
	// Get OTP record from DB
	otpRecord, err := s.repo.GetOTPByEmail(email)
	if err != nil {
		return 0, fmt.Errorf("OTP not found: %w", err)
	}

	// Check if OTP is expired
	if time.Now().After(otpRecord.ExpiresAt) {
		_ = s.repo.DeleteOTPByID(otpRecord.ID)
		return 0, errors.New("OTP expired")
	}

	// Check if max attempts exceeded
	if otpRecord.Attempts >= otpRecord.MaxAttempts {
		_ = s.repo.DeleteOTPByID(otpRecord.ID)
		return 0, errors.New("too many failed attempts, OTP locked")
	}

	// Compare OTP (bcrypt comparison)
	err = bcrypt.CompareHashAndPassword([]byte(otpRecord.CodeHash), []byte(otpCode))
	if err != nil {
		// Increment attempts
		otpRecord.Attempts++
		_ = s.repo.UpdateOTPAttempts(otpRecord.ID, otpRecord.Attempts)
		return 0, fmt.Errorf("invalid OTP (attempts: %d/%d)", otpRecord.Attempts, otpRecord.MaxAttempts)
	}

	// OTP is valid, delete it from DB
	_ = s.repo.DeleteOTPByID(otpRecord.ID)

	return otpRecord.UserID, nil
}

// CleanupExpiredOTPs removes expired OTP records
func (s *otpService) CleanupExpiredOTPs() error {
	return s.repo.DeleteExpiredOTPs(time.Now())
}
