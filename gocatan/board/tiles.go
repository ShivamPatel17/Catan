package board

import (
	"encoding/json"
	"gocatan/config"
	"net/http"
)

func HexagonHandler(w http.ResponseWriter, r *http.Request) {
	// can use this to pass configuation from the client in the future
	engine := Engine{
		hexSideSize:    config.HexagonImageHeight / 2,
		hexTotalWidth:  config.HexagonImageWidth,
		hexTotalHeight: config.HexagonImageHeight,
	}

	regularMap := RegularBoard()

	concreteHexTiles, _ := engine.BuildMap(&regularMap)

	resp, _ := json.Marshal(concreteHexTiles)

	w.Write(resp)
}
