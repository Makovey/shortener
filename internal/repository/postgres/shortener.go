package postgres

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/service"
)

type repo struct {
	db *sql.DB
}

func (r repo) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func NewPingerRepo(cfg config.Config) service.Pinger {
	db, err := sql.Open("pgx", cfg.DatabaseDSN())
	if err != nil {
		panic(err)
	}

	return &repo{db: db}
}

func (r repo) Close() error {
	return r.db.Close()
}
