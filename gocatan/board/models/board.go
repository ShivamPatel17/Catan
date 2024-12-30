package models

import "github.com/google/uuid"

type GameBoard struct {
	Tiles             map[uuid.UUID]*ConcreteHexagonTile `json:"tiles"`
	Vertices          map[uuid.UUID]*Vertice             `json:"vertices"`
	AdjacentVerticies map[uuid.UUID][]Vertice            `json:"adjacent_vertices"`
	Edges             map[uuid.UUID]*Edge                `json:"edges"`
	AdjacentEdges     map[uuid.UUID][]Edge               `json:"adjacent_edges"`
}
