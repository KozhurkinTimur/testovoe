package apiserver

import (
	//"io"
	"net/http"
	//"testovoe/internal/apiserver/rest/google/handlers"
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

func (s *APIServer) Start() error {
	s.routers()
	return http.ListenAndServe(s.Address, s.Router)
}

func (s *APIServer) routers() {
	s.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

	})
	
}
