package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/Insid1/with-auth/url-shortener/internal/config"
	"github.com/Insid1/with-auth/url-shortener/internal/model"
	"github.com/Insid1/with-auth/url-shortener/internal/repository"
)

const (
	shortIDLength = 10
)

type Service struct {
	ShortenerRepo repository.ShortenerRepository
	Config        *config.Config
}

func (s *Service) Get(shortenID string) (*model.URLDocument, error) {
	return s.ShortenerRepo.Get(shortenID)
}

func (s *Service) Set(totalURL string) (*model.URLDocument, error) {
	shortID, err := s.generateShortID(shortIDLength)
	if err != nil {
		return nil, err
	}

	doc := &model.URLDocument{
		ID:          [12]byte{},
		ShortID:     shortID,
		OriginalURL: totalURL,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	updatedDoc, err := s.ShortenerRepo.Set(doc)
	if err != nil {
		return nil, err
	}

	return updatedDoc, nil
}

func (s *Service) Delete(shortenID string) error {
	return s.ShortenerRepo.Delete(shortenID)
}

func (s *Service) GenerateURL(doc *model.URLDocument) (string, error) {
	if s.Config.GetRedirectServiceAddress() == "" {
		return "", fmt.Errorf("redirect url is not provided")
	}

	if doc.ShortID == "" {
		return "", fmt.Errorf("invalid id")
	}

	rawShortenURL, err := url.JoinPath(s.Config.GetRedirectServiceAddress(), doc.ShortID)
	if err != nil {
		return "", fmt.Errorf("unable to join url")
	}

	shortenURL, _ := strings.CutPrefix(rawShortenURL, "//")

	return shortenURL, nil
}

func (s *Service) generateShortID(length int) (string, error) {
	// Создаем байтовый массив для генерации случайной строки
	bArr := make([]byte, length)

	_, err := rand.Read(bArr)
	if err != nil {
		return "", err
	}

	// Кодируем в base64 URL-friendly формат
	shortID := base64.URLEncoding.EncodeToString(bArr)

	// Обрезаем до нужной длины
	return shortID[:length], nil
}
