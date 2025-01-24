package mock

import (
	"context"
)

type CheckerMock struct {
	returnedError error
}

func NewCheckerMock(
	returnedError error,
) *CheckerMock {
	return &CheckerMock{
		returnedError: returnedError,
	}
}

func (s *CheckerMock) CheckPing(ctx context.Context) error {
	if s.returnedError != nil {
		return s.returnedError
	}
	return nil
}
