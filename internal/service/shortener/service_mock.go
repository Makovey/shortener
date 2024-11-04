package shortener

import (
	"errors"
	def "github.com/Makovey/shortener/internal/api"
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

func NewMockService(isErrorNeeded bool) def.Shortener {
	return &mockService{
		isErrorNeeded: isErrorNeeded,
	}
}
