package inmemory

import (
	"sync"
)

type Repo struct {
	storage []storage
	mu      sync.RWMutex
}

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
