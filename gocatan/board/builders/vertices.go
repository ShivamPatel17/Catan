package builders

import (
	models "gocatan/board/models"
	"math"
)

func (e *HexagonEngine) BuildVertices(concreteTiles []models.ConcreteHexagonTile) []models.Vertice {
	vertices := make([]models.Vertice, 0)
	for _, concreteTile := range concreteTiles {
		// TODO use a real for loop lol
		arr := []int{1, 2, 3, 4, 5, 6}
		for range arr {
			vertices = append(vertices, models.Vertice{
				X: 0,
				Y: 0,
			})
			x, y := concreteTile.X, concreteTile.Y
			// top vertice
			vertices = append(vertices,
				models.Vertice{
					X: x,
					Y: y - (int(e.HexSideSize)),
				},
			)
			// top right
			vertices = append(vertices,
				models.Vertice{
					X: x + int((math.Sqrt(3)/2.0)*float64(e.HexSideSize)),
					Y: y - (int(e.HexSideSize) / 2),
				},
			)
			// top left
			vertices = append(vertices,
				models.Vertice{
					X: x - int((math.Sqrt(3)/2.0)*float64(e.HexSideSize)),
					Y: y - (int(e.HexSideSize) / 2),
				},
			)
			// bottom left
			vertices = append(vertices,
				models.Vertice{
					X: x - int((math.Sqrt(3)/2.0)*float64(e.HexSideSize)),
					Y: y + (int(e.HexSideSize) / 2),
				},
			)
			// bottom right
			vertices = append(vertices,
				models.Vertice{
					X: x + int((math.Sqrt(3)/2.0)*float64(e.HexSideSize)),
					Y: y + (int(e.HexSideSize) / 2),
				},
			)
			// bottom
			vertices = append(vertices,
				models.Vertice{
					X: x,
					Y: y + (int(e.HexSideSize)),
				},
			)
		}
	}
	return vertices
}
