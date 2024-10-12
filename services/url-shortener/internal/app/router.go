package app

import (
	"net/http"

	"github.com/Insid1/with-auth/pkg/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func (p *Provider) getRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(chiMiddleware.RequestID)
	router.Use(middleware.GetLoggerMiddleware(p.Logger))

	router.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Route("/short", func(r chi.Router) {
		shortenerHandler := p.GetShortenerHandler()

		r.Post("/", shortenerHandler.Set)

		r.Route("/{shortenURL}", func(r chi.Router) {
			r.Get("/", shortenerHandler.Get)
			r.Delete("/", shortenerHandler.Delete)
		})
	})

	return router
}
