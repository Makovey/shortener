package shortener

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/config"
	repo "github.com/Makovey/shortener/internal/service"
)

type service struct {
	repo   repo.Shortener
	pinger repo.Pinger
	cfg    config.Config
}

func (s *service) Short(url, userID string) (string, error) {
	shortURL := s.generateShortURL(url)[:7]
	err := s.repo.Store(shortURL, url, userID)
	fullShortURL := fmt.Sprintf("%s/%s", s.cfg.BaseReturnedURL(), shortURL)
	if err != nil {
		return fullShortURL, err
	}

	return fullShortURL, nil
}

func (s *service) Get(shortURL, userID string) (string, error) {
	return s.repo.Get(shortURL, userID)
}

func (s *service) CheckPing() error {
	return s.pinger.Ping()
}

func (s *service) ShortBatch(batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error) {
	var b []model.ShortenBatch
	for _, req := range batch {
		tmp := model.ShortenBatch{
			CorrelationID: req.CorrelationID,
			ShortURL:      s.generateShortURL(req.OriginalURL)[:7],
			OriginalURL:   req.OriginalURL,
		}

		b = append(b, tmp)
	}

	err := s.repo.StoreBatch(b, userID)
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

func (s *service) GetAll(userID string) ([]model.ShortenBatch, error) {
	models, err := s.repo.GetAll(userID)
	if err != nil {
		return nil, err
	}

	for i := range models {
		models[i].ShortURL = fmt.Sprintf("%s/%s", s.cfg.BaseReturnedURL(), models[i].ShortURL)
	}

	return models, nil
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
