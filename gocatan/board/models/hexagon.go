package models

type Direction int

const (
	TopLeft     Direction = 0
	TopRight    Direction = 1
	MiddleLeft  Direction = 2
	MiddleRight Direction = 3
	BottomLeft  Direction = 4
	BottomRight Direction = 5
)

type Resource string

const (
	Sheep Resource = "sheep"
	Wheat Resource = "wheat"
	Ore   Resource = "ore"
	Wood  Resource = "wood"
	Brick Resource = "brick"
)

type DirectionalHexagonTile struct {
	Direction       Direction           `json:"direction"`
	RelativeHexTile RelativeHexagonTile `json:"relative_hex_tile"`
}

type RelativeHexagonTile struct {
	AdjacentTiles []DirectionalHexagonTile `json:"adjacent_tiles"`
	Concrete      *ConcreteHexagonTile     `json:"concrete"`
}

type ConcreteHexagonTile struct {
	Id       int64    `json:"id"`
	X        float64  `json:"x"`
	Y        float64  `json:"y"`
	Resource Resource `json:"resource"`
}
