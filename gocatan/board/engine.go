package board

import (
	"math"
	"math/rand"
)

type Engine struct {
	HexSideSize    float32
	HexTotalWidth  float32
	HexTotalHeight float32
}

func (e *Engine) BuildHexagons(relativeTiles *RelativeHexagonTile) ([](ConcreteHexagonTile), error) {
	concreteTiles := make([]ConcreteHexagonTile, 0)
	ConcreteTile := relativeTiles.Concrete
	if ConcreteTile != nil {
		concreteTiles = append(concreteTiles, ConcreteHexagonTile{
			X: ConcreteTile.X,
			Y: ConcreteTile.Y,
		})
	}
	e.buildRelativeHexTiles(&relativeTiles.AdjacentTiles, &concreteTiles[0], &concreteTiles)
	assignResources(&concreteTiles)
	return concreteTiles, nil
}

func (e *Engine) buildRelativeHexTiles(tiles *[]DirectionalHexagonTile, relativeTo *ConcreteHexagonTile, concreteTiles *[]ConcreteHexagonTile) {
	if relativeTo == nil {
		return
	}

	for _, directionalTile := range *tiles {
		var ConcreteTile ConcreteHexagonTile
		switch directionalTile.direction {
		case TopRight:
			ConcreteTile = ConcreteHexagonTile{
				X: relativeTo.X + int(math.Sqrt(3)*float64(e.HexSideSize)/2.0),
				Y: relativeTo.Y - int((3*e.HexSideSize)/2.0),
			}
		case TopLeft:
			ConcreteTile = ConcreteHexagonTile{
				X: relativeTo.X - int(math.Sqrt(3)*float64(e.HexSideSize)/2.0),
				Y: relativeTo.Y - int((3*e.HexSideSize)/2.0),
			}
		case MiddleRight:
			ConcreteTile = ConcreteHexagonTile{
				X: relativeTo.X + int(e.HexTotalWidth),
				Y: relativeTo.Y,
			}
		case MiddleLeft:
			ConcreteTile = ConcreteHexagonTile{
				X:        relativeTo.X - int(e.HexTotalWidth),
				Y:        relativeTo.Y,
				Resource: Sheep,
			}
		case BottomRight:
			ConcreteTile = ConcreteHexagonTile{
				X: relativeTo.X + int(math.Sqrt(3)*float64(e.HexSideSize)/2.0),
				Y: relativeTo.Y + int((3*e.HexSideSize)/2.0),
			}
		case BottomLeft:
			ConcreteTile = ConcreteHexagonTile{
				X: relativeTo.X - int(math.Sqrt(3)*float64(e.HexSideSize)/2.0),
				Y: relativeTo.Y + int((3*e.HexSideSize)/2.0),
			}
		}

		*concreteTiles = append(*concreteTiles, ConcreteTile)
		e.buildRelativeHexTiles(&directionalTile.relativeHexTile.AdjacentTiles, &ConcreteTile, concreteTiles)
	}
}

func assignResources(concreteTiles *[]ConcreteHexagonTile) {
	for i := range *concreteTiles {
		(*concreteTiles)[i].Resource = getRandomResource()
	}
}
func getRandomResource() Resource {
	// Seed the random number generator to ensure different results on each run

	// Create a slice of all possible resources
	resources := []Resource{Sheep, Wheat, Ore, Wood, Brick}

	// Select a random index
	randomIndex := rand.Intn(len(resources))

	// Return the randomly selected resource
	return resources[randomIndex]
}
