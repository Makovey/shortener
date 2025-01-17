package disc

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/service/shortener"
)

type repo struct {
	file   *os.File
	path   string
	writer *bufio.Writer
	log    logger.Logger
	mu     sync.RWMutex
}

func NewFileStorage(filePath string, log logger.Logger) (shortener.Repository, error) {
	fn := "disc.NewFileStorage"

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	return &repo{
		file:   f,
		path:   filePath,
		writer: bufio.NewWriter(f),
		log:    log,
		mu:     sync.RWMutex{},
	}, nil
}

func (r *repo) Close() error {
	return r.file.Close()
}
