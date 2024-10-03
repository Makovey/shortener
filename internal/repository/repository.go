package repository

type ShortenerRepository interface {
	Store(shortUrl, longUrl string)
	Get(shortUrl string) (string, error)
}
