package shortener

import (
	"context"
	"errors"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/model"
)

type mockService struct {
	isErrorNeeded bool
}

func NewMockService(isErrorNeeded bool) def.Shortener {
	return &mockService{
		isErrorNeeded: isErrorNeeded,
	}
}

func (m *mockService) Shorten(ctx context.Context, url, userID string) (string, error) {
	if m.isErrorNeeded {
		return "", errors.New("mock error")
	}

	return "a1b2c3", nil
}

func (m *mockService) GetFullURL(ctx context.Context, shortURL, userID string) (model.UserFullURL, error) {
	if m.isErrorNeeded {
		return model.UserFullURL{}, errors.New("mock error")
	}

	return model.UserFullURL{OriginalURL: "https://github.com", IsDeleted: false}, nil
}

func (m *mockService) ShortBatch(ctx context.Context, batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error) {
	if m.isErrorNeeded {
		return nil, errors.New("mock error")
	}

	return []model.ShortenBatchResponse{
		{
			CorrelationID: "mockId",
			ShortURL:      "a1b2c3",
		},
	}, nil
}

func (m *mockService) GetAllURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error) {
	if m.isErrorNeeded {
		return nil, errors.New("mock error")
	}

	return []model.ShortenBatch{
		{
			ShortURL:    "a1b2c3",
			OriginalURL: "https://github.com",
		},
		{
			ShortURL:    "d4e5f6",
			OriginalURL: "https://gitlab.com",
		},
	}, nil
}

func (m *mockService) DeleteUsersURLs(ctx context.Context, userID string, shortURLs []string) []error {
	if m.isErrorNeeded {
		return []error{errors.New("mock error")}
	}

	return nil
}
