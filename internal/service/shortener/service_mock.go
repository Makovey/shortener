package shortener

import (
	"context"
	"errors"

	comModel "github.com/Makovey/shortener/internal/service/model"
	"github.com/Makovey/shortener/internal/transport/model"
)

type MockService struct {
	isErrorNeeded bool
}

func NewMockService(isErrorNeeded bool) *MockService {
	return &MockService{
		isErrorNeeded: isErrorNeeded,
	}
}

func (m *MockService) Shorten(ctx context.Context, url, userID string) (string, error) {
	if m.isErrorNeeded {
		return "", errors.New("mock error")
	}

	return "a1b2c3", nil
}

func (m *MockService) GetFullURL(ctx context.Context, shortURL, userID string) (model.UserFullURL, error) {
	if m.isErrorNeeded {
		return model.UserFullURL{}, errors.New("mock error")
	}

	return model.UserFullURL{OriginalURL: "https://github.com", IsDeleted: false}, nil
}

func (m *MockService) ShortBatch(ctx context.Context, batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error) {
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

func (m *MockService) GetAllURLs(ctx context.Context, userID string) ([]comModel.ShortenBatch, error) {
	if m.isErrorNeeded {
		return nil, errors.New("mock error")
	}

	return []comModel.ShortenBatch{
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

func (m *MockService) DeleteUsersURLs(ctx context.Context, userID string, shortURLs []string) []error {
	if m.isErrorNeeded {
		return []error{errors.New("mock error")}
	}

	return nil
}
