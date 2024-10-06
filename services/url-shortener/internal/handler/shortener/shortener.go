package shortener

import "net/http"

type Handler struct{}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Set(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
}
