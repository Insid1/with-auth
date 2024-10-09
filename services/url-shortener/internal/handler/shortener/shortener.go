package shortener

import (
	"net/http"

	"github.com/Insid1/with-auth/url-shortener/internal/service"
)

type Handler struct {
	ShortenerService service.ShortenerService
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
}
