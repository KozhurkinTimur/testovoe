package handlersgoogle

import (
	"net/http"
)

type GoogleHandler struct {
}

func New() *GoogleHandler {
	return &GoogleHandler{}
}

func (h *GoogleHandler) HandlerRegistration(w http.ResponseWriter, r *http.Request) {}

func (h *GoogleHandler) HandlerSignIn(w http.ResponseWriter, r *http.Request) {}

