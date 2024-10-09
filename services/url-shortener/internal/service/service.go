package service

import (
	"github.com/Insid1/with-auth/url-shortener/internal/config"
	"github.com/Insid1/with-auth/url-shortener/internal/model"
	"github.com/Insid1/with-auth/url-shortener/internal/repository"
	"github.com/Insid1/with-auth/url-shortener/internal/service/shortener"
)

type ShortenerService interface {
	Get(shortenURL string) *model.URLDocument
	Set(totalURL string, URLPrefix string) (*model.URLDocument, error)
	Delete(shortenURL string) error
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
