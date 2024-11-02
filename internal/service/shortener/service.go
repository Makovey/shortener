package shortener

import (
	"crypto/md5"
	"encoding/hex"

	def "github.com/Makovey/shortener/internal/api"
	repo "github.com/Makovey/shortener/internal/service"
)

type service struct {
	repo   repo.Shortener
	pinger repo.Pinger
}

func (s *service) Short(url string) string {
	shortURL := s.generateShortURL(url)[:7]
	s.repo.Store(shortURL, url)

	return shortURL
}

func (s *service) Get(shortURL string) (string, error) {
	return s.repo.Get(shortURL)
}

func (s *service) CheckPing() error {
	return s.pinger.Ping()
}

func (s *service) generateShortURL(url string) string {
	h := md5.New()
	h.Write([]byte(url))

	return hex.EncodeToString(h.Sum(nil))
}

func NewShortenerService(shortenerRepo repo.Shortener) def.Shortener {
	return &service{repo: shortenerRepo}
}

func NewChecker(pingerRepo repo.Pinger) def.Checker {
	return &service{pinger: pingerRepo}
}
