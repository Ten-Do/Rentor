package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"rentor/internal/logger"

	"github.com/go-chi/chi/v5"
)

type adCreateReq struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Type        string   `json:"type"`
	Rooms       string   `json:"rooms"`
	City        string   `json:"city"`
	Address     string   `json:"address"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	Square      *float64 `json:"square,omitempty"`
}

// ListAdvertisements handles GET /advertisements
func ListAdvertisements(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page := 1
	limit := 20
	if p := q.Get("page"); p != "" {
		if n, err := strconv.Atoi(p); err == nil && n > 0 {
			page = n
		}
	}
	if l := q.Get("limit"); l != "" {
		if n, err := strconv.Atoi(l); err == nil && n > 0 {
			limit = n
		}
	}

	resp := map[string]interface{}{
		"data": []interface{}{},
		"pagination": map[string]interface{}{
			"page":        page,
			"limit":       limit,
			"total":       0,
			"total_pages": 0,
		},
	}

	logger.Info("ListAdvertisements called", logger.Field("page", page), logger.Field("limit", limit))
	writeJSON(w, http.StatusOK, resp)
}

// CreateAdvertisement handles POST /advertisements
func CreateAdvertisement(w http.ResponseWriter, r *http.Request) {
	var req adCreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	// Basic validation
	if req.Title == "" || req.Description == "" || req.Price < 0 || req.Type == "" || req.Rooms == "" || req.City == "" || req.Address == "" {
		writeError(w, http.StatusBadRequest, "missing required fields")
		return
	}

	ad := map[string]interface{}{
		"id":             strconv.FormatInt(time.Now().UnixNano(), 10),
		"title":          req.Title,
		"description":    req.Description,
		"price":          req.Price,
		"type":           req.Type,
		"rooms":          req.Rooms,
		"city":           req.City,
		"address":        req.Address,
		"latitude":       req.Latitude,
		"longitude":      req.Longitude,
		"square":         req.Square,
		"image_urls":     []string{},
		"landlord_name":  "John Doe",
		"landlord_email": "john.doe@example.com",
		"landlord_phone": "+70000000000",
		"status":         "active",
		"created_at":     time.Now().Format(time.RFC3339),
	}

	logger.Info("CreateAdvertisement called", logger.Field("title", req.Title))
	writeJSON(w, http.StatusCreated, ad)
}

// GetAdvertisement handles GET /advertisements/{id}
func GetAdvertisement(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// For now return 404 (not found)
	logger.Info("GetAdvertisement called", logger.Field("id", id))
	writeError(w, http.StatusNotFound, "advertisement not found")
}

// UpdateAdvertisement handles PUT /advertisements/{id}
func UpdateAdvertisement(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	logger.Info("UpdateAdvertisement called", logger.Field("id", id))
	// Not implemented - return 501
	writeError(w, http.StatusNotImplemented, "not implemented")
}

// DeleteAdvertisement handles DELETE /advertisements/{id}
func DeleteAdvertisement(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	logger.Info("DeleteAdvertisement called", logger.Field("id", id))
	// Deletion successful (stub) - return 204 No Content
	w.WriteHeader(http.StatusNoContent)
}

// AddAdImages handles POST /advertisements/{id}/images
func AddAdImages(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "image upload not implemented")
}

// DeleteAdImage handles DELETE /advertisements/{ad_id}/images/{image_id}
func DeleteAdImage(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "delete image not implemented")
}
