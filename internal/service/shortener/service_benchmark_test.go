package shortener

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/middleware"
	"github.com/Makovey/shortener/internal/repository/inmemory"
)

func BenchmarkService_CreateShortURL(b *testing.B) {
	log := stdout.NewLoggerDummy()
	repo := inmemory.NewRepositoryInMemory()
	cfg := config.NewConfig(log)
	service := NewShortenerService(repo, cfg, log)

	userID := uuid.NewString()
	ctx := context.WithValue(context.Background(), middleware.CtxUserIDKey, userID)

	for i := 0; i < b.N; i++ {
		_, err := service.CreateShortURL(ctx, "https://github.com", userID)
		if err != nil {
			b.Fatal(err)
		}
	}
}
