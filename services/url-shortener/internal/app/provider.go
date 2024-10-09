package app

import (
	"github.com/Insid1/with-auth/url-shortener/internal/config"
	"github.com/Insid1/with-auth/url-shortener/internal/handler"
	"github.com/Insid1/with-auth/url-shortener/internal/repository"
	"github.com/Insid1/with-auth/url-shortener/internal/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type Provider struct {
	config *config.Config
	db     *mongo.Client
	Logger *zap.SugaredLogger

	shortenerHandler    handler.ShortenerHandler
	shortenerService    service.ShortenerService
	shortenerRepository repository.ShortenerRepository
}

func newProvider(
	config *config.Config,
	db *mongo.Client,
	logger *zap.SugaredLogger,
) *Provider {
	return &Provider{
		config: config,
		db:     db,
		Logger: logger,

		shortenerHandler:    nil,
		shortenerService:    nil,
		shortenerRepository: nil,
	}
}

func (p *Provider) GetShortenerHandler() handler.ShortenerHandler {
	if p.shortenerHandler == nil {
		p.shortenerHandler = handler.NewShortenerHandler(p.GetShortenerService())
	}

	return p.shortenerHandler
}

func (p *Provider) GetShortenerService() service.ShortenerService {
	if p.shortenerService == nil {
		p.shortenerService = service.NewShortenerService(p.GetShortenerRepo(), p.config)
	}

	return p.shortenerService
}

func (p *Provider) GetShortenerRepo() repository.ShortenerRepository {
	if p.shortenerRepository == nil {
		p.shortenerRepository = repository.NewShortenerRepository(p.db.Database("url_shortener_db"))
	}

	return p.shortenerRepository
}
