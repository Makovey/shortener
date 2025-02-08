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

// Repository основной интерфейс для репозитория, отвечает за хранение данных
//
//go:generate mockgen -source=service.go -destination=../../repository/mocks/repository_mock.go -package=mocks
type Repository interface {
	SaveUserURL(ctx context.Context, shortURL, longURL, userID string) error
	GetFullURL(ctx context.Context, shortURL, userID string) (*repoModel.UserURL, error)
	SaveUserURLs(ctx context.Context, models []comModel.ShortenBatch, userID string) error
	GetUserURLs(ctx context.Context, userID string) ([]comModel.ShortenBatch, error)
	MarkURLAsDeleted(ctx context.Context, userID string, url string) error
}

// Service он же useCase, слой отвечающий за бизнес-логику приложения
type Service struct {
	repo   Repository
	pinger driver.Pinger
	cfg    config.Config
	log    logger.Logger
}

// NewShortenerService конструктор сервиса
func NewShortenerService(
	shortenerRepo Repository,
	cfg config.Config,
	log logger.Logger,
) *Service {
	return &Service{repo: shortenerRepo, cfg: cfg, log: log}
}

// NewChecker конструктор чекера
func NewChecker(pingerRepo driver.Pinger) *Service {
	return &Service{pinger: pingerRepo}
}

// CreateShortURL метод по созданию короткого урла
// Возвращает полный адрес с коротким урлом
func (s *Service) CreateShortURL(ctx context.Context, url, userID string) (string, error) {
	fn := "shortener.CreateShortURL"

	shortURL := generateShortURL(url)
	err := s.repo.SaveUserURL(ctx, shortURL, url, userID)
	fullShortURL := fmt.Sprintf("%s/%s", s.cfg.BaseReturnedURL(), shortURL)
	if err != nil {
		return fullShortURL, fmt.Errorf("[%s]: %w", fn, err)
	}

	return fullShortURL, nil
}

// GetFullURL метод по получению урла, на вход принимает короткую версию урла
func (s *Service) GetFullURL(ctx context.Context, shortURL, userID string) (model.UserFullURL, error) {
	fn := "shortener.GetFullURL"

	userURL, err := s.repo.GetFullURL(ctx, shortURL, userID)
	if err != nil {
		return model.UserFullURL{}, fmt.Errorf("[%s]: %w", fn, err)
	}

	return model.UserFullURL{OriginalURL: userURL.OriginalURL, IsDeleted: userURL.IsDeleted}, nil
}

// ShortBatch метод по созданию списка коротких урлов
// Принимает на вход список полных урлов
func (s *Service) ShortBatch(
	ctx context.Context,
	batch []model.ShortenBatchRequest,
	userID string,
) ([]model.ShortenBatchResponse, error) {
	b := fromTransportToRepoShortenBatch(batch)

	err := s.repo.SaveUserURLs(ctx, b, userID)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", userID, err)
	}

	return fromRepoToShortenBatchResponse(b, s.cfg.BaseReturnedURL()), nil
}

// GetAllURLs метод по получению всех урлов юзера по UserID
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

// DeleteUsersURLs помечает урлы юзера как удаленные
// Принимает на вход список коротких урлов
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

// CheckPing метод по пингу репозитрия
func (s *Service) CheckPing(ctx context.Context) error {
	return s.pinger.Ping(ctx)
}

func generateShortURL(url string) string {
	h := md5.New()
	h.Write([]byte(url))

	return hex.EncodeToString(h.Sum(nil)[:7])
}
