package api

import (
	"context"

	"github.com/Makovey/shortener/internal/api/model"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
)

type Shortener interface {
	Short(url, userID string) (string, error)
	Get(shortURL, userID string) (repoModel.ShortenGet, error)
	ShortBatch(batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error)
	GetAll(userID string) ([]model.ShortenBatch, error)
	DeleteUsersURLS(ctx context.Context, userID string, shortURLs []string) []error
}

type Checker interface {
	CheckPing() error
}
