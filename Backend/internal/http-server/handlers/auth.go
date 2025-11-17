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

// AuthHandler handles authentication-related endpoints
type AuthHandler struct {
	userService    service.UserService
	otpService     service.OTPService
	jwtService     service.JWTService
	otpLength      int
	otpExpMin      int
	otpMaxAttempts int
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(userSvc service.UserService, otpSvc service.OTPService, jwtSvc service.JWTService,
	otpLen int, otpExpMin int, otpMaxAttempts int) *AuthHandler {
	return &AuthHandler{
		userService:    userSvc,
		otpService:     otpSvc,
		jwtService:     jwtSvc,
		otpLength:      otpLen,
		otpExpMin:      otpExpMin,
		otpMaxAttempts: otpMaxAttempts,
	}
}

// SendOTP sends an OTP to the provided email address
func (h *AuthHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	var req models.OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if req.Email == "" {
		writeError(w, http.StatusBadRequest, "email is required")
		return
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		writeError(w, http.StatusBadRequest, "invalid email format")
		return
	}

	logger.Info("SendOTP called", logger.Field("email", req.Email))

	user, err := h.userService.FindOrCreateUserByEmail(req.Email)
	if err != nil {
		logger.Error("failed to find/create user", logger.Field("error", err.Error()))
		writeError(w, http.StatusInternalServerError, "failed to process request")
		return
	}

	err = h.otpService.GenerateAndStoreOTP(user.UserID, req.Email, h.otpLength, h.otpExpMin, h.otpMaxAttempts)
	if err != nil {
		logger.Error("failed to generate OTP", logger.Field("error", err.Error()))
		writeError(w, http.StatusInternalServerError, "failed to send OTP")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "OTP sent to email"})
}

// VerifyOTP verifies the provided OTP and generates JWT tokens
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

	userID, err := h.otpService.VerifyOTP(req.Email, req.OtpCode, h.otpMaxAttempts)
	if err != nil {
		logger.Warn("OTP verification failed", logger.Field("error", err.Error()), logger.Field("email", req.Email))
		writeError(w, http.StatusUnauthorized, fmt.Sprintf("invalid OTP: %s", err.Error()))
		return
	}

	user, err := h.userService.GetUser(userID)
	if err != nil {
		logger.Error("failed to get user", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		writeError(w, http.StatusInternalServerError, "failed to authenticate")
		return
	}

	// Generate JWT tokens
	accessToken, err := h.jwtService.GenerateAccessToken(user.UserID, user.Email)
	if err != nil {
		logger.Error("failed to generate access token", logger.Field("error", err.Error()))
		writeError(w, http.StatusInternalServerError, "failed to authenticate")
		return
	}

	refreshToken, err := h.jwtService.GenerateRefreshToken(user.UserID, user.Email)
	if err != nil {
		logger.Error("failed to generate refresh token", logger.Field("error", err.Error()))
		writeError(w, http.StatusInternalServerError, "failed to authenticate")
		return
	}

	// Set refresh token in httpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(h.jwtService.GetRefreshTokenTTL()),
	})

	logger.Info("user authenticated", logger.Field("user_id", user.UserID), logger.Field("email", user.Email))

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"access_token": accessToken,
		"user": map[string]interface{}{
			"id":         user.UserID,
			"email":      user.Email,
			"phone":      user.Phone,
			"created_at": user.CreatedAt,
		},
	})
}

// RefreshToken refresh access token
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil || cookie.Value == "" {
		logger.Warn("refresh token missing")
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "refresh token required"})
		return
	}

	claims, err := h.jwtService.ValidateRefreshToken(cookie.Value)
	if err != nil {
		logger.Warn("invalid refresh token", logger.Field("error", err.Error()))
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		return
	}

	userID := claims.UserID
	userEmail := claims.Email

	accessToken, err := h.jwtService.GenerateAccessToken(userID, userEmail)
	if err != nil {
		logger.Error("failed to generate access token", logger.Field("error", err.Error()))
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to generate access token"})
		return
	}

	// Optionally refresh refresh token
	newRefreshToken, err := h.jwtService.GenerateRefreshToken(userID, userEmail)
	if err == nil {
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    newRefreshToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
			Expires:  time.Now().Add(h.jwtService.GetRefreshTokenTTL()),
		})
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"access_token": accessToken,
	})
}

// Logout clears refresh token
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
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

// --- helpers ---
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]interface{}{"error": message})
}
