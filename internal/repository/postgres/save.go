package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Makovey/shortener/internal/repository"
	"github.com/Makovey/shortener/internal/service/model"
)

// SaveUserURL сохраняет полную информацию по урлу с UserID
func (r *Repo) SaveUserURL(ctx context.Context, shortURL, longURL, userID string) error {
	fn := "postgres.SaveUserURL"

	_, err := r.db.ExecContext(ctx,
		`INSERT INTO shortener (short_url, original_url, owner_user_id) VALUES ($1, $2, $3)`,
		shortURL,
		longURL,
		userID,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok && pgErr.Code == errUniqueViolatesCode {
			return fmt.Errorf("[%s]: %w", fn, repository.ErrURLIsAlreadyExists)
		}
		return fmt.Errorf("[%s]: %w", fn, err)
	}

	return nil
}

// SaveUserURLs сохраняет полную информацию по урлам с UserID
func (r *Repo) SaveUserURLs(ctx context.Context, models []model.ShortenBatch, userID string) error {
	fn := "postgres.SaveUserURLs"

	if len(models) == 0 {
		return nil
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(
		ctx,
		`INSERT INTO shortener (short_url, original_url, owner_user_id)
		VALUES`+generateValuesPlaceholder(len(models), 3)+
			`ON CONFLICT (original_url) DO UPDATE SET original_url = excluded.original_url`,
	)
	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}
	defer stmt.Close()

	args := make([]any, 0, cap(models)*3)
	for _, m := range models {
		args = append(args, m.ShortURL, m.OriginalURL, userID)
	}

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}

	return tx.Commit()
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
