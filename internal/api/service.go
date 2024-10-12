package api

type Shortener interface {
	Short(url string) string
	Get(shortURL string) (string, error)
}