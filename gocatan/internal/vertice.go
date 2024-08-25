package internal

import (
	"gocatan/board/models"
	"math"
)

// returns true if vert1 is the same as vert2
func IsSameVertice(vert1 models.Vertice, vert2 models.Vertice, tolerance float64) bool {
	return math.Abs(vert1.X-vert2.X) < tolerance && math.Abs(vert1.Y-vert2.Y) < tolerance
}
