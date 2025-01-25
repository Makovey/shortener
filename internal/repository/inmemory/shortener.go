package inmemory

import (
	"sync"
)

// Repo репозиторий inmemmory
type Repo struct {
	storage []storage
	mu      sync.RWMutex
}

// NewRepositoryInMemory конструктор для репозитрия inmemory
func NewRepositoryInMemory() *Repo {
	return &Repo{
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
