package board

import (
	"encoding/json"
	"net/http"
)

type hexagonTile struct {
	X int
	Y int
}

func HexagonHandler(w http.ResponseWriter, r *http.Request) {
	// Create a slice of anonymous structs with fields X and Y
	points := []hexagonTile{
		{X: 100, Y: 200},
		{X: 300, Y: 400},
	}
	hex, _ := json.Marshal(points)
	w.Write(hex)
}
