package repository

import (
	"database/sql"
	"errors"
	"time"

	"rentor/internal/models"
)

// otpRepository implements OTPRepository
type otpRepository struct {
	db *sql.DB
}

// NewOTPRepository creates a new OTP repository
func NewOTPRepository(db *sql.DB) OTPRepository {
	return &otpRepository{db: db}
}

// CreateOTP creates a new OTP record
func (r *otpRepository) CreateOTP(userID int, email string, codeHash string, maxAttempts int, expiresAt time.Time) error {
	_, err := r.db.Exec(
		"INSERT INTO otp_codes (user_id, email, code_hash, max_attempts, expires_at) VALUES (?, ?, ?, ?, ?)",
		userID,
		email,
		codeHash,
		maxAttempts,
		expiresAt,
	)
	return err
}

// GetOTPByEmail retrieves OTP record by email
func (r *otpRepository) GetOTPByEmail(email string) (*models.OTPCode, error) {
	otp := &models.OTPCode{}
	err := r.db.QueryRow(
		"SELECT id, user_id, email, code_hash, attempts, max_attempts, expires_at, created_at FROM otp_codes WHERE email = ?",
		email,
	).Scan(&otp.ID, &otp.UserID, &otp.Email, &otp.CodeHash, &otp.Attempts, &otp.MaxAttempts, &otp.ExpiresAt, &otp.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("OTP not found")
		}
		return nil, err
	}

	return otp, nil
}

// GetOTPByID retrieves OTP record by ID
func (r *otpRepository) GetOTPByID(id int) (*models.OTPCode, error) {
	otp := &models.OTPCode{}
	err := r.db.QueryRow(
		"SELECT id, user_id, email, code_hash, attempts, max_attempts, expires_at, created_at FROM otp_codes WHERE id = ?",
		id,
	).Scan(&otp.ID, &otp.UserID, &otp.Email, &otp.CodeHash, &otp.Attempts, &otp.MaxAttempts, &otp.ExpiresAt, &otp.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("OTP not found")
		}
		return nil, err
	}

	return otp, nil
}

// UpdateOTPAttempts updates the attempts counter
func (r *otpRepository) UpdateOTPAttempts(id int, attempts int) error {
	_, err := r.db.Exec(
		"UPDATE otp_codes SET attempts = ? WHERE id = ?",
		attempts,
		id,
	)
	return err
}

// DeleteOTPByID deletes OTP record by ID
func (r *otpRepository) DeleteOTPByID(id int) error {
	_, err := r.db.Exec("DELETE FROM otp_codes WHERE id = ?", id)
	return err
}

// DeleteOTPByEmail deletes OTP record by email
func (r *otpRepository) DeleteOTPByEmail(email string) error {
	_, err := r.db.Exec("DELETE FROM otp_codes WHERE email = ?", email)
	return err
}

// DeleteExpiredOTPs deletes all expired OTP records
func (r *otpRepository) DeleteExpiredOTPs(now time.Time) error {
	_, err := r.db.Exec("DELETE FROM otp_codes WHERE expires_at < ?", now)
	return err
}
