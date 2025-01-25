package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"path/filepath"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
)

const (
	errUniqueViolatesCode = "23505"
	migrationPath         = "internal/db/migrations"
)

// Repo репозиторий базы данных
type Repo struct {
	log logger.Logger
	db  *sql.DB
}

// NewPostgresRepository конструктор репозитория базы данных
func NewPostgresRepository(cfg config.Config, log logger.Logger) (*Repo, error) {
	fn := "postgresql.NewPostgresRepo"

	db, err := sql.Open("pgx", cfg.DatabaseDSN())
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	path, err := filepath.Abs(migrationPath)
	if err != nil {
		return nil, fmt.Errorf("[%s] could not determine absolute path for migrations: %w", fn, err)
	}

	if err = goose.Up(db, path); err != nil {
		return nil, fmt.Errorf("[%s] could not up migrations: %w", fn, err)
	}

	return &Repo{
		db:  db,
		log: log,
	}, nil
}

// NewPingerRepo на случай если если репозиторий - не Postgres, для поддержки ручки ping
// Если репозиторий Postgres, то метод вызван не будет
func NewPingerRepo(cfg config.Config) (driver.Pinger, error) {
	fn := "postgres.NewPingerRepo"

	db, err := sql.Open("pgx", cfg.DatabaseDSN())
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	return &Repo{
		db: db,
	}, nil
}

// Ping хелсчек для базы данных
func (r *Repo) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := r.db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

// Close закрывает базу данных
func (r *Repo) Close() error {
	return r.db.Close()
}
