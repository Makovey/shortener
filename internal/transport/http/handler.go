package http

import (
	"context"

	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/service"
	"github.com/Makovey/shortener/internal/transport"
)

const (
	uuidLength         = 36
	reloginAndTryAgain = "please, relogin again, to get access to this resource"
)

type handler struct {
	service service.Service
	checker Checker
	logger  logger.Logger
}

// NewHTTPHandler конструктор HTTPHandler
func NewHTTPHandler(
	service service.Service,
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
