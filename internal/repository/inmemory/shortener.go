package inmemory

import (
	"sync"

	"github.com/Makovey/shortener/internal/service/shortener"
)

type repo struct {
	storage []storage
	mu      sync.RWMutex
}

func NewRepositoryInMemory() shortener.Repository {
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
