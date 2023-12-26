package main

import (
	"fmt"
	"testovoe/internal/apiserver"
	"testovoe/internal/config"

	// "testovoe/internal/models"
	// "testovoe/internal/lib/jwt"

	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting auth service", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	s := apiserver.New(cfg)
	err := s.Start()
	if err != nil {
		log.Error("doesnt start server ",err)
	}
	log.Debug("1111")

}

func init() {
	err := godotenv.Load("../../local.env")

	if err != nil {
		fmt.Println("Error")
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
