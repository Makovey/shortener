package grpc

import (
	"context"

	"github.com/Makovey/shortener/internal/interceptor"
)

// Тексты ошибок
const (
	uuidLength          = 36
	ReloginAndTryAgain  = "please, relogin again, to get access to this resource"
	InternalServerError = "internal server error"
)

// GetUserIDFromContext - хелпер, вытаскивает id из контекста
func GetUserIDFromContext(ctx context.Context) (string, error) {
	if ctx.Value(interceptor.CtxUserIDKey) == nil {
		return "", nil
	}

	userID := ctx.Value(interceptor.CtxUserIDKey).(string)
	if userID == "" || len(userID) != uuidLength {
		return "", nil
	}

	return userID, nil
}
