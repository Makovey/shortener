package shortener

import (
	"crypto/md5"
	"encoding/hex"

	repo "github.com/Makovey/shortener/internal/repository"
	def "github.com/Makovey/shortener/internal/service"
)

type service struct {
	repo repo.ShortenerRepository
}

func (s *service) Short(url string) (string, error) {
	shortURL := s.generateShortURL(url)[:7]
	s.repo.Store(shortURL, url)

	return shortURL, nil
}

func (s *service) Get(shortURL string) (string, error) {
	return s.repo.Get(shortURL)
}

func (s *service) generateShortURL(url string) string {
	h := md5.New()
	h.Write([]byte(url))

	return hex.EncodeToString(h.Sum(nil))
}

func NewShortenerService(shortenerRepo repo.ShortenerRepository) def.ShortenerService {
	return &service{repo: shortenerRepo}
}
