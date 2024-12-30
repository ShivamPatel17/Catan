package models

import (
	"math"

	"github.com/google/uuid"
)

type Orientation int

const (
	Incline  Orientation = 1 // for the top left and bottom right of a hexagon
	Decline  Orientation = 2 // for the top right and bottom left a hexagon
	Vertical Orientation = 3 // for the left and right of a hexagon
)

type Edge struct {
	Uuid        uuid.UUID   `json:"uuid"`
	X           float64     `json:"x"`
	Y           float64     `json:"y"`
	Orientation Orientation `json:"orientation"`
}

func NewEdge(x float64, y float64, o Orientation) Edge {
	uuid, _ := uuid.NewUUID()
	return Edge{
		Uuid:        uuid,
		X:           x,
		Y:           y,
		Orientation: o,
	}
}

func IsSameEdge(e1 Edge, e2 Edge, tolerance float64) bool {
	return math.Abs(e1.X-e2.X) < tolerance && math.Abs(e1.Y-e2.Y) < tolerance
}
