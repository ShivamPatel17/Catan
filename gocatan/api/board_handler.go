package api

import (
	"context"
	"encoding/json"
	builders "gocatan/board/builders/hexagon"
	"gocatan/board/models"
	"gocatan/config"
	"net/http"

	"github.com/google/uuid"
)

func BoardHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.NewConfig()
	ctx := context.Background()
	gb := BuildBoard(ctx, cfg)
	resp, _ := json.Marshal(gb)

	w.Write(resp)
}

func BuildBoard(_ context.Context, cfg config.Config) models.GameBoard {
	// can use this to pass configuation from the client in the future
	engine := builders.NewHexagonEngine(cfg)

	regularMap := models.RegularBoard()

	concreteHexTiles, _ := engine.BuildHexagons(&regularMap)
	vertices := engine.BuildVertices(concreteHexTiles)
	adjVerticies := engine.BuildAdjacentVerticesMap(vertices)
	edges := engine.BuildEdges(concreteHexTiles)

	verticesMap := make(map[uuid.UUID]*models.Vertice)

	for _, v := range vertices {
		verticesMap[v.Uuid] = &v
	}

	concreteHexTilesMap := make(map[uuid.UUID]*models.ConcreteHexagonTile)

	for _, t := range concreteHexTiles {
		concreteHexTilesMap[t.Uuid] = t
	}

	edgesMap := make(map[uuid.UUID]*models.Edge)

	for _, e := range edges {
		edgesMap[e.Uuid] = &e
	}

	gb := models.GameBoard{
		Tiles:             concreteHexTilesMap,
		Vertices:          verticesMap,
		AdjacentVerticies: adjVerticies,
		Edges:             edgesMap,
	}
	return gb
}
