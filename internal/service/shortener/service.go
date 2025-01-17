package shortener

import (
	"context"
	"crypto/md5"
	"database/sql/driver"
	"encoding/hex"
	"fmt"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
	comModel "github.com/Makovey/shortener/internal/service/model"
	"github.com/Makovey/shortener/internal/service/utils"
	"github.com/Makovey/shortener/internal/transport/model"
)

type Repository interface {
	SaveUserURL(ctx context.Context, shortURL, longURL, userID string) error
	GetFullURL(ctx context.Context, shortURL, userID string) (repoModel.UserURL, error)
	SaveUserURLs(ctx context.Context, models []comModel.ShortenBatch, userID string) error
	GetUserURLs(ctx context.Context, userID string) ([]comModel.ShortenBatch, error)
	MarkURLAsDeleted(ctx context.Context, userID string, url string) error
}

type Service struct {
	repo   Repository
	pinger driver.Pinger
	cfg    config.Config
	log    logger.Logger
}

func NewShortenerService(
	shortenerRepo Repository,
	cfg config.Config,
	log logger.Logger,
) *Service {
	return &Service{repo: shortenerRepo, cfg: cfg, log: log}
}

func NewChecker(pingerRepo driver.Pinger) *Service {
	return &Service{pinger: pingerRepo}
}

func (s *Service) Shorten(ctx context.Context, url, userID string) (string, error) {
	fn := "shortener.Shorten"

	shortURL := s.generateShortURL(url)
	err := s.repo.SaveUserURL(ctx, shortURL, url, userID)
	fullShortURL := fmt.Sprintf("%s/%s", s.cfg.BaseReturnedURL(), shortURL)
	if err != nil {
		return fullShortURL, fmt.Errorf("[%s]: %w", fn, err)
	}

	return fullShortURL, nil
}

func (s *Service) GetFullURL(ctx context.Context, shortURL, userID string) (model.UserFullURL, error) {
	fn := "shortener.GetFullURL"

	userURL, err := s.repo.GetFullURL(ctx, shortURL, userID)
	if err != nil {
		return model.UserFullURL{}, fmt.Errorf("[%s]: %w", fn, err)
	}

	return model.UserFullURL{OriginalURL: userURL.OriginalURL, IsDeleted: userURL.IsDeleted}, nil
}

func (s *Service) ShortBatch(
	ctx context.Context,
	batch []model.ShortenBatchRequest,
	userID string,
) ([]model.ShortenBatchResponse, error) {
	var b []comModel.ShortenBatch
	for _, req := range batch {
		tmp := comModel.ShortenBatch{
			CorrelationID: req.CorrelationID,
			ShortURL:      s.generateShortURL(req.OriginalURL),
			OriginalURL:   req.OriginalURL,
		}

		b = append(b, tmp)
	}

	err := s.repo.SaveUserURLs(ctx, b, userID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", userID, err)
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

func (s *Service) GetAllURLs(ctx context.Context, userID string) ([]comModel.ShortenBatch, error) {
	fn := "shortener.GetAllURLs"

	models, err := s.repo.GetUserURLs(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	for i := range models {
		models[i].ShortURL = fmt.Sprintf("%s/%s", s.cfg.BaseReturnedURL(), models[i].ShortURL)
	}

	return models, nil
}

func (s *Service) DeleteUsersURLs(ctx context.Context, userID string, shortURLs []string) []error {
	fn := "shortener.DeleteUsersURLs"

	ch := utils.Generator(ctx, shortURLs)
	results := utils.FanOut(ctx, 5, func(ctx context.Context) chan error {
		errors := make(chan error)

		go func() {
			defer close(errors)

			for url := range ch {
				err := s.repo.MarkURLAsDeleted(ctx, userID, url)
				if err != nil {
					s.log.Warning(fmt.Sprintf("[%s]: with url %s, %s", fn, url, err.Error()))
				}

				select {
				case <-ctx.Done():
					return
				case errors <- err:
				}
			}
		}()
		return errors
	})

	errorsCh := utils.FanIn(ctx, 3, results...)

	errors := make([]error, 0)
	for e := range errorsCh {
		errors = append(errors, e)
	}

	return errors
}

func (s *Service) CheckPing(ctx context.Context) error {
	return s.pinger.Ping(ctx)
}

func (s *Service) generateShortURL(url string) string {
	h := md5.New()
	h.Write([]byte(url))

	return hex.EncodeToString(h.Sum(nil)[:7])
}
