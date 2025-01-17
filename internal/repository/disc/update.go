package disc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/Makovey/shortener/internal/repository/model"
)

func (r *repo) MarkURLAsDeleted(ctx context.Context, userID string, url string) error {
	fn := "disc.MarkURLAsDeleted"

	r.mu.Lock()
	defer r.mu.Unlock()

	urls, err := r.fetchAllURLs()
	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}

	for i, u := range urls {
		if u.ShortURL == url && u.OwnerID == userID {
			urls[i].IsDeleted = true
		}
	}

	if err = os.Truncate(r.path, 0); err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}

	err = r.RewriteURLS(urls, userID)
	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}

	return nil
}

func (r *repo) RewriteURLS(models []model.ShortenerURL, userID string) error {
	fn := "disc.RewriteURLS"

	for _, m := range models {
		var id string
		if m.OwnerID == userID {
			id = m.OwnerID
		} else {
			id = userID
		}

		url := model.ShortenerURL{
			UUID:        uuid.NewString(),
			ShortURL:    m.ShortURL,
			OriginalURL: m.OriginalURL,
			IsDeleted:   m.IsDeleted,
			OwnerID:     id,
		}

		b, err := json.Marshal(&url)
		if err != nil {
			return fmt.Errorf("[%s]: %w", fn, err)
		}
		b = append(b, '\n')

		_, err = r.writer.Write(b)
		if err != nil {
			return fmt.Errorf("[%s]: %w", fn, err)
		}

		_ = r.writer.Flush()
	}

	return nil
}
