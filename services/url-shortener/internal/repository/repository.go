package repository

import "github.com/Insid1/with-auth/url-shortener/internal/repository/shortener"

type ShortenerRepository interface {
	Get(shortenURL string) (string, error)
	Set(URL string) (string, error)
	Delete(shortenURL string) error
}

func NewShortenerRepository() ShortenerRepository {
	return &shortener.Repository{}
}
