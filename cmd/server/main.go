package main

import (
	"testovoe/internal/apiserver"
	"testovoe/internal/config"

	// "testovoe/internal/models"
	// "testovoe/internal/lib/jwt"

	"log"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

var ()

func main() {
	cfg := config.MustLoad()

	log.Print("starting test service")
	log.Print("debug messages are enabled")

	s := apiserver.New(cfg)
	err := s.Start()
	if err != nil {
		log.Fatalf("doesnt start server %v", err)
	}
	log.Print("1111")
}
