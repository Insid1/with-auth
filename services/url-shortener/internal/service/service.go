package service

import (
	"github.com/Insid1/with-auth/url-shortener/internal/config"
	"github.com/Insid1/with-auth/url-shortener/internal/model"
	"github.com/Insid1/with-auth/url-shortener/internal/repository"
	"github.com/Insid1/with-auth/url-shortener/internal/service/shortener"
)

type ShortenerService interface {
	Get(shortenID string) (*model.URLDocument, error)
	Set(totalURL string) (*model.URLDocument, error)
	Delete(shortenID string) error

	GenerateURL(doc *model.URLDocument) (string, error)
}

func NewShortenerService(
	shortenerRepo repository.ShortenerRepository,
	cfg *config.Config,
) ShortenerService {
	return &shortener.Service{
		ShortenerRepo: shortenerRepo,
		Config:        cfg,
	}
}
