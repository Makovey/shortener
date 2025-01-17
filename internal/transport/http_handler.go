package transport

import "net/http"

type HTTPHandler interface {
	PostNewURL(w http.ResponseWriter, r *http.Request)
	GetURL(w http.ResponseWriter, r *http.Request)
	PostShortenURL(w http.ResponseWriter, r *http.Request)
	GetPing(w http.ResponseWriter, r *http.Request)
	PostBatch(w http.ResponseWriter, r *http.Request)
	GetAllURLS(w http.ResponseWriter, r *http.Request)
	DeleteURLS(w http.ResponseWriter, r *http.Request)
}
