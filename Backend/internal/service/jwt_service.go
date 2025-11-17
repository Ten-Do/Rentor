package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTService handles JWT token operations
type JWTService interface {
	GenerateAccessToken(userID int, email string) (string, error)
	GenerateRefreshToken(userID int) (string, error)
	ValidateToken(tokenString string) (int, error) // returns userID
	RefreshAccessToken(refreshToken string) (string, error)
}

type jwtService struct {
	secret          string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

// JWTClaims custom claims structure
type JWTClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Type   string `json:"type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

// NewJWTService creates a new JWT service
func NewJWTService(secret string, accessTTL, refreshTTL time.Duration) JWTService {
	return &jwtService{
		secret:          secret,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}

// GenerateAccessToken creates a new access JWT token
func (s *jwtService) GenerateAccessToken(userID int, email string) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "rentor",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

// GenerateRefreshToken creates a new refresh JWT token
func (s *jwtService) GenerateRefreshToken(userID int) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID: userID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "rentor",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

// ValidateToken validates and extracts claims from token
func (s *jwtService) ValidateToken(tokenString string) (int, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return 0, errors.New("token expired")
	}

	return claims.UserID, nil
}

// RefreshAccessToken generates new access token from refresh token
func (s *jwtService) RefreshAccessToken(refreshToken string) (string, error) {
	userID, err := s.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.secret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	if claims.Type != "refresh" {
		return "", errors.New("token is not a refresh token")
	}

	return s.GenerateAccessToken(userID, claims.Email)
}
