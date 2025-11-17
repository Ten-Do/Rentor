package models

import "time"

// OTPCode represents an OTP record in the system
type OTPCode struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Email       string    `json:"email"`
	CodeHash    string    `json:"-"` // never expose hash
	Attempts    int       `json:"attempts"`
	MaxAttempts int       `json:"max_attempts"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// OTPRequest for sending OTP
type OTPRequest struct {
	Email string `json:"email"`
}

// OTPVerifyRequest for verifying OTP
type OTPVerifyRequest struct {
	Email   string `json:"email"`
	OtpCode string `json:"otp_code"`
}
