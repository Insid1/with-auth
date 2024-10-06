package handler

import (
	"net/http"

	"github.com/Insid1/with-auth/url-shortener/internal/handler/shortener"
)

type ShortenerHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Set(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewShortenerHandler() ShortenerHandler {
	return &shortener.Handler{}
}
