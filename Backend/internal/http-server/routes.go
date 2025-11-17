package httpserver

import (
	"rentor/internal/http-server/handlers"
	"rentor/internal/http-server/middleware"
	"rentor/internal/store"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers HTTP routes.
func RegisterRoutes(router chi.Router, dataStore *store.Store, otpLen int, otpExpMin int, otpMaxAttempts int) {
	// Authentication (no middleware required)
	authHandler := handlers.NewAuthHandler(dataStore.UserService, dataStore.OTPService, dataStore.JWTService, otpLen, otpExpMin, otpMaxAttempts)
	router.Post("/auth/send-otp", authHandler.SendOTP)
	router.Post("/auth/verify-otp", authHandler.VerifyOTP)

	// Protected routes (require JWT)
	authMiddleware := middleware.AuthMiddleware(dataStore.JWTService)

	// logout
	router.With(authMiddleware).Post("/auth/logout", authHandler.Logout)

	// User profile
	userProfileHandler := handlers.NewUserProfileHandler(dataStore.UserProfileService)
	router.With(authMiddleware).Get("/user/profile", userProfileHandler.GetUserProfile)
	router.With(authMiddleware).Put("/user/profile", userProfileHandler.UpdateUserProfile)

	// Advertisements (for now, not protected)
	router.Get("/advertisements", handlers.ListAdvertisements)
	router.Post("/advertisements", handlers.CreateAdvertisement)
	router.Get("/advertisements/{id}", handlers.GetAdvertisement)
	router.Put("/advertisements/{id}", handlers.UpdateAdvertisement)
	router.Delete("/advertisements/{id}", handlers.DeleteAdvertisement)

	router.Post("/advertisements/{id}/images", handlers.AddAdImages)
	router.Delete("/advertisements/{ad_id}/images/{image_id}", handlers.DeleteAdImage)
}
