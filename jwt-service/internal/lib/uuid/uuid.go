package uuid

import (
	"log/slog"

	"github.com/google/uuid"
)

func GenerateUUID(log *slog.Logger) uuid.UUID {
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		log.Error("error generating UUIDV4", err)
	}
	log.Info("generated UUIDV4")
	return uuidV4
}