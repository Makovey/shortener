package inmemory

import "context"

// MarkURLAsDeleted помечает урл как удаленный, при этом из памяти он не удаляется
func (r *Repo) MarkURLAsDeleted(ctx context.Context, userID string, url string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, row := range r.storage {
		if row.shortURL == url && row.userID == userID {
			r.storage[i].isDeleted = true
		}
	}

	return nil
}
