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
	router.Post("/auth/refresh", authHandler.RefreshToken)

	// Protected routes (require JWT)
	authMiddleware := middleware.AuthMiddlewareWithRefresh(dataStore.JWTService, "access_token", "refresh_token")

	// logout
	router.With(authMiddleware).Post("/auth/logout", authHandler.Logout)

	// User profile
	userProfileHandler := handlers.NewUserProfileHandler(dataStore.UserService, dataStore.UserProfileService)
	router.With(authMiddleware).Get("/user/profile", userProfileHandler.GetUserProfile)
	router.With(authMiddleware).Put("/user/profile", userProfileHandler.UpdateUserProfile)

	// Advertisements
	adsHandler := handlers.NewAdvertisementHandlers(dataStore.AdService, dataStore.ImageService)
	router.Get("/advertisements", adsHandler.ListAdvertisements)
	router.With(authMiddleware).Post("/advertisements", adsHandler.CreateAdvertisement)
	router.Get("/advertisements/{id}", adsHandler.GetAdvertisement)
	router.With(authMiddleware).Put("/advertisements/{id}", adsHandler.UpdateAdvertisement)
	router.With(authMiddleware).Delete("/advertisements/{id}", adsHandler.DeleteAdvertisement)
	router.With(authMiddleware).Post("/advertisements/{id}/images", adsHandler.AddAdImages)
	router.With(authMiddleware).Delete("/advertisements/{ad_id}/images/{image_id}", adsHandler.DeleteAdImage)
	router.With(authMiddleware).Get("/advertisements/my", adsHandler.GetMyAdvertisements)
}
