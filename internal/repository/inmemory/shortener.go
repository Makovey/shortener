package inmemory

import (
	"context"
	"sync"

	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/repository"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
	"github.com/Makovey/shortener/internal/service"
)

type repo struct {
	storage []storage
	mu      sync.RWMutex
}

func NewRepositoryInMemory() service.Shortener {
	return &repo{
		storage: make([]storage, 0),
		mu:      sync.RWMutex{},
	}
}

type storage struct {
	originalURL string
	shortURL    string
	userID      string
	isDeleted   bool
}

func (r *repo) SaveUserURL(ctx context.Context, shortURL, longURL, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage = append(r.storage, storage{shortURL: shortURL, originalURL: longURL, userID: userID, isDeleted: false})
	return nil
}

func (r *repo) GetFullURL(ctx context.Context, shortURL, userID string) (repoModel.UserURL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, row := range r.storage {
		if row.shortURL == shortURL {
			return repoModel.UserURL{OriginalURL: row.originalURL, IsDeleted: row.isDeleted}, nil
		}
	}

	return repoModel.UserURL{}, repository.ErrURLNotFound
}

func (r *repo) SaveUserURLs(ctx context.Context, models []model.ShortenBatch, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, m := range models {
		r.storage = append(r.storage, storage{originalURL: m.OriginalURL, userID: userID})
	}

	return nil
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

func (r *repo) MarkURLAsDeleted(ctx context.Context, userID string, url string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, row := range r.storage {
		if row.shortURL == url && row.userID == userID {
			r.storage[i].isDeleted = true
		}
	}

	return nil
}
