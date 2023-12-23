package handlers

import (
	"net/http"
)

type SMSHandler struct{
	Id int
}

func New() *SMSHandler {
	return &SMSHandler{
		Id: 2,
	}
}

func (h *SMSHandler) handlerRegistration(w http.ResponseWriter, r *http.Request, login string, passwrod string, id uuid) {
	// TODO: verify uuid in db, registartion id in another provider and db. if  
	// TODO: verify correctly login and password
	// TODO: send sms with confirmation code
	// TODO: create user in db


}

func (h *SMSHandler) handlerSignIn(w http.ResponseWriter, r *http.Request) {
	
}

// func (h *SMSHandler) handlerPay(w http.ResponseWriter, r *http.Request, l login, id uuid) {

// }

// func (h *SMSHandler) handlerConfirmations(w http.ResponseWriter, r *http.Request, id uuid, code string) {

// }