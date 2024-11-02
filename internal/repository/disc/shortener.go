package disc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/service"
)

var errURLNotFound = errors.New("url is not existed yet")

type repo struct {
	file   *os.File
	path   string
	writer *bufio.Writer
	log    logger.Logger
}

func (r *repo) Get(shortURL string) (string, error) {
	shortenerURLs := r.fetchAllURLs()
	for _, shortenerURL := range shortenerURLs {
		if shortenerURL.ShortURL == shortURL {
			return shortenerURL.OriginalURL, nil
		}
	}

	return "", errURLNotFound
}

func (r *repo) Store(shortURL, longURL string) error {
	currentURL := ShortenerURL{
		UUID:        uuid.New(),
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
	}
}
