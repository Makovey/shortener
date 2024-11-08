package shortener

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/Makovey/shortener/internal/config"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/model"
	repo "github.com/Makovey/shortener/internal/service"
)

type service struct {
	repo   repo.Shortener
	pinger repo.Pinger
	cfg    config.Config
}

func (s *service) Short(url string) (string, error) {
	shortURL := s.generateShortURL(url)[:7]
	err := s.repo.Store(shortURL, url)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", s.cfg.BaseReturnedURL(), shortURL), nil
}

func (s *service) Get(shortURL string) (string, error) {
	return s.repo.Get(shortURL)
}

func (s *service) CheckPing() error {
	return s.pinger.Ping()
}

func (s *service) ShortBatch(batch []model.ShortenBatchRequest) ([]model.ShortenBatchResponse, error) {
	var b []model.ShortenBatch
	for _, req := range batch {
		tmp := model.ShortenBatch{
			CorrelationID: req.CorrelationID,
			ShortURL:      s.generateShortURL(req.OriginalURL)[:7],
			OriginalURL:   req.OriginalURL,
		}

		b = append(b, tmp)
	}

	err := s.repo.StoreBatch(b)
	if err != nil {
		return nil, err
	}

	var res []model.ShortenBatchResponse
	for _, req := range b {
		tmp := model.ShortenBatchResponse{
			CorrelationID: req.CorrelationID,
			ShortURL:      fmt.Sprintf("%s/%s", s.cfg.BaseReturnedURL(), req.ShortURL),
		}

		res = append(res, tmp)
	}

	return res, nil
}

func (s *service) generateShortURL(url string) string {
	h := md5.New()
	h.Write([]byte(url))

	return hex.EncodeToString(h.Sum(nil))
}

func NewShortenerService(shortenerRepo repo.Shortener, cfg config.Config) def.Shortener {
	return &service{repo: shortenerRepo, cfg: cfg}
}

func NewChecker(pingerRepo repo.Pinger) def.Checker {
	return &service{pinger: pingerRepo}
}
