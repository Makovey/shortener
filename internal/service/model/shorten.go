package model

// ShortenBatch модель, которую передают батчем
type ShortenBatch struct {
	CorrelationID string `json:"correlation_id,omitempty"`
	OriginalURL   string `json:"original_url"`
	ShortURL      string `json:"short_url"`
}
