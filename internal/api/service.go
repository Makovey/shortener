package api

import (
	"context"

	"github.com/Makovey/shortener/internal/api/model"
)

type Shortener interface {
	Shorten(ctx context.Context, url, userID string) (string, error)
	GetFullURL(ctx context.Context, shortURL, userID string) (model.UserFullURL, error)
	ShortBatch(ctx context.Context, batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error)
	GetAllURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error)
	DeleteUsersURLs(ctx context.Context, userID string, shortURLs []string) []error
}

type Checker interface {
	CheckPing(ctx context.Context) error
}
