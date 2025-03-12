package shortener

import (
	proto "github.com/Makovey/shortener/internal/generated/shortener"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/service"
)

// Server сервер основных методов по сокращению урлов
type Server struct {
	proto.UnimplementedShortenerServer

	log     logger.Logger
	service service.Service
}

// NewServer констурктор Server
func NewServer(
	log logger.Logger,
	service service.Service,
) *Server {
	return &Server{
		log:     log,
		service: service,
	}
}
