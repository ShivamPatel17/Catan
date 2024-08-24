package builders

import models "gocatan/board/models"

func BuildVertices(concreteTiles []models.ConcreteHexagonTile) []models.Vertice {
	vertices := make([]models.Vertice, 0)
	for _, concreteTile := range concreteTiles {
		vertices = append(vertices,
			models.Vertice{
				X: concreteTile.X,
				Y: concreteTile.Y,
			},
		)
	}
	return vertices
}
