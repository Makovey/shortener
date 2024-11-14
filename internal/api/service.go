package api

import "github.com/Makovey/shortener/internal/api/model"

type Shortener interface {
	Short(url string) (string, error)
	Get(shortURL string) (string, error)
	ShortBatch(batch []model.ShortenBatchRequest) ([]model.ShortenBatchResponse, error)
}

type Checker interface {
	CheckPing() error
}
