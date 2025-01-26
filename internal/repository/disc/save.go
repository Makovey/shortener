package disc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	repoModel "github.com/Makovey/shortener/internal/repository/model"
	"github.com/Makovey/shortener/internal/service/model"
)

// SaveUserURL сохраняет полную информацию по урлу с UserID
func (r *Repo) SaveUserURL(ctx context.Context, shortURL, longURL, userID string) error {
	fn := "disc.SaveUserURL"

	r.mu.Lock()
	defer r.mu.Unlock()

	currentURL := repoModel.ShortenerURL{
		UUID:        uuid.New().String(),
		ShortURL:    shortURL,
		OriginalURL: longURL,
		OwnerID:     userID,
		IsDeleted:   false,
	}

	b, err := json.Marshal(&currentURL)
	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}
	b = append(b, '\n')

	_, err = r.writer.Write(b)
	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}

	_ = r.writer.Flush()

	return nil
}

// SaveUserURLs сохраняет полную информацию по урлам с UserID
func (r *Repo) SaveUserURLs(ctx context.Context, models []model.ShortenBatch, userID string) error {
	fn := "disc.SaveUserURLs"

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, m := range models {
		url := repoModel.ShortenerURL{
			UUID:        m.CorrelationID,
			ShortURL:    m.ShortURL,
			OriginalURL: m.OriginalURL,
			OwnerID:     userID,
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
