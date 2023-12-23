package handlers

import (
	"net/http"
)

type GoogleHandler struct {
	Id int
}

func New() *GoogleHandler {
	return &GoogleHandler{
		Id: 1,
	}
}

func (h *GoogleHandler) handlerRegistration(w http.ResponseWriter, r *http.Request) {

}

func (h *GoogleHandler) handlerSignIn(w http.ResponseWriter, r *http.Request) {

}

