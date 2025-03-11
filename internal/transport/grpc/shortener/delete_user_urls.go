package shortener

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Makovey/shortener/internal/generated/shortener"
	"github.com/Makovey/shortener/internal/transport/grpc"
)

func (s *Server) DeleteUserURLs(ctx context.Context, req *shortener.DeleteUserURLsRequest) (*emptypb.Empty, error) {
	fn := "shortener.DeleteUserURLs"

	userID, err := grpc.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, grpc.ReloginAndTryAgain)
	}

	if len(req.GetShortURLs()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no ids provided")
	}

	errs := s.service.DeleteUsersURLs(ctx, userID, req.GetShortURLs())

	for _, err = range errs {
		if err != nil {
			s.log.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		}
	}

	return &emptypb.Empty{}, nil
}
