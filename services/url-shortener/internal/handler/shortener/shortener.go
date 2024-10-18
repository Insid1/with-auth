package shortener

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Insid1/with-auth/url-shortener/internal/service"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Handler struct {
	ShortenerService service.ShortenerService
}

func (h *Handler) Get(w http.ResponseWriter, req *http.Request) {
	shortenURL := chi.URLParam(req, "shortenURL")

	doc, err := h.ShortenerService.Get(shortenURL)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, err.Error(), http.StatusNotFound)

			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	h.sendJSON(w, doc)
}

func (h *Handler) GetLink(w http.ResponseWriter, req *http.Request) {
	shortenURL := chi.URLParam(req, "shortenURL")

	doc, err := h.ShortenerService.Get(shortenURL)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, err.Error(), http.StatusNotFound)

			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, err = w.Write([]byte(doc.OriginalURL))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Redirect(w http.ResponseWriter, req *http.Request) {
	shortenURL := chi.URLParam(req, "shortenURL")

	doc, err := h.ShortenerService.Get(shortenURL)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, err.Error(), http.StatusNotFound)

			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	http.Redirect(w, req, doc.OriginalURL, http.StatusTemporaryRedirect)
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
