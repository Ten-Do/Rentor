package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"rentor/internal/logger"
	"rentor/internal/models"
	"rentor/internal/service"
)

type AuthHandler struct {
	userService    service.UserService
	otpService     service.OTPService
	jwtService     service.JWTService
	otpLength      int
	otpExpMin      int
	otpMaxAttempts int
}

func NewAuthHandler(userSvc service.UserService, otpSvc service.OTPService, jwtSvc service.JWTService, otpLen int, otpExpMin int, otpMaxAttempts int) *AuthHandler {
	return &AuthHandler{
		userService:    userSvc,
		otpService:     otpSvc,
		jwtService:     jwtSvc,
		otpLength:      otpLen,
		otpExpMin:      otpExpMin,
		otpMaxAttempts: otpMaxAttempts,
	}
}

// SendOTP handles POST /auth/send-otp
func (h *AuthHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	var req models.OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	// Validate email format
	req.Email = json.Number(req.Email).String() // no-op, just ensures it's a string
	if req.Email == "" {
		writeError(w, http.StatusBadRequest, "email is required")
		return
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		writeError(w, http.StatusBadRequest, "invalid email format")
		return
	}

	logger.Info("SendOTP called", logger.Field("email", req.Email))

	// Find or create user
	user, err := h.userService.FindOrCreateUserByEmail(req.Email)
	if err != nil {
		logger.Error("failed to find/create user", logger.Field("error", err.Error()))
		writeError(w, http.StatusInternalServerError, "failed to process request")
		return
	}

	// Generate and store OTP
	err = h.otpService.GenerateAndStoreOTP(user.ID, req.Email, h.otpLength, h.otpExpMin, h.otpMaxAttempts)
	if err != nil {
		logger.Error("failed to generate OTP", logger.Field("error", err.Error()))
		writeError(w, http.StatusInternalServerError, "failed to send OTP")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "OTP sent to email",
	})
}

// VerifyOTP handles POST /auth/verify-otp
func (h *AuthHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req models.OTPVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if req.Email == "" || req.OtpCode == "" {
		writeError(w, http.StatusBadRequest, "email and otp_code are required")
		return
	}

	logger.Info("VerifyOTP called", logger.Field("email", req.Email))

	// Verify OTP
	userID, err := h.otpService.VerifyOTP(req.Email, req.OtpCode, h.otpMaxAttempts)
	if err != nil {
		logger.Warn("OTP verification failed", logger.Field("error", err.Error()), logger.Field("email", req.Email))
		writeError(w, http.StatusUnauthorized, fmt.Sprintf("invalid OTP: %s", err.Error()))
		return
	}

	// Get user
	user, err := h.userService.GetUser(userID)
	if err != nil {
		logger.Error("failed to get user", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		writeError(w, http.StatusInternalServerError, "failed to authenticate")
		return
	}

	// Generate JWT token
	accessToken, err := h.jwtService.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		logger.Error("failed to generate JWT", logger.Field("error", err.Error()))
		writeError(w, http.StatusInternalServerError, "failed to authenticate")
		return
	}

	// Set JWT in httpOnly secure cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,                            // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,             // Use Strict in production
		Expires:  time.Now().Add(15 * time.Minute), // Match JWT TTL
	})

	logger.Info("user authenticated", logger.Field("user_id", user.ID), logger.Field("email", user.Email))

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"user": map[string]interface{}{
			"id":         user.ID,
			"email":      user.Email,
			"phone":      user.Phone,
			"created_at": user.CreatedAt,
		},
	})
}

// Logout handles POST /auth/logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Clear JWT cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
	})

	logger.Info("user logged out")
	writeJSON(w, http.StatusOK, map[string]string{"message": "logged out"})
}

// common helpers
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]interface{}{"error": message})
}
