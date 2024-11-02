package api

type dummyChecker struct {
}

func (s *dummyChecker) CheckPing() error {
	return nil
}

func NewDummyChecker() Checker {
	return &dummyChecker{}
}
