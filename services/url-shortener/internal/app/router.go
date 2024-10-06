package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Insid1/with-auth/pkg/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func (a *App) getRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(chiMiddleware.RequestID)
	router.Use(middleware.GetLoggerMiddleware(a.Logger))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("sup"))
	})

	router.Get("/slow", func(w http.ResponseWriter, r *http.Request) {
		// Simulates some hard work.
		//
		// We want this handler to complete successfully during a shutdown signal,
		// so consider the work here as some background routine to fetch a long running
		// search query to find as many results as possible, but, instead we cut it short
		// and respond with what we have so far. How a shutdown is handled is entirely
		// up to the developer, as some code blocks are preemptible, and others are not.
		time.Sleep(5 * time.Second)

		w.Write([]byte(fmt.Sprintf("all done.\n")))
	})

	return router
}
