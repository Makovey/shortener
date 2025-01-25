package disc

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/Makovey/shortener/internal/logger"
)

// Repo файловый репозитрий
type Repo struct {
	file   *os.File
	path   string
	writer *bufio.Writer
	log    logger.Logger
	mu     sync.RWMutex
}

// NewFileStorage по filePath открывает файл или создает его в случай отсутствия
func NewFileStorage(filePath string, log logger.Logger) (*Repo, error) {
	fn := "disc.NewFileStorage"

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	return &Repo{
		file:   f,
		path:   filePath,
		writer: bufio.NewWriter(f),
		log:    log,
		mu:     sync.RWMutex{},
	}, nil
}

// Close закрывает файл
func (r *Repo) Close() error {
	return r.file.Close()
}
