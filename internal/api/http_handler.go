package api

import "net/http"

type HttpHandler interface {
	PostNewUrlHandler(w http.ResponseWriter, r *http.Request)
	GetUrlHandler(w http.ResponseWriter, r *http.Request)
}
