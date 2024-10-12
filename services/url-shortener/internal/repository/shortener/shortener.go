package shortener

import (
	"context"
	"fmt"

	"github.com/Insid1/with-auth/url-shortener/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository struct {
	DB *mongo.Database
}

func (r *Repository) Get(shortenID string) (*model.URLDocument, error) {
	collection := r.DB.Collection("urls")

	// Ищем документ по short_id
	filter := bson.M{"short_id": shortenID}

	var result model.URLDocument

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) Set(doc *model.URLDocument) (*model.URLDocument, error) {
	collection := r.DB.Collection("urls")

	// Вставляем документ
	result, err := collection.InsertOne(context.TODO(), *doc)
	if err != nil {
		return nil, err
	}

	oid, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, fmt.Errorf("cannot parse id: %v", result.InsertedID)
	}

	// Обогащаем URLDocument id
	doc.ID = oid

	return doc, nil
}

func (r *Repository) Delete(shortenID string) error {
	collection := r.DB.Collection("urls")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"short_id": shortenID})
	if err != nil {
		return err
	}

	return nil
}
