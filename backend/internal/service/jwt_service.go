package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTClaims custom claims structure
type JWTClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Type   string `json:"type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

// jwtService implementation
type jwtService struct {
	secret          string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
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
func (s *jwtService) GenerateRefreshToken(userID int, email string) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
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

// ValidateAccessToken validates and returns claims from an access token
func (s *jwtService) ValidateAccessToken(tokenString string) (*JWTClaims, error) {
	return s.validateToken(tokenString, "access")
}

func (s *jwtService) ValidateRefreshToken(tokenString string) (*JWTClaims, error) {
	return s.validateToken(tokenString, "refresh")
}

func (s *jwtService) validateToken(tokenString, expectedType string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	if claims.Type != expectedType {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}

func (s *jwtService) GetRefreshTokenTTL() time.Duration {
	return s.refreshTokenTTL
}

func (s *jwtService) GetAccessTokenTTL() time.Duration {
	return s.accessTokenTTL
}

// RefreshAccessToken generates a new access token from a refresh token
func (s *jwtService) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := s.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	return s.GenerateAccessToken(claims.UserID, claims.Email)
}
