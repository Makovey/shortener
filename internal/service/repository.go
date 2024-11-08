package service

import "github.com/Makovey/shortener/internal/api/model"

type Shortener interface {
	Store(shortURL, longURL string) error
	Get(shortURL string) (string, error)
	StoreBatch(models []model.ShortenBatch) error
}
