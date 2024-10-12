package shortener

import (
	"encoding/json"
	"net/http"

	"github.com/Insid1/with-auth/url-shortener/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	ShortenerService service.ShortenerService
}

func (h *Handler) Get(w http.ResponseWriter, req *http.Request) {
	shortenURL := chi.URLParam(req, "shortenURL")

	doc, err := h.ShortenerService.Get(shortenURL)
	if err != nil {
		// todo добавить проверку на "не найден"
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	h.sendJSON(w, doc)
}

func (h *Handler) Set(w http.ResponseWriter, req *http.Request) {
	var body SetReqBody

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	doc, err := h.ShortenerService.Set(body.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	ShortenURL, err := h.ShortenerService.GenerateURL(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	h.sendJSON(w, &SetResBody{
		URLDocument: doc,
		ShortenURL:  ShortenURL,
	})
}

func (h *Handler) Delete(w http.ResponseWriter, req *http.Request) {
	shortenURL := chi.URLParam(req, "shortenURL")

	err := h.ShortenerService.Delete(shortenURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) sendJSON(w http.ResponseWriter, valueToEncode interface{}) {
	respBody, err := json.Marshal(valueToEncode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(respBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
