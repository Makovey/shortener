package mock

import (
	"context"
)

// CheckerMock мок чекера, использовать только для тестов
type CheckerMock struct {
	returnedError error
}

// NewCheckerMock конструктор мока
func NewCheckerMock(
	returnedError error,
) *CheckerMock {
	return &CheckerMock{
		returnedError: returnedError,
	}
}

// CheckPing замоканный метод для тестов.
// Если при создании структуры передали ошибку, то метод ее вернет
func (s *CheckerMock) CheckPing(ctx context.Context) error {
	if s.returnedError != nil {
		return s.returnedError
	}
	return nil
}
