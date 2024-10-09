package repository

import (
	"github.com/Insid1/with-auth/url-shortener/internal/model"
	"github.com/Insid1/with-auth/url-shortener/internal/repository/shortener"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ShortenerRepository interface {
	Get(shortenURL string) *model.URLDocument
	Set(doc *model.URLDocument) (*model.URLDocument, error)
	Delete(shortenURL string) error
}

func NewShortenerRepository(db *mongo.Database) ShortenerRepository {
	return &shortener.Repository{
		DB: db,
	}
}
