package service

type Shortener interface {
	Store(shortURL, longURL string)
	Get(shortURL string) (string, error)
}
