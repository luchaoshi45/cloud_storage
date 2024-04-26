package handler

import "net/http"

type Handler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}
