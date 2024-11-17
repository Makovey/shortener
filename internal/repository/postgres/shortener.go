package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
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

func (r *repo) Store(shortURL, longURL, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `
		INSERT INTO shortener (short_url, original_url, owner_user_id) VALUES ($1, $2, $3)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, shortURL, longURL, userID)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return repository.ErrURLIsAlreadyExists
		}
		return err
	}

	r.log.Info(fmt.Sprintf("executed insert, with %s and %s", shortURL, longURL))

	return nil
}

func (r *repo) Get(shortURL, userID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `SELECT original_url FROM shortener WHERE short_url = $1 AND owner_user_id = $2`)
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, shortURL, userID)

	r.log.Info(fmt.Sprintf("queried select, with %s", shortURL))

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

func (r *repo) StoreBatch(models []model.ShortenBatch, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO shortener (short_url, original_url, owner_user_id) VALUES ($1, $2, $3)`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, m := range models {
		_, err := stmt.ExecContext(ctx, m.ShortURL, m.OriginalURL, userID)
		if err != nil {
			tx.Rollback()
			return err
		}
		r.log.Info(fmt.Sprintf("executed insert, with %s and %s", m.ShortURL, m.OriginalURL))
	}

	return tx.Commit()
}

func (r *repo) GetAll(userID string) ([]model.ShortenBatch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, `SELECT short_url, original_url FROM shortener WHERE owner_user_id = $1`)
	if err != nil {
		return []model.ShortenBatch{}, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return []model.ShortenBatch{}, err
	}

	defer rows.Close()

	r.log.Info("queried all URL")

	var models []model.ShortenBatch
	for rows.Next() {
		var shorten model.ShortenBatch
		err := rows.Scan(&shorten.ShortURL, &shorten.OriginalURL)
		if err != nil {
			return nil, err
		}
		models = append(models, shorten)
	}

	return models, nil
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
			created_at TIMESTAMP default CURRENT_TIMESTAMP,
			owner_user_id TEXT NOT NULL,
			UNIQUE (original_url)
		);`

	_, err := r.db.ExecContext(ctx, stmt)

	r.log.Info(fmt.Sprintf("executed %s", stmt))

	if err != nil {
		panic(err)
	}
}
