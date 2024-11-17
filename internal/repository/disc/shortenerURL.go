package disc

type ShortenerURL struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	OwnerID     string `json:"owner_id"`
}
