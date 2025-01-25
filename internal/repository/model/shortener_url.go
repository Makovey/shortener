package model

// ShortenerURL полная информация по урлу
type ShortenerURL struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	IsDeleted   bool   `json:"is_deleted"`
	OwnerID     string `json:"owner_id"`
}
