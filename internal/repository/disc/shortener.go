package disc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"

	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/repository"
	"github.com/Makovey/shortener/internal/service"
)

type repo struct {
	file   *os.File
	path   string
	writer *bufio.Writer
	log    logger.Logger
	mu     sync.RWMutex
}

func (r *repo) Get(shortURL string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	shortenerURLs := r.fetchAllURLs()
	for _, shortenerURL := range shortenerURLs {
		if shortenerURL.ShortURL == shortURL {
			return shortenerURL.OriginalURL, nil
		}
	}

	return "", repository.ErrURLNotFound
}

func (r *repo) Store(shortURL, longURL string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	currentURL := ShortenerURL{
		UUID:        uuid.New().String(),
		ShortURL:    shortURL,
		OriginalURL: longURL,
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

func (r *repo) StoreBatch(models []model.ShortenBatch) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, m := range models {
		url := ShortenerURL{
			UUID:        m.CorrelationID,
			ShortURL:    m.ShortURL,
			OriginalURL: m.OriginalURL,
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
