package models

import (
	"github.com/google/uuid"
)

type Vertice struct {
	Id uuid.UUID `json:"id"`
	X  float64   `json:"x"`
	Y  float64   `json:"y"`
}

func NewVertice(x float64, y float64) Vertice {
	uuid, _ := uuid.NewUUID()

	return Vertice{
		Id: uuid,
		X:  x,
		Y:  y,
	}
}
