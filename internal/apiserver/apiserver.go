package apiserver

import (
	//"io"
	"encoding/json"
	"log"
	"net/http"
	"testovoe/internal/apiserver/rest/google/handlersgoogle"
	"testovoe/internal/apiserver/rest/guest/handlersjwt"
	"testovoe/internal/apiserver/rest/sms/handlerssms"
	"testovoe/internal/config"
	"testovoe/internal/models"
	"testovoe/internal/storage/users"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Address string
}

func New(config *config.Config) *APIServer {
	return &APIServer{
		Address: config.Address,
	}
}

// refactor: create intrface for handlers, implement it in handlers packages, inject in Start
func (s *APIServer) Start() error {
	usersRepo, err := users.NewUsersRepository()
	if err != nil {
		log.Fatal(err)
	}
	if err := usersRepo.Save(&models.User{Login: "user"}); err != nil {
		log.Fatal(err)
	}
	googleHandler := handlersgoogle.New()
	smsHandler := handlerssms.New(usersRepo)
	jwtHandler := handlersjwt.New()

	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	}).Methods("GET")

	router.HandleFunc("/jwt/guest", jwtHandler.HandlerGuest).Methods("GET")
	router.HandleFunc("/jwt/signin", jwtHandler.HandlerSignIn).Methods("POST")

	router.HandleFunc("/google/registration", googleHandler.HandlerRegistration).Methods("POST")
	router.HandleFunc("/google/signin", googleHandler.HandlerSignIn).Methods("GET")

	router.HandleFunc("/sms/registration", smsHandler.HandlerRegistration).Methods("POST")
	router.HandleFunc("/sms/pay", smsHandler.HandlerPay).Methods("GET")
	router.HandleFunc("/sms/confirmations", smsHandler.HandlerConfirmations).Methods("POST")

	log.Printf("server starting on: %s", s.Address)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv.ListenAndServe()
}
