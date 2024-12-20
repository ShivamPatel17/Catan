package api

import (
	"context"
	"encoding/json"
	builders "gocatan/board/builders/hexagon"
	"gocatan/board/models"
	"gocatan/config"
	"net/http"
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

	gb := models.GameBoard{
		Tiles:             concreteHexTiles,
		Vertices:          vertices,
		AdjacentVerticies: adjVerticies,
		Edges:             edges,
	}
	return gb
}
