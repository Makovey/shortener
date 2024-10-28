package service

type Shortener interface {
	Store(shortURL, longURL string) error
	Get(shortURL string) (string, error)
}
