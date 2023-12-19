package usecase

import (
	"jwt-service/internal/domain/entity"
	"jwt-service/internal/lib/uuid"
	"log/slog"
)

func NewUserUseCase(log *slog.Logger) *entity.User {
	log.Info("Created new user")
	return &entity.User{
		ID:     uuid.GenerateUUID(log),
		Secret: "sasa",
	}
}
