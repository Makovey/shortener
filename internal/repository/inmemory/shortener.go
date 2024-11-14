package inmemory

import (
	"sync"

	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/repository"
	"github.com/Makovey/shortener/internal/service"
)

type repo struct {
	storage map[string]string
	mu      sync.RWMutex
}

func (r *repo) Store(shortURL, longURL string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage[shortURL] = longURL
	return nil
}

func (r *repo) Get(shortURL string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	longURL, ok := r.storage[shortURL]
	if !ok {
		return "", repository.ErrURLNotFound
	}

	return longURL, nil
}

func (r *repo) StoreBatch(models []model.ShortenBatch) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, m := range models {
		r.storage[m.ShortURL] = m.OriginalURL
	}

	return nil
}

func NewRepositoryInMemory() service.Shortener {
	return &repo{
		storage: make(map[string]string),
		mu:      sync.RWMutex{},
	}
}
