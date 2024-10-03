package in_memory

import (
	"fmt"
	"github.com/Makovey/shortener/internal/repository"
)

type repo struct {
	storage map[string]string
}

func (r *repo) Store(shortUrl, longUrl string) {
	if r.storage == nil {
		r.storage = make(map[string]string)
	}

	r.storage[shortUrl] = longUrl
}

func (r *repo) Get(shortUrl string) (string, error) {
	longUrl, ok := r.storage[shortUrl]
	if !ok {
		return "", fmt.Errorf("long url by -> %s not found", shortUrl)
	}

	return longUrl, nil
}

func NewRepositoryInMemory() repository.ShortenerRepository {
	return &repo{}
}
