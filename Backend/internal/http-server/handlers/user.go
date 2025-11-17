package handlers

import (
	"encoding/json"
	"net/http"

	"rentor/internal/http-server/middleware"
	"rentor/internal/logger"
	"rentor/internal/models"
	"rentor/internal/service"
)

type UserProfileHandler struct {
	service service.UserProfileService
}

func NewUserProfileHandler(svc service.UserProfileService) *UserProfileHandler {
	return &UserProfileHandler{
		service: svc,
	}
}

// GetUserProfile handles GET /user/profile
func (h *UserProfileHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	profile, err := h.service.GetUserProfile(userID)
	if err != nil {
		logger.Error("failed to get user profile", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		writeError(w, http.StatusNotFound, "profile not found")
		return
	}

	logger.Info("GetUserProfile called", logger.Field("user_id", userID))
	writeJSON(w, http.StatusOK, profile)
}

// UpdateUserProfile handles PUT /user/profile
func (h *UserProfileHandler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	var req models.UpdateUserProfileInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	err = h.service.UpdateUserProfile(userID, &req)
	if err != nil {
		logger.Error("failed to update user profile", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		writeError(w, http.StatusInternalServerError, "failed to update profile")
		return
	}

	// Return updated profile
	profile, _ := h.service.GetUserProfile(userID)

	logger.Info("UpdateUserProfile called", logger.Field("user_id", userID))
	writeJSON(w, http.StatusOK, profile)
}
