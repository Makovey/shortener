package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/repository"
	"github.com/Makovey/shortener/internal/service"
)

type repo struct {
	log logger.Logger
	db  *sql.DB
	mu  sync.RWMutex
}

func (r *repo) Store(shortURL, longURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `INSERT INTO shortener (short_url, original_url) VALUES ($1, $2)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, shortURL, longURL)

	r.log.Info(fmt.Sprintf("Store: Executed insert, with %s and %s", shortURL, longURL))

	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Get(shortURL string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `SELECT original_url FROM shortener WHERE short_url = $1`)
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, shortURL)

	r.log.Info(fmt.Sprintf("Queried select, with %s", shortURL))

	var originalURL string
	err = row.Scan(&originalURL)
	if errors.Is(err, sql.ErrNoRows) {
		return "", repository.ErrURLNotFound
	}

	return originalURL, nil
}

func (r *repo) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := r.db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func (r *repo) StoreBatch(models []model.ShortenBatch) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO shortener (short_url, original_url) VALUES ($1, $2)`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, m := range models {
		_, err := stmt.ExecContext(ctx, m.ShortURL, m.OriginalURL)
		if err != nil {
			tx.Rollback()
			return err
		}
		r.log.Info(fmt.Sprintf("Batch: Executed insert, with %s and %s", m.ShortURL, m.OriginalURL))
	}

	return tx.Commit()
}

func (r *repo) Close() error {
	return r.db.Close()
}

func NewPostgresRepository(cfg config.Config, log logger.Logger) service.Shortener {
	db, err := sql.Open("pgx", cfg.DatabaseDSN())
	if err != nil {
		panic(err)
	}

	r := &repo{
		db:  db,
		log: log,
		mu:  sync.RWMutex{},
	}
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

	return &repo{
		db: db,
		mu: sync.RWMutex{},
	}
}

func (r *repo) prepareDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt := `
		CREATE TABLE IF NOT EXISTS shortener (
			id SERIAL PRIMARY KEY,
			short_url TEXT,
			original_url TEXT,
			created_at TIMESTAMP default CURRENT_TIMESTAMP
		);`

	_, err := r.db.ExecContext(ctx, stmt)

	r.log.Info(fmt.Sprintf("Executed %s", stmt))

	if err != nil {
		panic(err)
	}
}
