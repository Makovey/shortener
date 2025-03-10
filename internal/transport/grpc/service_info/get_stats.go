package service_info

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	proto "github.com/Makovey/shortener/internal/generated/service_info"
	"github.com/Makovey/shortener/internal/transport/grpc"
	"github.com/Makovey/shortener/internal/transport/grpc/model_mapper"
)

// Stats предоставляет статистику по сервису
func (s *InfoServer) Stats(ctx context.Context, empty *emptypb.Empty) (*proto.StatsResponse, error) {
	fn := "service_info.GetStats"

	_, err := grpc.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, grpc.ReloginAndTryAgain)
	}

	model, err := s.service.GetStats(ctx)
	if err != nil {
		s.log.Error(fmt.Sprintf("[%s]: %v", fn, err))
		return nil, status.Error(codes.Internal, grpc.InternalServerError)
	}

	return model_mapper.ToProtoFromStats(&model), nil
}
