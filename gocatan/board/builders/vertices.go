package builders

import (
	models "gocatan/board/models"
	"gocatan/internal"
	mathhelper "gocatan/internal"
)

func (e *HexagonEngine) BuildVertices(concreteTiles []models.ConcreteHexagonTile) []models.Vertice {
	allVertices := e.buildAllVertices(concreteTiles)
	dedupedVertices := dedup(allVertices)
	return dedupedVertices
}

func (e *HexagonEngine) buildAllVertices(concreteTiles []models.ConcreteHexagonTile) []models.Vertice {
	vertices := make([]models.Vertice, 0)
	for _, concreteTile := range concreteTiles {

		x, y := concreteTile.X, concreteTile.Y
		// top vertice
		vertices = append(vertices,
			models.Vertice{
				X: x,
				Y: y - (e.HexSideSize),
			},
		)
		// top right
		vertices = append(vertices,
			models.Vertice{
				X: x + mathhelper.HeightOfEqualateralTriangle(e.HexSideSize),
				Y: y - (e.HexSideSize / 2),
			},
		)
		// top left
		vertices = append(vertices,
			models.Vertice{
				X: x - mathhelper.HeightOfEqualateralTriangle(e.HexSideSize),
				Y: y - (e.HexSideSize / 2),
			},
		)
		// bottom left
		vertices = append(vertices,
			models.Vertice{
				X: x - mathhelper.HeightOfEqualateralTriangle(e.HexSideSize),
				Y: y + (e.HexSideSize / 2),
			},
		)
		// bottom right
		vertices = append(vertices,
			models.Vertice{
				X: x + mathhelper.HeightOfEqualateralTriangle(e.HexSideSize),
				Y: y + (e.HexSideSize / 2),
			},
		)
		// bottom
		vertices = append(vertices,
			models.Vertice{
				X: x,
				Y: y + e.HexSideSize,
			},
		)
	}
	return vertices
}

func dedup(vertices []models.Vertice) []models.Vertice {
	tolerance := 1.0
	dedupedVerts := make([]models.Vertice, 0)
	for _, vert := range vertices {
		dup := false
		for _, dedupedVert := range dedupedVerts {
			if internal.IsSameVertice(vert, dedupedVert, tolerance) {
				dup = true
				break
			}
		}
		if !dup {
			dedupedVerts = append(dedupedVerts, vert)
		}
	}
	return dedupedVerts
}
