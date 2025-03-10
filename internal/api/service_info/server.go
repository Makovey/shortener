package service_info

import (
	proto "github.com/Makovey/shortener/internal/generated/service_info"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/transport/http"
)

// InfoServer сервер для утилитарных хендлеров
type InfoServer struct {
	proto.UnimplementedServiceInfoServer

	log     logger.Logger
	checker http.Checker
}

// NewInfoServer констурктор InfoServer
func NewInfoServer(
	log logger.Logger,
	checker http.Checker,
) *InfoServer {
	return &InfoServer{
		log:     log,
		checker: checker,
	}
}
