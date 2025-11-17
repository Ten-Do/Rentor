package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"rentor/internal/logger"
	"rentor/internal/service"
)

type ContextKey string

const UserIDKey ContextKey = "user_id"

func AuthMiddleware(jwtService service.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Get JWT from cookie
			cookie, err := r.Cookie("session")
			if err != nil {
				logger.Warn("missing session cookie", logger.Field("path", r.URL.Path))
				http.Error(w, `{"error": "unauthorized"}`, http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimSpace(cookie.Value)
			if tokenString == "" {
				logger.Warn("empty session cookie", logger.Field("path", r.URL.Path))
				http.Error(w, `{"error": "unauthorized"}`, http.StatusUnauthorized)
				return
			}

			// Validate JWT
			userID, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				logger.Warn("invalid JWT token",
					logger.Field("error", err.Error()),
					logger.Field("path", r.URL.Path),
				)
				http.Error(w, `{"error": "unauthorized"}`, http.StatusUnauthorized)
				return
			}

			// Add user ID into context
			ctx := context.WithValue(r.Context(), UserIDKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIDFromContext(r *http.Request) (int, error) {
	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		return 0, errors.New("user id not found in context")
	}
	return userID, nil
}
