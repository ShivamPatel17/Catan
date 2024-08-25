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
	Direction       Direction
	RelativeHexTile RelativeHexagonTile
}
type RelativeHexagonTile struct {
	AdjacentTiles []DirectionalHexagonTile
	Concrete      *ConcreteHexagonTile
}

type ConcreteHexagonTile struct {
	Id       int64
	X        float64
	Y        float64
	Resource Resource
}

type Vertice struct {
	Id int64
	X  float64
	Y  float64
}
