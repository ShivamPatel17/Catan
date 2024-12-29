package models

import (
	"math"

	"github.com/google/uuid"
)

type Vertice struct {
	Uuid     uuid.UUID `json:"uuid"`
	X        float64   `json:"x"`
	Y        float64   `json:"y"`
	Player   Player    `json:"player"`
	Building `json:"building"`
}

type Building int

const (
	Settlement Building = 1
	City       Building = 2
)

func NewVertice(x float64, y float64) Vertice {
	uuid, _ := uuid.NewUUID()

	return Vertice{
		Uuid: uuid,
		X:    x,
		Y:    y,
	}
}

// returns true if vert1 is the same as vert2
func IsSameVertice(vert1 Vertice, vert2 Vertice, tolerance float64) bool {
	return math.Abs(vert1.X-vert2.X) < tolerance && math.Abs(vert1.Y-vert2.Y) < tolerance
}
