package interceptor

import (
	"context"
	"net"
	"slices"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const realIPMetaName = "X-Real-IP"

// CheckSubnet проверяет досутпен ли ресурс для ip клиента
func CheckSubnet(subnet string, acceptMethods []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		method := strings.Split(info.FullMethod, "/")[len(strings.Split(info.FullMethod, "/"))-1]
		if slices.Contains(acceptMethods, method) {
			var ip string

			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
			}

			values := md.Get(realIPMetaName)
			if len(values) > 0 {
				ip = values[0]
			}

			_, trustedIP, err := net.ParseCIDR(subnet)
			if err != nil {
				return nil, status.Error(codes.FailedPrecondition, err.Error())
			}

			if trustedIP.Contains(net.ParseIP(ip)) {
				return nil, status.Error(codes.PermissionDenied, "your ip is not trusted")
			}
		}

		return handler(ctx, req)
	}
}
