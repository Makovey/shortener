package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/repository"
	"github.com/Makovey/shortener/internal/service"
)

type repo struct {
	log logger.Logger
	db  *sql.DB
}

func (r repo) Store(shortURL, longURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `INSERT INTO shortener (short_url, original_url) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, stmt, shortURL, longURL)

	r.log.Info(fmt.Sprintf("Executed %s, with %s and %s", stmt, shortURL, longURL))

	if err != nil {
		return err
	}

	return nil
}

func (r repo) Get(shortURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `SELECT original_url FROM shortener WHERE short_url = $1`
	row := r.db.QueryRowContext(ctx, stmt, shortURL)

	r.log.Info(fmt.Sprintf("Queried %s, with %s", stmt, shortURL))

	var originalURL string
	err := row.Scan(&originalURL)
	if errors.Is(err, sql.ErrNoRows) {
		return "", repository.ErrURLNotFound
	}

	return originalURL, nil
}

func (r repo) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func NewPostgresRepository(cfg config.Config, log logger.Logger) service.Shortener {
	db, err := sql.Open("pgx", cfg.DatabaseDSN())
	if err != nil {
		panic(err)
	}

	r := repo{db: db, log: log}
	r.prepareDB()

	return r
}

// NewPingerRepo нужен, на случай если если репозиторий - не Postgres, для поддержки ручки ping
// Если репозиторий Postgres, то метод вызван не будет
func NewPingerRepo(cfg config.Config) service.Pinger {
	db, err := sql.Open("pgx", cfg.DatabaseDSN())
	if err != nil {
		panic(err)
	}

	return &repo{db: db}
}

func (r repo) prepareDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt := `
		CREATE TABLE IF NOT EXISTS shortener (
			id SERIAL PRIMARY KEY,
			short_url TEXT,
			original_url TEXT
		);`

	_, err := r.db.ExecContext(ctx, stmt)

	r.log.Info(fmt.Sprintf("Executed %s", stmt))

	if err != nil {
		panic(err)
	}
}

func (r repo) Close() error {
	return r.db.Close()
}
