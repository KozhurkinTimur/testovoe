package main

import (
	"testovoe/internal/apiserver"

	"log"
)

func main() {
	log.Print("starting test service")
	log.Print("debug messages are enabled")

	s := apiserver.New()
	err := s.Start()
	if err != nil {
		log.Fatalf("doesnt start server %v", err)
	}
	log.Print("1111")
}
