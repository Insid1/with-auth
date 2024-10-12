package shortener

import "github.com/Insid1/with-auth/url-shortener/internal/model"

type SetReqBody struct {
	OriginalURL string `json:"originalUrl"`
}

type SetResBody struct {
	*model.URLDocument
	ShortenURL string `json:"shortenUrl"`
}
