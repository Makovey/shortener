package shortener

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Makovey/shortener/internal/generated/shortener"
	"github.com/Makovey/shortener/internal/transport/grpc"
	"github.com/Makovey/shortener/internal/transport/grpc/mapper"
)

func (s *Server) GetUserURLs(ctx context.Context, req *emptypb.Empty) (*shortener.GetUserURLsResponse, error) {
	fn := "shortener.GetUserURLs"

	userID, err := grpc.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, grpc.ReloginAndTryAgain)
	}

	models, err := s.service.GetAllURLs(ctx, userID)
	if err != nil {
		s.log.Error(fmt.Sprintf("[%s]: %v", fn, err))
		return nil, status.Error(codes.Internal, grpc.InternalServerError)
	}

	return mapper.FromBatchToGetUserURLsResponse(models), nil
}
