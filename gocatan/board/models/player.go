package models

import (
	"github.com/google/uuid"
)

type Player struct {
	Uuid uuid.UUID `json:"uuid"`
}
