package httpserver

import (
	"net/http"
	"rentor/internal/config"
	"rentor/internal/http-server/handlers"
	"rentor/internal/http-server/middleware"
	"rentor/internal/logger"
	"rentor/internal/store"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers HTTP routes.
func RegisterRoutes(router chi.Router, dataStore *store.Store, cfg *config.Config) {
	log := logger.With(logger.Field("component", "http-server"))

	// Authentication (no middleware required)
	authHandler := handlers.NewAuthHandler(dataStore.UserService, dataStore.OTPService, dataStore.JWTService, cfg.Auth.OTPLength, cfg.Auth.OTPExpirationMinutes, cfg.Auth.OTPMaxAttempts)
	router.Post("/auth/send-otp", authHandler.SendOTP)
	log.Info("registered route", logger.Field("path", "/auth/send-otp"), logger.Field("method", "POST"))
	router.Post("/auth/verify-otp", authHandler.VerifyOTP)
	log.Info("registered route", logger.Field("path", "/auth/verify-otp"), logger.Field("method", "POST"))
	router.Post("/auth/refresh", authHandler.RefreshToken)
	log.Info("registered route", logger.Field("path", "/auth/refresh"), logger.Field("method", "POST"))

	// Protected routes (require JWT)
	authMiddleware := middleware.AuthMiddlewareWithRefresh(dataStore.JWTService, "access_token", "refresh_token")

	// logout
	router.With(authMiddleware).Post("/auth/logout", authHandler.Logout)
	log.Info("registered route", logger.Field("path", "/auth/logout"), logger.Field("method", "POST"))

	// User profile
	userProfileHandler := handlers.NewUserProfileHandler(dataStore.UserService, dataStore.UserProfileService)
	router.With(authMiddleware).Get("/user/profile", userProfileHandler.GetUserProfile)
	log.Info("registered route", logger.Field("path", "/user/profile"), logger.Field("method", "GET"))
	router.With(authMiddleware).Put("/user/profile", userProfileHandler.UpdateUserProfile)
	log.Info("registered route", logger.Field("path", "/user/profile"), logger.Field("method", "PUT"))

	// Advertisements
	adsHandler := handlers.NewAdvertisementHandlers(dataStore.AdService, dataStore.ImageService)
	router.Get("/advertisements", adsHandler.ListAdvertisements)
	log.Info("registered route", logger.Field("path", "/advertisements"), logger.Field("method", "GET"))
	router.With(authMiddleware).Post("/advertisements", adsHandler.CreateAdvertisement)
	log.Info("registered route", logger.Field("path", "/advertisements"), logger.Field("method", "POST"))
	router.Get("/advertisements/{id}", adsHandler.GetAdvertisement)
	log.Info("registered route", logger.Field("path", "/advertisements/{id}"), logger.Field("method", "GET"))
	router.With(authMiddleware).Put("/advertisements/{id}", adsHandler.UpdateAdvertisement)
	log.Info("registered route", logger.Field("path", "/advertisements/{id}"), logger.Field("method", "PUT"))
	router.With(authMiddleware).Delete("/advertisements/{id}", adsHandler.DeleteAdvertisement)
	log.Info("registered route", logger.Field("path", "/advertisements/{id}"), logger.Field("method", "DELETE"))
	router.With(authMiddleware).Post("/advertisements/{id}/images", adsHandler.AddAdImages)
	log.Info("registered route", logger.Field("path", "/advertisements/{id}/images"), logger.Field("method", "POST"))
	router.With(authMiddleware).Delete("/advertisements/{ad_id}/images/{image_id}", adsHandler.DeleteAdImage)
	log.Info("registered route", logger.Field("path", "/advertisements/{ad_id}/images/{image_id}"), logger.Field("method", "DELETE"))
	router.With(authMiddleware).Get("/advertisements/my", adsHandler.GetMyAdvertisements)
	log.Info("registered route", logger.Field("path", "/advertisements/my"), logger.Field("method", "GET"))

	// static
	router.Handle(cfg.BaseURL+"*", http.StripPrefix(cfg.BaseURL, http.FileServer(http.Dir(cfg.ImageStoragePath))))
	log.Info("registered static route", logger.Field("path", cfg.BaseURL+"*"), logger.Field("method", "GET"))

}
