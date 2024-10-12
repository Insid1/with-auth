package repository

import (
	"github.com/Insid1/with-auth/url-shortener/internal/model"
	"github.com/Insid1/with-auth/url-shortener/internal/repository/shortener"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ShortenerRepository interface {
	Get(shortenID string) (*model.URLDocument, error)
	Set(doc *model.URLDocument) (*model.URLDocument, error)
	Delete(shortenID string) error
}

func NewShortenerRepository(db *mongo.Database) ShortenerRepository {
	return &shortener.Repository{
		DB: db,
	}
}
