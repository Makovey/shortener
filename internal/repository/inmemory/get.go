package inmemory

import (
	"context"
	"fmt"

	"github.com/Makovey/shortener/internal/repository"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
	"github.com/Makovey/shortener/internal/service/model"
)

// GetFullURL возвращает полный урл по короткому урлу, если он есть в памяти
func (r *Repo) GetFullURL(ctx context.Context, shortURL, userID string) (*repoModel.UserURL, error) {
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

// GetUserURLs возвращает все урлы, которые есть в памяти
func (r *Repo) GetUserURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	models := make([]model.ShortenBatch, 0)
	for _, val := range r.storage {
		models = append(models, model.ShortenBatch{ShortURL: val.originalURL, OriginalURL: val.originalURL})
	}

	return models, nil
}

// GetStats возвращает стистику по сервису, количество пользователей и сокращенных адресов
func (r *Repo) GetStats(ctx context.Context) (model.Stats, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make(map[string]bool)
	shortURLs := make(map[string]bool)

	for _, row := range r.storage {
		users[row.userID] = true
		shortURLs[row.shortURL] = true
	}

	return model.Stats{Users: len(users), URLS: len(shortURLs)}, nil
}
