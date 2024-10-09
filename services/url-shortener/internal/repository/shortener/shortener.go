package shortener

import (
	"context"

	"github.com/Insid1/with-auth/url-shortener/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	DB *mongo.Database
}

func (r *Repository) Get(shortenURL string) *model.URLDocument {
	collection := r.DB.Collection("urls")

	// Ищем документ по short_url
	filter := bson.M{"short_url": shortenURL}

	var result model.URLDocument

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return &result
	}

	return nil
}

func (r *Repository) Set(doc *model.URLDocument) (*model.URLDocument, error) {
	collection := r.DB.Collection("urls")

	// Вставляем документ
	result, err := collection.InsertOne(context.TODO(), *doc)
	if err != nil {
		return nil, err
	}

	parsedID, ok := result.InsertedID.([12]byte)
	if !ok {
		return nil, err
	}

	// Обогащаем URLDocument id
	doc.ID = parsedID

	return doc, nil
}

func (r *Repository) Delete(shortenURL string) error {
	collection := r.DB.Collection("urls")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"short_url": shortenURL})
	if err != nil {
		return err
	}

	return nil
}
