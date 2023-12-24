package handlerssms

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

func (h *SMSHandler) HandlerRegistration(w http.ResponseWriter, r *http.Request) {
	// TODO: verify uuid in db, registartion id in another provider and db  
	// TODO: verify correctly login and password
	// TODO: send sms with confirmation code
	// TODO: create user in db
}

func (h *SMSHandler) HandlerSignIn(w http.ResponseWriter, r *http.Request) {
	// TODO: verify uuid in db, login and password
	// TODO: verify confirmation code
	// TODO: login user
}

func (h *SMSHandler) HandlerPay(w http.ResponseWriter, r *http.Request) {
	// TODO: get number, key uuid, some data
	// TODO: send sms with confirmation code
}

func (h *SMSHandler) HandlerConfirmations(w http.ResponseWriter, r *http.Request) {
	// TODO: get uuid and code
	// TODO: verify code and uuid
}