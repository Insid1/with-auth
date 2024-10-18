package handler

import (
	"net/http"

	"github.com/Insid1/with-auth/url-shortener/internal/handler/shortener"
	"github.com/Insid1/with-auth/url-shortener/internal/service"
)

type ShortenerHandler interface {
	GetLink(w http.ResponseWriter, r *http.Request)
	Redirect(w http.ResponseWriter, r *http.Request)

	Get(w http.ResponseWriter, r *http.Request)
	Set(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewShortenerHandler(shortenerService service.ShortenerService) ShortenerHandler {
	return &shortener.Handler{
		ShortenerService: shortenerService,
	}
}
