package shortener

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Makovey/shortener/internal/generated/shortener"
	"github.com/Makovey/shortener/internal/transport/grpc"
)

// PostURL - gRPC хендлер по вставке нового урла
func (s *Server) PostURL(ctx context.Context, req *shortener.PostURLRequest) (*shortener.PostURLResponse, error) {
	fn := "shortener.PostURL"

	userID, err := grpc.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, grpc.ReloginAndTryAgain)
	}

	url, err := s.service.CreateShortURL(ctx, req.GetLongUrl(), userID)
	if err != nil {
		s.log.Error(fmt.Sprintf("[%s]: %v", fn, err))
		return nil, status.Error(codes.Internal, grpc.InternalServerError)
	}

	return &shortener.PostURLResponse{FullShortUrl: url}, nil
}
