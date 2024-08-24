package api

import (
	"encoding/json"
	"gocatan/board"
	builders "gocatan/board/builders"
	"gocatan/config"
	"net/http"
)

type GameBoard struct {
	Tiles    []board.ConcreteHexagonTile
	Vertices []board.Vertice
}

func BoardHandler(w http.ResponseWriter, r *http.Request) {
	// can use this to pass configuation from the client in the future
	engine := board.Engine{
		HexSideSize:    config.HexagonImageHeight / 2,
		HexTotalWidth:  config.HexagonImageWidth,
		HexTotalHeight: config.HexagonImageHeight,
	}

	regularMap := board.RegularBoard()

	concreteHexTiles, _ := engine.BuildHexagons(&regularMap)
	vertices := builders.BuildVertices(concreteHexTiles)

	gb := GameBoard{
		Tiles:    concreteHexTiles,
		Vertices: vertices,
	}

	resp, _ := json.Marshal(gb)

	w.Write(resp)
}
