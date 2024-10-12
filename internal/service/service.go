package service

type ShortenerService interface {
	Short(url string) (string, error)
	Get(shortURL string) (string, error)
}
