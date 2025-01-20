package http

import (
	"context"

	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/middleware"
	comModel "github.com/Makovey/shortener/internal/service/model"
	"github.com/Makovey/shortener/internal/transport"
	"github.com/Makovey/shortener/internal/transport/model"
)

const (
	uuidLength         = 36
	reloginAndTryAgain = "please, relogin again, to get access to this resource"
)

type Service interface {
	CreateShortURL(ctx context.Context, url, userID string) (string, error)
	GetFullURL(ctx context.Context, shortURL, userID string) (model.UserFullURL, error)
	ShortBatch(ctx context.Context, batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error)
	GetAllURLs(ctx context.Context, userID string) ([]comModel.ShortenBatch, error)
	DeleteUsersURLs(ctx context.Context, userID string, shortURLs []string) []error
}

type handler struct {
	service Service
	checker Checker
	logger  logger.Logger
}

func NewHTTPHandler(
	service Service,
	logger logger.Logger,
	checker Checker,
) transport.HTTPHandler {
	return &handler{
		service: service,
		logger:  logger,
		checker: checker,
	}
}

func getUserIDFromContext(ctx context.Context) string {
	if ctx.Value(middleware.CtxUserIDKey) == nil {
		return ""
	}

	return ctx.Value(middleware.CtxUserIDKey).(string)
}
