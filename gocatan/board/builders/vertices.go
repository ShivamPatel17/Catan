package board

import "gocatan/board"

func BuildVertices(concreteTiles []board.ConcreteHexagonTile) []board.Vertice {
	vertices := make([]board.Vertice, 0)
	for _, concreteTile := range concreteTiles {
		vertices = append(vertices,
			board.Vertice{
				X: concreteTile.X,
				Y: concreteTile.Y,
			},
		)
	}
	return vertices
}
