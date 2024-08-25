package models

import (
	"math"

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

// returns true if vert1 is the same as vert2
func IsSameVertice(vert1 Vertice, vert2 Vertice, tolerance float64) bool {
	return math.Abs(vert1.X-vert2.X) < tolerance && math.Abs(vert1.Y-vert2.Y) < tolerance
}
