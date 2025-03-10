package service

import (
	"context"

	comModel "github.com/Makovey/shortener/internal/service/model"
	"github.com/Makovey/shortener/internal/transport/model"
)

// Service интерфейс он же useCase, слой отвечающий за бизнес-логику приложения
type Service interface {
	CreateShortURL(ctx context.Context, url, userID string) (string, error)
	GetFullURL(ctx context.Context, shortURL, userID string) (model.UserFullURL, error)
	ShortBatch(ctx context.Context, batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error)
	GetAllURLs(ctx context.Context, userID string) ([]comModel.ShortenBatch, error)
	DeleteUsersURLs(ctx context.Context, userID string, shortURLs []string) []error
	GetStats(ctx context.Context) (comModel.Stats, error)
}
