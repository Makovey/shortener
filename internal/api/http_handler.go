package api

import "net/http"

type HTTPHandler interface {
	PostNewURLHandler(w http.ResponseWriter, r *http.Request)
	GetURLHandler(w http.ResponseWriter, r *http.Request)
	PostShortenURLHandler(w http.ResponseWriter, r *http.Request)
	GetPing(w http.ResponseWriter, r *http.Request)
}
