package dummy

import (
	"context"
)

type DummyChecker struct {
}

func NewDummyChecker() *DummyChecker {
	return &DummyChecker{}
}

func (s *DummyChecker) CheckPing(ctx context.Context) error {
	return nil
}
