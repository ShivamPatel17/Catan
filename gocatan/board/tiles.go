package board

import (
	"encoding/json"
	"net/http"
)

type Direction int

const (
	TopLeft     Direction = 0
	TopRight    Direction = 1
	MiddleLeft  Direction = 2
	MiddleRight Direction = 3
	BottomLeft  Direction = 4
	BottomRight Direction = 5
)

type DirectionalHexagonTile struct {
	direction       Direction
	relativeHexTile RelativeHexagonTile
}
type RelativeHexagonTile struct {
	adjacentTiles []DirectionalHexagonTile
	concrete      *ConcreteHexagonTile
}

type ConcreteHexagonTile struct {
	X     int
	Y     int
	Color int
}

func HexagonHandler(w http.ResponseWriter, r *http.Request) {
	// can use this to pass configuation from the client in the future
	relativeHexTiles := RelativeHexagonTile{
		adjacentTiles: []DirectionalHexagonTile{
			{
				direction: TopRight,
				relativeHexTile: RelativeHexagonTile{
					adjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleRight,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       TopRight,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       TopLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: TopLeft,
				relativeHexTile: RelativeHexagonTile{
					adjacentTiles: []DirectionalHexagonTile{
						{
							direction:       TopLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       MiddleLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: MiddleRight,
				relativeHexTile: RelativeHexagonTile{
					adjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleRight,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: MiddleLeft,
				relativeHexTile: RelativeHexagonTile{
					adjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: BottomRight,
				relativeHexTile: RelativeHexagonTile{
					adjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleRight,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       BottomRight,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       BottomLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
			{
				direction: BottomLeft,
				relativeHexTile: RelativeHexagonTile{
					adjacentTiles: []DirectionalHexagonTile{
						{
							direction:       MiddleLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
						{
							direction:       BottomLeft,
							relativeHexTile: RelativeHexagonTile{},
						},
					},
				},
			},
		},
		concrete: &ConcreteHexagonTile{
			X: 600,
			Y: 450,
		},
	}

	hexScale := 0.3
	hexTotalHeight := 508
	hexTotalWidth := 440
	hexSideLength := hexTotalHeight / 2 // HEX PIXEL SIDE

	engine := Engine{
		hexSideSize:    int(hexScale * float64(hexSideLength)),
		hexTotalWidth:  int(hexScale * float64(hexTotalWidth)),
		hexTotalHeight: int(hexScale * float64(hexTotalHeight)),
	}
	concreteHexTiles, _ := engine.BuildMap(&relativeHexTiles)

	resp, _ := json.Marshal(concreteHexTiles)

	w.Write(resp)
}
