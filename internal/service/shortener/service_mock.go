package shortener

import (
	"errors"

	def "github.com/Makovey/shortener/internal/api"
	"github.com/Makovey/shortener/internal/api/model"
)

type mockService struct {
	isErrorNeeded bool
}

func (m *mockService) Short(url string) (string, error) {
	if m.isErrorNeeded {
		return "", errors.New("mock error")
	}

	return "a1b2c3", nil
}

func (m *mockService) Get(shortURL string) (string, error) {
	if m.isErrorNeeded {
		return "", errors.New("mock error")
	}

	return "https://github.com", nil
}

func (m *mockService) ShortBatch(batch []model.ShortenBatchRequest) ([]model.ShortenBatchResponse, error) {
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

func NewMockService(isErrorNeeded bool) def.Shortener {
	return &mockService{
		isErrorNeeded: isErrorNeeded,
	}
}
