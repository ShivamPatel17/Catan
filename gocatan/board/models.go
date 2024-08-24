package board

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
	direction       Direction
	relativeHexTile RelativeHexagonTile
}
type RelativeHexagonTile struct {
	AdjacentTiles []DirectionalHexagonTile
	Concrete      *ConcreteHexagonTile
}

type ConcreteHexagonTile struct {
	id       int64
	X        int
	Y        int
	Resource Resource
}

type Vertice struct {
	id int
	X  int
	Y  int
}
