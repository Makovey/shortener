package model

// ShortenRequest клиентский запрос
type ShortenRequest struct {
	URL string `json:"url"`
}

// ShortenResponse ответ клиенту
type ShortenResponse struct {
	Result string `json:"result"`
}
