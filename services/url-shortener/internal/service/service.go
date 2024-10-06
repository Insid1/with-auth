package service

import "github.com/Insid1/with-auth/url-shortener/internal/service/shortener"

type ShortenerService interface {
	Get(shortenURL string) (string, error)
	Set(totalURL string) (string, error)
	Delete(shortenURL string) error
}

func NewShortenerService() ShortenerService {
	return &shortener.Service{}
}
