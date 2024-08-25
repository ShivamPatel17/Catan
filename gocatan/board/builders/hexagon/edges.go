package hexagon

import (
	"gocatan/board/models"
	mathhelper "gocatan/internal/math"
	"math"
)

func (e *HexagonEngine) BuildEdges(concreteTiles []models.ConcreteHexagonTile) []models.Edge {
	allEdges := e.buildAllEdges(concreteTiles)
	dedupedEdges := e.dedupEdges(allEdges)
	return dedupedEdges
}

func (e *HexagonEngine) buildAllEdges(concreteTiles []models.ConcreteHexagonTile) []models.Edge {
	edges := make([]models.Edge, 0)
	for _, tile := range concreteTiles {
		x, y, height := tile.X, tile.Y, mathhelper.HeightOfEqualateralTriangle(e.HexSideSize)
		vert := math.Sqrt((height * height) - ((height / 2.0) * (height / 2.0)))
		edges = append(edges,
			// top left
			models.NewEdge(x-height/2.0, y-vert, models.Incline),
			// top right
			models.NewEdge(x+height/2.0, y-vert, models.Decline),
			// bottom left
			models.NewEdge(x-height/2.0, y+vert, models.Decline),
			// bottom right
			models.NewEdge(x+height/2.0, y+vert, models.Incline),
			// right
			models.NewEdge(x+height, y, models.Vertical),
			// left
			models.NewEdge(x-height, y, models.Incline),
		)
	}
	return edges
}

func (e *HexagonEngine) dedupEdges(edges []models.Edge) []models.Edge {
	dedupedEdges := make([]models.Edge, 0)
	tolerance := 1.0
	for _, e := range edges {
		dup := false
		for _, de := range dedupedEdges {
			if models.IsSameEdge(e, de, tolerance) {
				dup = true
				break
			}
		}
		if !dup {
			dedupedEdges = append(dedupedEdges, e)
		}
	}

	return dedupedEdges
}
