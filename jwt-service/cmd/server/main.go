package main

import (
	"fmt"
	"jwt-service/internal/config"
	"jwt-service/internal/domain/usecase"
	"jwt-service/internal/lib/jwt"
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

	log.Info("starting jwt service", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	token, err := jwt.NewToken(usecase.NewUserUseCase(log), log)

	fmt.Println(token, err)
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
