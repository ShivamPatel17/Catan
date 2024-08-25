package models

import "github.com/google/uuid"

type GameBoard struct {
	Tiles             []ConcreteHexagonTile   `json:"tiles"`
	Vertices          []Vertice               `json:"vertices"`
	AdjacentVerticies map[uuid.UUID][]Vertice `json:"adjacent_vertices"`
	Edges             []Edge                  `json:"edges"`
	AdjacentEdges     map[uuid.UUID][]Edge    `json:"adjacent_edges"`
}
