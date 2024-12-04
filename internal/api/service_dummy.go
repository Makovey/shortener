package api

import "context"

type dummyChecker struct {
}

func NewDummyChecker() Checker {
	return &dummyChecker{}
}

func (s *dummyChecker) CheckPing(ctx context.Context) error {
	return nil
}
