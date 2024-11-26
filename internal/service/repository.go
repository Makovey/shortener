package service

import (
	"context"

	"github.com/Makovey/shortener/internal/api/model"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
)

type Shortener interface {
	Store(shortURL, longURL, userID string) error
	Get(shortURL, userID string) (repoModel.ShortenGet, error)
	StoreBatch(models []model.ShortenBatch, userID string) error
	GetAll(userID string) ([]model.ShortenBatch, error)
	DeleteUsersURL(ctx context.Context, userID string, url string) error
}
