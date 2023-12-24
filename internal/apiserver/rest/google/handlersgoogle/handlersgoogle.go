package handlersgoogle

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

func (h *GoogleHandler) HandlerRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO: verify uuid in db, registartion id in another provider and db  
	// TODO: verify correctly login and password
	// TODO: create user in db 
	// TODO: OAuth...
}

func (h *GoogleHandler) HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	// TODO: get OAuth token, key uuid, some data
}

