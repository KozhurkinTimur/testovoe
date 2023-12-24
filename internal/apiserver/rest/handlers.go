package rest

import "net/http"

type Handler interface {
	HandlerRegistration(w http.ResponseWriter, r *http.Request)
}