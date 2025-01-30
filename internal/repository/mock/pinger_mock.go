package mock

import (
	"context"

	"database/sql/driver"
)

type pingerMock struct {
	pingErr error
}

// NewPingerMock конструктор для пинга, только для тестов
func NewPingerMock(pingErr error) driver.Pinger {
	return &pingerMock{pingErr: pingErr}
}

// Ping стаба для пинга
func (p *pingerMock) Ping(ctx context.Context) error {
	if p.pingErr != nil {
		return p.pingErr
	}

	return nil
}
