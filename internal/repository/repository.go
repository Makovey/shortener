package repository

type ShortenerRepository interface {
	Store(shortURL, longURL string)
	Get(shortURL string) (string, error)
}
