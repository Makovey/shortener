package service_info

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	proto "github.com/Makovey/shortener/internal/generated/service_info"
)

// Stats предоставляет статистику по сервису
func (s *InfoServer) Stats(ctx context.Context, empty *emptypb.Empty) (*proto.StatsResponse, error) {
	// TODO
	return &proto.StatsResponse{}, nil
}
