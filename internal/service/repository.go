package service

import "github.com/Makovey/shortener/internal/api/model"

type Shortener interface {
	Store(shortURL, longURL, userID string) error
	Get(shortURL, userID string) (string, error)
	StoreBatch(models []model.ShortenBatch, userID string) error
	GetAll(userID string) ([]model.ShortenBatch, error)
}
