package apiserver

import (
	//"io"
	"net/http"
	"testovoe/internal/apiserver/rest/google/handlersgoogle"
	"testovoe/internal/apiserver/rest/sms/handlerssms"
	"testovoe/internal/apiserver/rest/guest/handlersjwt"
	"testovoe/internal/config"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Address string
	Router  *mux.Router
}

func New(config *config.Config) *APIServer {
	return &APIServer{
		Address: config.Address,
		Router:  mux.NewRouter(),
	}
}
// refactor: create intrface for handlers, implement it in handlers packages, inject in Start
func (s *APIServer) Start() error {
	s.routers()
	return http.ListenAndServe(s.Address, s.Router)
}
// refactor: create intrface for handlers, implement it in handlers packages, inject in routers
func (s *APIServer) routers() {
	googleHandler := handlersgoogle.New()
	smsHandler := handlerssms.New()
	jwtHandler := handlersjwt.New()
	// s.Router.HandleFunc("/jwt", handlersjwt.LogHandler(jwtHandler.HandlerGuest)).Methods("GET")
	s.Router.HandleFunc("/jwt/guest", jwtHandler.HandlerGuest).Methods("GET")
	s.Router.HandleFunc("/jwt/signin", jwtHandler.HandlerSignIn).Methods("POST")

	s.Router.HandleFunc("/google/registration", googleHandler.HandlerRegistration).Methods("POST")
	s.Router.HandleFunc("/google/signin", googleHandler.HandlerSignIn).Methods("GET")

	s.Router.HandleFunc("/sms/registration", smsHandler.HandlerRegistration).Methods("POST")
	s.Router.HandleFunc("/sms/pay", smsHandler.HandlerPay).Methods("GET")
	s.Router.HandleFunc("/sms/confirmations", smsHandler.HandlerConfirmations).Methods("POST")
	
}
