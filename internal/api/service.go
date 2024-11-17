package api

import "github.com/Makovey/shortener/internal/api/model"

type Shortener interface {
	Short(url, userID string) (string, error)
	Get(shortURL, userID string) (string, error)
	ShortBatch(batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error)
	GetAll(userID string) ([]model.ShortenBatch, error)
}

type Checker interface {
	CheckPing() error
}
