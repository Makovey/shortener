package postgres

import (
	"context"
	"fmt"
)

func (r *Repo) MarkURLAsDeleted(ctx context.Context, userID string, url string) error {
	fn := "postgres.MarkURLAsDeleted"

	_, err := r.db.ExecContext(ctx,
		`UPDATE shortener SET is_deleted = true WHERE owner_user_id = $1 AND short_url = $2`,
		userID,
		url,
	)
	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}

	return nil
}
