package service

import (
	"errors"
	"rentor/internal/models"
	"rentor/internal/repository"
)

type advertisementService struct {
	adRepo repository.AdRepository
}

func NewadvertisementService(adRepo repository.AdRepository) *advertisementService {
	return &advertisementService{
		adRepo: adRepo,
	}
}

// ==========================
// CREATE
// ==========================
func (s *advertisementService) CreateAdvertisement(userID int, input *models.CreateAdvertisementInput) (*models.GetAd, error) {
	adID, err := s.adRepo.CreateAdvertisement(userID, input)
	if err != nil {
		return nil, err
	}

	return s.adRepo.GetAdvertisement(adID)
}

// ==========================
// GET BY ID
// ==========================
func (s *advertisementService) GetAdvertisement(id int) (*models.GetAd, error) {
	return s.adRepo.GetAdvertisement(id)
}

// ==========================
// FILTERED LIST
// ==========================
func (s *advertisementService) GetAdvertisementsPaged(filters *models.AdFilters) (*models.GetAdPreviewsList, error) {
	return s.adRepo.GetAdvertisementsPaged(filters)
}

// ==========================
// GET MY ADS
// ==========================
func (s *advertisementService) GetMyAdvertisements(userID, page, limit int) (*models.GetAdPreviewsList, error) {
	f := &models.AdFilters{
		Page:   page,
		Limit:  limit,
		UserID: &userID,
	}
	return s.adRepo.GetAdvertisementsPaged(f)
}

// ==========================
// UPDATE
// ==========================
func (s *advertisementService) UpdateAdvertisement(userID, adID int, input *models.UpdateAdvertisementInput) error {
	// Проверка принадлежности
	owner, err := s.adRepo.GetUserID(adID)
	if err != nil {
		return err
	}

	if userID != owner {
		return errors.New("not owner")
	}

	return s.adRepo.UpdateAdvertisement(adID, input)
}

// ==========================
// DELETE
// ==========================
func (s *advertisementService) DeleteAdvertisement(userID, adID int) error {
	// Проверка принадлежности
	owner, err := s.adRepo.GetUserID(adID)
	if err != nil {
		return err
	}

	if userID != owner {
		return errors.New("not owner")
	}

	// TODO — привязать к user_profile_id, когда ты сделаешь таблицу профилей
	// сейчас у нас нет прямой связи — значит, удаляем без проверки

	return s.adRepo.DeleteAdvertisement(adID)
}

// ==========================
// ADD IMAGES
// ==========================
func (s *advertisementService) AddImages(userID, adID int, urls []string) (*models.ImagesUploadResponse, error) {
	// Проверка принадлежности
	owner, err := s.adRepo.GetUserID(adID)
	if err != nil {
		return nil, err
	}

	if userID != owner {
		return nil, errors.New("not owner")
	}

	err = s.adRepo.CreateAdvertisementImages(adID, urls)
	if err != nil {
		return nil, err
	}

	return &models.ImagesUploadResponse{
		Uploaded: urls,
		Count:    len(urls),
	}, nil
}

// ==========================
// DELETE IMAGE
// ==========================
func (s *advertisementService) DeleteImage(userID, adID, imageID int) error {
	// Проверка принадлежности
	owner, err := s.adRepo.GetUserID(adID)
	if err != nil {
		return err
	}

	if userID != owner {
		return errors.New("not owner")
	}

	return s.adRepo.DeleteAdvertisementImage(adID, imageID)
}

func (s *advertisementService) GetImagePath(adID, imageID int) (string, error) {
	return s.adRepo.GetImagePath(adID, imageID)
}
