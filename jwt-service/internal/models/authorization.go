package models

import (
	"github.com/google/uuid"
)

type Authorization struct {
	Id uuid.UUID `json:"id"`
	Secret string `json:""`
	Role string
	Provider
}


