package models

import (
	"testovoe/internal/config"

	"github.com/google/uuid"
)

type Authorization struct {
	Id     uuid.UUID
	Secret string
	Role   string
	//Provider int
}

func New(config *config.Config) *Authorization {
	return &Authorization{
		Id:     uuid.New(),
		Secret: config.Secret,
		Role:   "",
	}
}


