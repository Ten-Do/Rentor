package middleware

import (
	"context"
	"errors"
	"net/http"

	"rentor/internal/logger"
	"rentor/internal/service"
)

type ContextKey string

const UserIDKey ContextKey = "user_id"

// AuthMiddlewareWithRefresh checks access token from cookie
// cookieAccessName — cookie with access token
// cookieRefreshName — cookie with refresh token
func AuthMiddlewareWithRefresh(jwtService service.JWTService, cookieAccessName, cookieRefreshName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Попытка получить access token из cookie
			accessCookie, err := r.Cookie(cookieAccessName)
			if err != nil || accessCookie.Value == "" {
				// Если нет access, пробуем refresh
				userID, newAccess, err := refreshAccessFromCookie(jwtService, r, cookieRefreshName)
				if err != nil {
					logger.Warn("no valid token", logger.Field("error", err.Error()))
					http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
					return
				}

				// Кладём новый access в cookie
				http.SetCookie(w, &http.Cookie{
					Name:     cookieAccessName,
					Value:    newAccess,
					Path:     "/",
					HttpOnly: true,
					Secure:   false, // true на продакшене
					SameSite: http.SameSiteLaxMode,
				})

				ctx := context.WithValue(r.Context(), UserIDKey, userID)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// Валидация access token
			claims, err := jwtService.ValidateAccessToken(accessCookie.Value)
			if err != nil {
				// Попытка обновить через refresh
				userID, newAccess, err := refreshAccessFromCookie(jwtService, r, cookieRefreshName)
				if err != nil {
					logger.Warn("invalid access token and refresh failed", logger.Field("error", err.Error()))
					http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
					return
				}

				// Кладём новый access в cookie
				http.SetCookie(w, &http.Cookie{
					Name:     cookieAccessName,
					Value:    newAccess,
					Path:     "/",
					HttpOnly: true,
					Secure:   false,
					SameSite: http.SameSiteLaxMode,
				})

				ctx := context.WithValue(r.Context(), UserIDKey, userID)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// Всё ок, access токен валиден
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// refreshAccessFromCookie проверяет refresh token и создаёт новый access token
func refreshAccessFromCookie(jwtService service.JWTService, r *http.Request, cookieRefreshName string) (userID int, newAccess string, err error) {
	refreshCookie, err := r.Cookie(cookieRefreshName)
	if err != nil || refreshCookie.Value == "" {
		return 0, "", errors.New("refresh token missing")
	}

	// Получаем userID из refresh token
	claims, err := jwtService.ValidateRefreshToken(refreshCookie.Value)
	if err != nil {
		return 0, "", err
	}

	newAccess, err = jwtService.GenerateAccessToken(claims.UserID, claims.Email)
	if err != nil {
		return 0, "", err
	}

	return claims.UserID, newAccess, nil

}

// GetUserIDFromContext достаёт userID из контекста
func GetUserIDFromContext(r *http.Request) (int, error) {
	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		return 0, errors.New("user id not found in context")
	}
	return userID, nil
}
