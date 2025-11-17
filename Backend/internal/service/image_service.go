package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type imageService struct {
	StoragePath string // например, "./storage"
	BaseURL     string // например, "/static/"
}

// NewimageService создаёт сервис для работы с изображениями
func NewimageService(storagePath, baseURL string) *imageService {
	return &imageService{
		StoragePath: storagePath,
		BaseURL:     baseURL,
	}
}

// SaveAdvertisementImages сохраняет массив файлов для объявления и возвращает URL
func (s *imageService) SaveAdvertisementImages(adID int, files []*multipart.FileHeader) ([]string, error) {
	var urls []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Генерация уникального имени файла
		timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
		ext := filepath.Ext(fileHeader.Filename)
		filename := fmt.Sprintf("ad_%d_%s%s", adID, timestamp, ext)

		savePath := filepath.Join(s.StoragePath, filename)

		dst, err := os.Create(savePath)
		if err != nil {
			return nil, err
		}

		if _, err := io.Copy(dst, file); err != nil {
			dst.Close()
			return nil, err
		}
		dst.Close()

		urls = append(urls, s.BaseURL+filename)
	}

	return urls, nil
}
