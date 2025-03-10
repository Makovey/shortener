package service_info

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Ping grpc хендлер /ping
func (s *InfoServer) Ping(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	fn := "service_info.Ping"

	err := s.checker.CheckPing(ctx)
	if err != nil {
		s.log.Error(fmt.Sprintf("[%s]: %s", fn, err.Error()))
		return &emptypb.Empty{}, status.Error(codes.Internal, "internal server error")
	}

	return &emptypb.Empty{}, nil
}
