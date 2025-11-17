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
	userService        service.UserService
	userProfileService service.UserProfileService
}

func NewUserProfileHandler(userService service.UserService, userProfileService service.UserProfileService) *UserProfileHandler {
	return &UserProfileHandler{
		userService:        userService,
		userProfileService: userProfileService,
	}
}

// GetUserProfile handles GET /user/profile
func (h *UserProfileHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "user not authenticated")
		return
	}

	user, err := h.userService.GetUser(userID)
	if err != nil {
		logger.Error("failed to get user", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	profile, err := h.userProfileService.GetUserProfile(userID)
	if err != nil {
		logger.Error("failed to get user profile", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		writeError(w, http.StatusNotFound, "profile not found")
		return
	}

	logger.Info("GetUserProfile called", logger.Field("user_id", userID))
	writeJSON(w, http.StatusOK, &models.GetUserProfileOutput{
		UserID:     user.UserID,
		Email:      user.Email,
		Phone:      user.Phone,
		FirstName:  profile.FirstName,
		Surname:    profile.Surname,
		Patronymic: profile.Patronymic,
		CreatedAt:  profile.CreatedAt,
	})
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

	err = h.userProfileService.UpdateUserProfile(userID, &req)
	if err != nil {
		logger.Error("failed to update user profile", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		writeError(w, http.StatusInternalServerError, "failed to update profile")
		return
	}

	user, err := h.userService.GetUser(userID)
	if err != nil {
		logger.Error("failed to get user", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	profile, err := h.userProfileService.GetUserProfile(userID)
	if err != nil {
		logger.Error("failed to get user profile", logger.Field("error", err.Error()), logger.Field("user_id", userID))
		writeError(w, http.StatusNotFound, "profile not found")
		return
	}

	logger.Info("UpdateUserProfile called", logger.Field("user_id", userID))
	writeJSON(w, http.StatusOK, &models.GetUserProfileOutput{
		UserID:     user.UserID,
		Email:      user.Email,
		Phone:      user.Phone,
		FirstName:  profile.FirstName,
		Surname:    profile.Surname,
		Patronymic: profile.Patronymic,
		CreatedAt:  profile.CreatedAt,
	})
}
