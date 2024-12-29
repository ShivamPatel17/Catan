package hexagon

import (
	models "gocatan/board/models"
	mathhelper "gocatan/internal/math"

	"github.com/google/uuid"
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
		vertices = append(vertices,
			// top
			models.NewVertice(x, y-e.HexSideSize),
			// top right
			models.NewVertice(x+mathhelper.HeightOfEqualateralTriangle(e.HexSideSize), y-(e.HexSideSize/2)),
			// top left
			models.NewVertice(x-mathhelper.HeightOfEqualateralTriangle(e.HexSideSize), y-(e.HexSideSize/2)),
			// bottom left
			models.NewVertice(x-mathhelper.HeightOfEqualateralTriangle(e.HexSideSize), y+(e.HexSideSize/2)),
			// bottom right
			models.NewVertice(x+mathhelper.HeightOfEqualateralTriangle(e.HexSideSize), y+(e.HexSideSize/2)),
			// bottom
			models.NewVertice(x, y+e.HexSideSize),
		)
	}
	return vertices
}

func dedup(vertices []models.Vertice) []models.Vertice {
	Tolerance := 1.0
	dedupedVerts := make([]models.Vertice, 0)
	for _, vert := range vertices {
		dup := false
		for _, dedupedVert := range dedupedVerts {
			if models.IsSameVertice(vert, dedupedVert, Tolerance) {
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

func (e *HexagonEngine) BuildAdjacentVerticesMap(vertices []models.Vertice) map[uuid.UUID][]models.Vertice {
	adjVertsMap := make(map[uuid.UUID][]models.Vertice)

	for _, vert1 := range vertices {
		adjArr := make([]models.Vertice, 0)
		for _, vert2 := range vertices {
			if e.isAdjacentVertice(vert2, vert1) {
				adjArr = append(adjArr, vert2)
			}
		}
		adjVertsMap[vert1.Uuid] = adjArr
	}

	return adjVertsMap
}

// returns ture if the v1 and v2 are adjacent to each other
func (e *HexagonEngine) isAdjacentVertice(v1 models.Vertice, v2 models.Vertice) bool {
	Tolerance := 1.0
	height := mathhelper.HeightOfEqualateralTriangle(e.HexSideSize)

	// check if v2 is top  of v1
	if mathhelper.WithinTolerance(v1.X, v2.X, Tolerance) && mathhelper.WithinTolerance(v1.Y, v2.Y+e.HexSideSize, Tolerance) {
		return true
	}
	// check if v2 is top right of v1
	if mathhelper.WithinTolerance(v1.X, v2.X-height, Tolerance) && mathhelper.WithinTolerance(v1.Y, v2.Y+(e.HexSideSize/2.0), Tolerance) {
		return true
	}
	// check if v2 is top left of v1
	if mathhelper.WithinTolerance(v1.X, v2.X+height, Tolerance) && mathhelper.WithinTolerance(v1.Y, v2.Y+(e.HexSideSize/2.0), Tolerance) {
		return true
	}

	// check if v2 is bottom of v1
	if mathhelper.WithinTolerance(v1.X, v2.X, Tolerance) && mathhelper.WithinTolerance(v1.Y, v2.Y-e.HexSideSize, Tolerance) {
		return true
	}

	// check if v2 is bottom right of v1
	if mathhelper.WithinTolerance(v1.X, v2.X-height, Tolerance) && mathhelper.WithinTolerance(v1.Y, v2.Y-(e.HexSideSize/2.0), Tolerance) {
		return true
	}

	// check if v2 is bottom left of v1
	if mathhelper.WithinTolerance(v1.X, v2.X+height, Tolerance) && mathhelper.WithinTolerance(v1.Y, v2.Y-(e.HexSideSize/2.0), Tolerance) {
		return true
	}
	return false
}
