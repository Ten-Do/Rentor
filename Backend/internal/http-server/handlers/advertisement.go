package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"rentor/internal/http-server/middleware"
	"rentor/internal/logger"
	"rentor/internal/models"
	"rentor/internal/service"

	"github.com/go-chi/chi/v5"
)

type AdvertisementHandlers struct {
	adService service.AdvertisementService
	imageSvc  service.ImageService
}

func NewAdvertisementHandlers(adService service.AdvertisementService, imageSvc service.ImageService) *AdvertisementHandlers {
	return &AdvertisementHandlers{
		adService: adService,
		imageSvc:  imageSvc,
	}
}

// ===========================
// CREATE
// ===========================
func (h *AdvertisementHandlers) CreateAdvertisement(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	var input models.CreateAdvertisementInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}

	res, err := h.adService.CreateAdvertisement(userID, &input)
	if err != nil {
		logger.Error("create ad failed", logger.Field("error", err.Error()))
		http.Error(w, `{"error":"cannot create advertisement"}`, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, res)
}

// ===========================
// GET /advertisements/{id}
// ===========================
func (h *AdvertisementHandlers) GetAdvertisement(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	ad, err := h.adService.GetAdvertisement(id)
	if err != nil {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, ad)
}

// ===========================
// LIST WITH FILTERS
// ===========================
func (h *AdvertisementHandlers) ListAdvertisements(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	filters := &models.AdFilters{
		Page:  parseIntDefault(q.Get("page"), 1),
		Limit: parseIntDefault(q.Get("limit"), 20),
	}

	if v := q.Get("minPrice"); v != "" {
		val := parseFloatPointer(v)
		filters.MinPrice = val
	}
	if v := q.Get("maxPrice"); v != "" {
		val := parseFloatPointer(v)
		filters.MaxPrice = val
	}
	if v := q.Get("type"); v != "" {
		filters.Type = &v
	}
	if v := q.Get("rooms"); v != "" {
		filters.Rooms = &v
	}
	if v := q.Get("city"); v != "" {
		filters.City = &v
	}
	if v := q.Get("keywords"); v != "" {
		filters.Keywords = &v
	}

	list, err := h.adService.GetAdvertisementsPaged(filters)
	if err != nil {
		http.Error(w, `{"error":"failed to fetch advertisements"}`, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, list)
}

// ===========================
// /advertisements/my
// ===========================
func (h *AdvertisementHandlers) GetMyAdvertisements(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	q := r.URL.Query()
	page := parseIntDefault(q.Get("page"), 1)
	limit := parseIntDefault(q.Get("limit"), 20)

	list, err := h.adService.GetMyAdvertisements(userID, page, limit)
	if err != nil {
		http.Error(w, `{"error":"failed to fetch my ads"}`, http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, list)
}

// ===========================
// UPDATE
// ===========================
func (h *AdvertisementHandlers) UpdateAdvertisement(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input models.UpdateAdvertisementInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}

	err = h.adService.UpdateAdvertisement(adID, userID, &input)
	if err != nil {
		http.Error(w, `{"error":"cannot update advertisement"}`, http.StatusForbidden)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "updated"})
}

// ===========================
// DELETE
// ===========================
func (h *AdvertisementHandlers) DeleteAdvertisement(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err = h.adService.DeleteAdvertisement(userID, adID)
	if err != nil {
		http.Error(w, `{"error":"cannot delete"}`, http.StatusForbidden)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// ===========================
// ADD IMAGES
// ===========================
func (h *AdvertisementHandlers) AddAdImages(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	// Parse multipart form
	r.ParseMultipartForm(20 << 20) // 20 MB limit

	files := r.MultipartForm.File["images"]
	if len(files) == 0 {
		http.Error(w, `{"error":"no images provided"}`, http.StatusBadRequest)
		return
	}

	// Сохраняем изображения через ImageService
	urls, err := h.imageSvc.SaveAdvertisementImages(adID, files)
	if err != nil {
		http.Error(w, `{"error":"cannot save images"}`, http.StatusInternalServerError)
		return
	}

	// Передаём в AdvertisementService для сохранения URL в БД
	resp, err := h.adService.AddImages(userID, adID, urls)
	if err != nil {
		http.Error(w, `{"error":"cannot link images"}`, http.StatusForbidden)
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// ===========================
// DELETE IMAGE
// ===========================
func (h *AdvertisementHandlers) DeleteAdImage(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	adID, _ := strconv.Atoi(chi.URLParam(r, "ad_id"))
	imgID, _ := strconv.Atoi(chi.URLParam(r, "image_id"))

	err = h.adService.DeleteImage(userID, adID, imgID)
	if err != nil {
		http.Error(w, `{"error":"delete failed"}`, http.StatusForbidden)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "image deleted"})
}

// ===========================
// Helpers
// ===========================
func parseIntDefault(s string, def int) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	return def
}

func parseFloatPointer(s string) *float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &val
}
