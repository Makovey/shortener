package service

import (
	"context"

	"github.com/Makovey/shortener/internal/api/model"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
)

type Shortener interface {
	SaveUserURL(ctx context.Context, shortURL, longURL, userID string) error
	GetFullURL(ctx context.Context, shortURL, userID string) (repoModel.UserURL, error)
	SaveUserURLs(ctx context.Context, models []model.ShortenBatch, userID string) error
	GetUserURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error)
	MarkURLAsDeleted(ctx context.Context, userID string, url string) error
}
