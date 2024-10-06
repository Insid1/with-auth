package app

import (
	"database/sql"

	"github.com/Insid1/with-auth/url-shortener/internal/config"
)

type Provider struct {
	config *config.Config
	db     *sql.DB

	// shortenerHandler    handler.Auth
	// shortenerService    service.Auth
	// shortenerRepository repository.Auth
}
