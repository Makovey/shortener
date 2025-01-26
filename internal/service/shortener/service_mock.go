package shortener

import (
	"context"

	comModel "github.com/Makovey/shortener/internal/service/model"
	"github.com/Makovey/shortener/internal/transport/model"
)

// MockService мок используемый только для тестов
type MockService struct {
	returnedError error
	returnedModel any
}

// NewMockService конструктор для мока сервиса
// Принимает на вход возвращаемую ошибку или модель, которую нужно вернуть
func NewMockService(returnedError error, returnedModel any) *MockService {
	return &MockService{
		returnedError: returnedError,
		returnedModel: returnedModel,
	}
}

// CreateShortURL мок по созданию урла
func (m *MockService) CreateShortURL(ctx context.Context, url, userID string) (string, error) {
	if m.returnedError != nil {
		return "", m.returnedError
	}

	return "a1b2c3", nil
}

// GetFullURL мок по получанию урла
func (m *MockService) GetFullURL(ctx context.Context, shortURL, userID string) (model.UserFullURL, error) {
	if m.returnedError != nil {
		return model.UserFullURL{}, m.returnedError
	}

	if m.returnedModel != nil {
		if mod, ok := m.returnedModel.(model.UserFullURL); ok {
			return mod, nil
		}
	}

	return model.UserFullURL{OriginalURL: "https://github.com", IsDeleted: false}, nil
}

// ShortBatch мок по созданию пачки коротких урлов
func (m *MockService) ShortBatch(ctx context.Context, batch []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error) {
	if m.returnedError != nil {
		return nil, m.returnedError
	}

	if m.returnedModel != nil {
		if mod, ok := m.returnedModel.([]model.ShortenBatchResponse); ok {
			return mod, nil
		}
	}

	return []model.ShortenBatchResponse{
		{
			CorrelationID: "mockId",
			ShortURL:      "a1b2c3",
		},
	}, nil
}

// GetAllURLs мок по получению вех урлов
func (m *MockService) GetAllURLs(ctx context.Context, userID string) ([]comModel.ShortenBatch, error) {
	if m.returnedError != nil {
		return nil, m.returnedError
	}

	if m.returnedModel != nil {
		if mod, ok := m.returnedModel.([]comModel.ShortenBatch); ok {
			return mod, nil
		}
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

// DeleteUsersURLs мок по пометке урла, как удаленный
func (m *MockService) DeleteUsersURLs(ctx context.Context, userID string, shortURLs []string) []error {
	if m.returnedError != nil {
		return []error{m.returnedError}
	}

	return nil
}
