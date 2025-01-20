package inmemory

import (
	"context"
	"fmt"

	"github.com/Makovey/shortener/internal/repository"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
	"github.com/Makovey/shortener/internal/service/model"
)

func (r *repo) GetFullURL(ctx context.Context, shortURL, userID string) (*repoModel.UserURL, error) {
	fn := "inmemory.GetFullURL"

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, row := range r.storage {
		if row.shortURL == shortURL {
			return &repoModel.UserURL{OriginalURL: row.originalURL, IsDeleted: row.isDeleted}, nil
		}
	}

	return nil, fmt.Errorf("[%s]: %w", fn, repository.ErrURLNotFound)
}

func (r *repo) GetUserURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	models := make([]model.ShortenBatch, 0)
	for _, val := range r.storage {
		models = append(models, model.ShortenBatch{ShortURL: val.originalURL, OriginalURL: val.originalURL})
	}

	return models, nil
}
