package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type URLDocument struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ShortID     string        `bson:"short_id"      json:"shortId,omitempty"`
	OriginalURL string        `bson:"original_url"  json:"originalUrl,omitempty"`
	CreatedAt   time.Time     `bson:"created_at"    json:"createdAt,omitempty"`
	UpdatedAt   time.Time     `bson:"updated_at"    json:"updatedAt,omitempty"`
}
