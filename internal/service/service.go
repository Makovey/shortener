package service

type ShortenerService interface {
	Short(url string) (string, error)
	Get(shortUrl string) (string, error)
}
