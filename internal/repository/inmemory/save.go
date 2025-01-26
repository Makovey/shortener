package inmemory

import (
	"context"

	"github.com/Makovey/shortener/internal/service/model"
)

// SaveUserURL сохраняет полную информацию по урлу с UserID
func (r *Repo) SaveUserURL(ctx context.Context, shortURL, longURL, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage = append(r.storage, storage{shortURL: shortURL, originalURL: longURL, userID: userID, isDeleted: false})
	return nil
}

// SaveUserURLs сохраняет полную информацию по урлам с UserID
func (r *Repo) SaveUserURLs(ctx context.Context, models []model.ShortenBatch, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, m := range models {
		r.storage = append(r.storage, storage{originalURL: m.OriginalURL, userID: userID})
	}

	return nil
}
