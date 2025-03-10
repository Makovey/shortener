package interceptor

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/middleware/utils"
)

const jwtMetaName = "jwt"

// JWTAuth проверяет наличие JWT, устанавливает новый, если необходимо
func JWTAuth(
	log logger.Logger,
	jwtUtils utils.JWTUtils,
) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fn := "interceptor.Auth"

		var token string
		var userID string

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		values := md.Get(jwtMetaName)
		if len(values) > 0 {
			token = values[0]
		}

		if len(token) != 0 {
			userID, err = jwtUtils.ParseUserID(token)
			if err != nil && errors.Is(err, utils.ErrParseToken) {
				return nil, status.Error(codes.Internal, "internal server error")
			}

			if userID == "" {
				log.Warning(fmt.Sprintf("[%s]: userID is empty", fn))
				return nil, status.Error(codes.Unauthenticated, "please, relogin and try again")
			}
		}

		if len(token) == 0 || errors.Is(err, utils.ErrTokenExpired) || errors.Is(err, utils.ErrInvalidToken) {
			userID = uuid.NewString()
			tokenString, err := jwtUtils.BuildNewJWT(userID)
			if err != nil {
				return nil, status.Error(codes.Internal, "internal server error")
			}

			metadata.AppendToOutgoingContext(ctx, jwtMetaName, tokenString)
		}

		return handler(ctx, req)
	}
}
