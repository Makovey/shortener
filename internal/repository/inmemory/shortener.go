package inmemory

import (
	"fmt"

	"github.com/Makovey/shortener/internal/service"
)

type repo struct {
	storage map[string]string
}

func (r *repo) Store(shortURL, longURL string) {
	if r.storage == nil {
		r.storage = make(map[string]string)
	}

	r.storage[shortURL] = longURL
}

func (r *repo) Get(shortURL string) (string, error) {
	longURL, ok := r.storage[shortURL]
	if !ok {
		return "", fmt.Errorf("long url by -> %s not found", shortURL)
	}

	return longURL, nil
}

func NewRepositoryInMemory() service.Shortener {
	return &repo{}
}
