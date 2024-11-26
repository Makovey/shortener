package shortener

import (
	"context"
	"errors"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/model"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
)

type mockService struct {
	isErrorNeeded bool
}

func (m *mockService) Short(url, userID string) (string, error) {
	if m.isErrorNeeded {
		return "", errors.New("mock error")
	}

	return "a1b2c3", nil
}

func (m *mockService) Get(shortURL, userID string) (repoModel.ShortenGet, error) {
	if m.isErrorNeeded {
		return repoModel.ShortenGet{}, errors.New("mock error")
	}

	return repoModel.ShortenGet{OriginalURL: "https://github.com", IsDeleted: false}, nil
}

func (m *mockService) ShortBatch(batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error) {
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

func (m *mockService) GetAll(userID string) ([]model.ShortenBatch, error) {
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

func (m *mockService) DeleteUsersURLS(ctx context.Context, userID string, shortURLs []string) []error {
	if m.isErrorNeeded {
		return []error{errors.New("mock error")}
	}

	return nil
}

func NewMockService(isErrorNeeded bool) def.Shortener {
	return &mockService{
		isErrorNeeded: isErrorNeeded,
	}
}
