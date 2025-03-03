package disc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Makovey/shortener/internal/repository"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
	"github.com/Makovey/shortener/internal/service/model"
)

// GetFullURL возвращает полный урл по короткому урлу, если он есть в файле
func (r *Repo) GetFullURL(ctx context.Context, shortURL, userID string) (*repoModel.UserURL, error) {
	fn := "disc.GetFullURL"

	r.mu.RLock()
	defer r.mu.RUnlock()

	shortenerURLs, err := r.fetchAllURLs()
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	for _, u := range shortenerURLs {
		if u.ShortURL == shortURL {
			return &repoModel.UserURL{OriginalURL: u.OriginalURL, IsDeleted: u.IsDeleted}, nil
		}
	}

	return nil, fmt.Errorf("[%s]: %w", fn, repository.ErrURLNotFound)
}

// GetUserURLs возвращает все урлы, которые есть в файле
func (r *Repo) GetUserURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error) {
	fn := "disc.GetUserURLs"

	r.mu.RLock()
	defer r.mu.RUnlock()

	models := make([]model.ShortenBatch, 0)

	urls, err := r.fetchAllURLs()
	if err != nil {
		return models, fmt.Errorf("[%s]: %w", fn, err)
	}

	for _, url := range urls {
		models = append(models, model.ShortenBatch{ShortURL: url.ShortURL, OriginalURL: url.OriginalURL})
	}

	return models, nil
}

func (r *Repo) fetchAllURLs() ([]repoModel.ShortenerURL, error) {
	fn := "disc.fetchAllURLs"

	var urls []repoModel.ShortenerURL
	b, err := os.ReadFile(r.path)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	for _, line := range bytes.Split(b, []byte("\n")) {
		if len(line) == 0 {
			break
		}
		var url repoModel.ShortenerURL
		err = json.Unmarshal(line, &url)
		if err != nil {
			r.log.Debug(fmt.Sprintf("[%s]: %s", fn, err.Error()))
			continue
		}
		urls = append(urls, url)
	}

	return urls, nil
}

// GetStats возвращает стистику по сервису, количество пользователей и сокращенных адресов
func (r *Repo) GetStats(ctx context.Context) (model.Stats, error) {
	fn := "disc.GetStats"

	r.mu.RLock()
	defer r.mu.RUnlock()

	models, err := r.fetchAllURLs()
	if err != nil {
		return model.Stats{}, fmt.Errorf("[%s]: %w", fn, err)
	}

	users := make(map[string]bool)
	shortURLs := make(map[string]bool)

	for _, u := range models {
		users[u.OwnerID] = true
		shortURLs[u.ShortURL] = true
	}

	return model.Stats{Users: len(users), URLS: len(shortURLs)}, nil
}
