package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Makovey/shortener/internal/api/model"
	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger"
	"github.com/Makovey/shortener/internal/repository"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
	"github.com/Makovey/shortener/internal/service"
)

const (
	errUniqueViolatesCode = "23505"
)

type repo struct {
	log logger.Logger
	db  *sql.DB
}

func NewPostgresRepository(cfg config.Config, log logger.Logger) service.Shortener {
	db, err := sql.Open("pgx", cfg.DatabaseDSN())
	if err != nil {
		panic(err)
	}

	r := &repo{
		db:  db,
		log: log,
	}
	r.prepareDB()

	return r
}

// NewPingerRepo нужен, на случай если если репозиторий - не Postgres, для поддержки ручки ping
// Если репозиторий Postgres, то метод вызван не будет
func NewPingerRepo(cfg config.Config) driver.Pinger {
	db, err := sql.Open("pgx", cfg.DatabaseDSN())
	if err != nil {
		panic(err)
	}

	return &repo{
		db: db,
	}
}

func (r *repo) SaveUserURL(ctx context.Context, shortURL, longURL, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx,
		`INSERT INTO shortener (short_url, original_url, owner_user_id) VALUES ($1, $2, $3)`,
		shortURL, longURL, userID,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok && pgErr.Code == errUniqueViolatesCode {
			return repository.ErrURLIsAlreadyExists
		}
		return err
	}

	r.log.Info(fmt.Sprintf("executed insert, with %s and %s", shortURL, longURL))

	return nil
}

func (r *repo) GetFullURL(ctx context.Context, shortURL, userID string) (repoModel.UserURL, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := r.db.QueryRowContext(ctx, `SELECT original_url, is_deleted FROM shortener WHERE short_url = $1`, shortURL)

	r.log.Info(fmt.Sprintf("queried select, with %s", shortURL))

	var getModel repoModel.UserURL
	err := row.Scan(&getModel.OriginalURL, &getModel.IsDeleted)
	if errors.Is(err, sql.ErrNoRows) {
		return repoModel.UserURL{}, repository.ErrURLNotFound
	}

	return getModel, nil
}

func (r *repo) SaveUserURLs(ctx context.Context, models []model.ShortenBatch, userID string) error {
	if len(models) == 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx,
		`INSERT INTO shortener (short_url, original_url, owner_user_id) 
			    VALUES`+generateValuesPlaceholder(len(models), 3)+
			`ON CONFLICT (original_url) DO UPDATE SET original_url = excluded.original_url`,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	args := make([]any, 0, cap(models)*3)
	for _, m := range models {
		args = append(args, m.ShortURL, m.OriginalURL, userID)
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	r.log.Info(fmt.Sprintf("executed batch insert for, %s", userID))

	return tx.Commit()
}

func (r *repo) GetUserURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error) {
	ctx, cancel := context.WithTimeout(ctx, 7*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx,
		`SELECT short_url, original_url FROM shortener WHERE owner_user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r.log.Info("queried all URL")

	var models []model.ShortenBatch
	for rows.Next() {
		var shorten model.ShortenBatch
		err = rows.Scan(&shorten.ShortURL, &shorten.OriginalURL)
		if err != nil {
			return nil, err
		}
		models = append(models, shorten)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return models, nil
}

func (r *repo) MarkURLAsDeleted(ctx context.Context, userID string, url string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx,
		`UPDATE shortener SET is_deleted = true WHERE owner_user_id = $1 AND short_url = $2`,
		userID, url,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := r.db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func (r *repo) Close() error {
	return r.db.Close()
}

func (r *repo) prepareDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt := `
		CREATE TABLE IF NOT EXISTS shortener (
			id SERIAL PRIMARY KEY,
			short_url TEXT NOT NULL,
			original_url TEXT NOT NULL,
			created_at TIMESTAMP default CURRENT_TIMESTAMP,
			owner_user_id TEXT NOT NULL,
			is_deleted BOOLEAN default FALSE,
			UNIQUE (original_url)
		);`

	_, err := r.db.ExecContext(ctx, stmt)

	r.log.Info(fmt.Sprintf("executed %s", stmt))

	if err != nil {
		panic(err)
	}
}

func generateValuesPlaceholder(count int, columns int) string {
	if count == 0 {
		return ""
	}

	var result strings.Builder
	for i := 0; i < count*columns; i++ {
		if i%columns == columns-1 {
			result.WriteString(fmt.Sprintf("$%d),", i+1))
		} else if i%columns == 0 {
			result.WriteString(fmt.Sprintf("($%d,", i+1))
		} else {
			result.WriteString(fmt.Sprintf("$%d,", i+1))
		}
	}

	buf := result.String()
	return strings.Trim(buf, ",")
}
