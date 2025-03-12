package shortener

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Makovey/shortener/internal/generated/shortener"
	"github.com/Makovey/shortener/internal/transport/grpc"
)

// GetURL - gRPC хендлер по получению урлов
func (s *Server) GetURL(ctx context.Context, req *shortener.GetURLRequest) (*shortener.GetURLResponse, error) {
	fn := "shortener.GetURL"

	userID, err := grpc.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, grpc.ReloginAndTryAgain)
	}

	url, err := s.service.GetFullURL(ctx, req.GetShortUrl(), userID)
	if err != nil {
		s.log.Error(fmt.Sprintf("[%s]: %v", fn, err))
		return nil, status.Error(codes.Internal, grpc.InternalServerError)
	}

	if url.IsDeleted {
		return nil, status.Errorf(codes.NotFound, "short url already deleted: %s", req.GetShortUrl())
	}

	return &shortener.GetURLResponse{LongUrl: url.OriginalURL}, nil
}
