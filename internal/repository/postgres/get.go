package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Makovey/shortener/internal/repository"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
	"github.com/Makovey/shortener/internal/service/model"
)

// GetFullURL возвращает полный урл по короткому урлу, если он есть в базе данных
func (r *Repo) GetFullURL(ctx context.Context, shortURL, userID string) (*repoModel.UserURL, error) {
	fn := "postgres.GetFullURL"

	row := r.db.QueryRowContext(
		ctx,
		`SELECT original_url, is_deleted FROM shortener WHERE short_url = $1`,
		shortURL,
	)

	var url repoModel.UserURL
	err := row.Scan(&url.OriginalURL, &url.IsDeleted)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("[%s]: %w", fn, repository.ErrURLNotFound)
	}

	return &url, nil
}

// GetUserURLs возвращает все урлы юзера, которые есть в базе данных
func (r *Repo) GetUserURLs(ctx context.Context, userID string) ([]model.ShortenBatch, error) {
	fn := "postgres.GetUserURLs"

	rows, err := r.db.QueryContext(ctx,
		`SELECT short_url, original_url FROM shortener WHERE owner_user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}
	defer rows.Close()

	var models []model.ShortenBatch
	for rows.Next() {
		var shorten model.ShortenBatch
		err = rows.Scan(&shorten.ShortURL, &shorten.OriginalURL)
		if err != nil {
			return nil, fmt.Errorf("[%s]: %w", fn, err)
		}
		models = append(models, shorten)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	return models, nil
}

// GetStats возвращает стистику по сервису, количество пользователей и сокращенных адресов
func (r *Repo) GetStats(ctx context.Context) (model.Stats, error) {
	fn := "postgres.GetStats"

	row := r.db.QueryRowContext(
		ctx,
		`SELECT COUNT(DISTINCT shortener.owner_user_id) AS "users", COUNT(DISTINCT shortener.short_url) as "urls" FROM shortener`,
	)

	var stats model.Stats
	err := row.Scan(&stats.URLS, &stats.Users)
	if err != nil {
		return model.Stats{}, fmt.Errorf("[%s]: %w", fn, err)
	}

	return stats, nil
}
