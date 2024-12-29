package models

import (
	"github.com/google/uuid"
)

type Player struct {
	id uuid.UUID `json:"id"`
}
