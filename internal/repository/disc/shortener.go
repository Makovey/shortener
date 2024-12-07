package disc

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"

	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/repository"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
	"github.com/Makovey/shortener/internal/service"
)

type repo struct {
	file   *os.File
	path   string
	writer *bufio.Writer
	log    logger.Logger
	mu     sync.RWMutex
}

func NewFileStorage(filePath string, log logger.Logger) service.Shortener {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Error(fmt.Sprintf("error opening disc: %v", err))
		panic(fmt.Sprintf("error opening disc: %v", err))
	}

	return &repo{
		file:   f,
		path:   filePath,
		writer: bufio.NewWriter(f),
		log:    log,
		mu:     sync.RWMutex{},
	}
}

func (r *repo) GetFullURL(ctx context.Context, shortURL, userID string) (repoModel.UserURL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	shortenerURLs := r.fetchAllURLs()
	for _, u := range shortenerURLs {
		if u.ShortURL == shortURL {
			return repoModel.UserURL{OriginalURL: u.OriginalURL, IsDeleted: u.IsDeleted}, nil
		}
	}

	return repoModel.UserURL{}, repository.ErrURLNotFound
}

func (r *repo) SaveUserURL(ctx context.Context, shortURL, longURL, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	currentURL := ShortenerURL{
		UUID:        uuid.New().String(),
		ShortURL:    shortURL,
		OriginalURL: longURL,
		OwnerID:     userID,
		IsDeleted:   false,
	}

	b, err := json.Marshal(&currentURL)
	if err != nil {
		r.log.Error(fmt.Sprintf("can't marshall shortener url cause: %s", err.Error()))
		return err
	}
	b = append(b, '\n')

	_, err = r.writer.Write(b)
	if err != nil {
		r.log.Error(fmt.Sprintf("can't write shortener url cause: %s", err.Error()))
		return err
	}
	_ = r.writer.Flush()

	return nil
}

func (r *repo) SaveUserURLs(ctx context.Context, models []model.ShortenBatch, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, m := range models {
		url := ShortenerURL{
			UUID:        m.CorrelationID,
			ShortURL:    m.ShortURL,
			OriginalURL: m.OriginalURL,
			OwnerID:     userID,
		}

		b, err := json.Marshal(&url)
		if err != nil {
			r.log.Error(fmt.Sprintf("can't marshall shortener url cause: %s", err.Error()))
			return err
		}
		b = append(b, '\n')

		_, err = r.writer.Write(b)
		if err != nil {
			r.log.Error(fmt.Sprintf("can't write shortener url cause: %s", err.Error()))
			return err
		}

		_ = r.writer.Flush()
	}

	return nil
}

func (r *repo) GetUserURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	models := make([]model.ShortenBatch, 0)

	urls := r.fetchAllURLs()
	for _, url := range urls {
		models = append(models, model.ShortenBatch{ShortURL: url.ShortURL, OriginalURL: url.OriginalURL})
	}

	return models, nil
}

func (r *repo) MarkURLAsDeleted(ctx context.Context, userID string, url string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	urls := r.fetchAllURLs()

	for i, u := range urls {
		if u.ShortURL == url && u.OwnerID == userID {
			urls[i].IsDeleted = true
		}
	}

	if err := os.Truncate(r.path, 0); err != nil {
		return err
	}

	err := r.RewriteURLS(urls, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) fetchAllURLs() []ShortenerURL {
	var shortenerURLs []ShortenerURL

	b, err := os.ReadFile(r.path)
	if err != nil {
		r.log.Error(fmt.Sprintf("can't read urls.txt: %s", err.Error()))
		return shortenerURLs
	}

	for _, line := range bytes.Split(b, []byte("\n")) {
		if len(line) == 0 {
			break
		}
		var url ShortenerURL
		err := json.Unmarshal(line, &url)
		if err != nil {
			r.log.Error(fmt.Sprintf("can't unmarshall shortener url cause: %s", err.Error()))
			continue
		}
		shortenerURLs = append(shortenerURLs, url)
	}

	return shortenerURLs
}

func (r *repo) RewriteURLS(models []ShortenerURL, userID string) error {
	for _, m := range models {
		var id string
		if m.OwnerID == userID {
			id = m.OwnerID
		} else {
			id = userID
		}

		url := ShortenerURL{
			UUID:        uuid.NewString(),
			ShortURL:    m.ShortURL,
			OriginalURL: m.OriginalURL,
			IsDeleted:   m.IsDeleted,
			OwnerID:     id,
		}

		b, err := json.Marshal(&url)
		if err != nil {
			r.log.Error(fmt.Sprintf("can't marshall shortener url cause: %s", err.Error()))
			return err
		}
		b = append(b, '\n')

		_, err = r.writer.Write(b)
		if err != nil {
			r.log.Error(fmt.Sprintf("can't write shortener url cause: %s", err.Error()))
			return err
		}

		_ = r.writer.Flush()
	}

	return nil
}

func (r *repo) Close() error {
	return r.file.Close()
}
