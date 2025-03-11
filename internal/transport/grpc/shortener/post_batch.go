package shortener

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Makovey/shortener/internal/generated/shortener"
	"github.com/Makovey/shortener/internal/transport/grpc"
	"github.com/Makovey/shortener/internal/transport/grpc/mapper"
)

// PostBatchURL - gRPC хендлер по вставке новых урлов
func (s *Server) PostBatchURL(ctx context.Context, req *shortener.PostBatchURLRequest) (*shortener.PostBatchURLResponse, error) {
	fn := "shortener.PostBatchURL"

	userID, err := grpc.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, grpc.ReloginAndTryAgain)
	}

	models, err := s.service.ShortBatch(ctx, mapper.FromBatchProtoToBatchRequest(req), userID)
	if err != nil {
		s.log.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		return nil, status.Error(codes.Internal, grpc.InternalServerError)
	}

	if len(models) == 0 {
		return nil, status.Error(codes.InvalidArgument, "request body is empty")
	}

	return mapper.FromBatchRequestToPostBatchProto(&models), nil
}
