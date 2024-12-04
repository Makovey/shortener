package shortener

import (
	"context"
	"crypto/md5"
	"database/sql/driver"
	"encoding/hex"
	"fmt"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	repo "github.com/Makovey/shortener/internal/service"
	"github.com/Makovey/shortener/internal/service/utils"
)

type service struct {
	repo   repo.Shortener
	pinger driver.Pinger
	cfg    config.Config
	log    logger.Logger
}

func NewShortenerService(
	shortenerRepo repo.Shortener,
	cfg config.Config,
	log logger.Logger,
) def.Shortener {
	return &service{repo: shortenerRepo, cfg: cfg, log: log}
}

func NewChecker(pingerRepo driver.Pinger) def.Checker {
	return &service{pinger: pingerRepo}
}

func (s *service) Shorten(ctx context.Context, url, userID string) (string, error) {
	shortURL := s.generateShortURL(url)
	err := s.repo.SaveUserURL(ctx, shortURL, url, userID)
	fullShortURL := fmt.Sprintf("%s/%s", s.cfg.BaseReturnedURL(), shortURL)
	if err != nil {
		return fullShortURL, err
	}

	return fullShortURL, nil
}

func (s *service) GetFullURL(ctx context.Context, shortURL, userID string) (model.UserFullURL, error) {
	repoModel, err := s.repo.GetFullURL(ctx, shortURL, userID)
	if err != nil {
		return model.UserFullURL{}, err
	}

	return model.UserFullURL{OriginalURL: repoModel.OriginalURL, IsDeleted: repoModel.IsDeleted}, nil
}

func (s *service) ShortBatch(ctx context.Context, batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error) {
	var b []model.ShortenBatch
	for _, req := range batch {
		tmp := model.ShortenBatch{
			CorrelationID: req.CorrelationID,
			ShortURL:      s.generateShortURL(req.OriginalURL),
			OriginalURL:   req.OriginalURL,
		}

		b = append(b, tmp)
	}

	err := s.repo.SaveUserURLs(ctx, b, userID)
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

func (s *service) GetAllURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error) {
	models, err := s.repo.GetUserURLs(ctx, userID)
	if err != nil {
		return nil, err
	}

	for i := range models {
		models[i].ShortURL = fmt.Sprintf("%s/%s", s.cfg.BaseReturnedURL(), models[i].ShortURL)
	}

	return models, nil
}

func (s *service) DeleteUsersURLs(ctx context.Context, userID string, shortURLs []string) []error {
	ch := utils.Generator(ctx, shortURLs)
	results := utils.FanOut(ctx, 5, func(ctx context.Context) chan error {
		errors := make(chan error)

		go func() {
			defer close(errors)

			for url := range ch {
				err := s.repo.MarkURLAsDeleted(ctx, userID, url)
				if err != nil {
					s.log.Warning(fmt.Sprintf("failed to delete users URL %s, error is: %s", url, err.Error()))
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

func (s *service) CheckPing(ctx context.Context) error {
	return s.pinger.Ping(ctx)
}

func (s *service) generateShortURL(url string) string {
	h := md5.New()
	h.Write([]byte(url))

	return hex.EncodeToString(h.Sum(nil)[:7])
}
