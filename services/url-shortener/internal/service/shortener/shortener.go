package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"
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

func (s *Service) Get(shortenURL string) *model.URLDocument {
	return s.ShortenerRepo.Get(shortenURL)
}

func (s *Service) Set(totalURL string, urlPrefix string) (*model.URLDocument, error) {
	// Если не передан URLPrefix назначаем по умолчанию тот, что в переменных окружения
	if urlPrefix == "" {
		urlPrefix = s.Config.ShortenerURLPrefix
	}

	shortID, err := s.generateShortID(shortIDLength)
	if err != nil {
		return nil, err
	}

	shortURL, err := url.JoinPath(urlPrefix, shortID)
	if err != nil {
		return nil, err
	}

	return s.ShortenerRepo.Set(&model.URLDocument{
		ID:          [12]byte{},
		ShortURL:    shortURL,
		OriginalURL: totalURL,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	})
}

func (s *Service) Delete(shortenURL string) error {
	return s.ShortenerRepo.Delete(shortenURL)
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
