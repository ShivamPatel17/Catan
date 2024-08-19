package board

import "math"

type Engine struct {
	hexSideSize    float32
	hexTotalWidth  float32
	hexTotalHeight float32
}

func (e *Engine) BuildMap(relativeTiles *RelativeHexagonTile) ([](ConcreteHexagonTile), error) {
	concreteTiles := make([]ConcreteHexagonTile, 0)
	ConcreteTile := relativeTiles.Concrete
	if ConcreteTile != nil {
		concreteTiles = append(concreteTiles, ConcreteHexagonTile{
			X: ConcreteTile.X,
			Y: ConcreteTile.Y,
		})
	}
	e.buildRelativeHexTiles(&relativeTiles.AdjacentTiles, &concreteTiles[0], &concreteTiles)
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
				X: relativeTo.X + int(math.Sqrt(3)*float64(e.hexSideSize)/2.0),
				Y: relativeTo.Y - int((3*e.hexSideSize)/2.0),
			}
		case TopLeft:
			ConcreteTile = ConcreteHexagonTile{
				X: relativeTo.X - int(math.Sqrt(3)*float64(e.hexSideSize)/2.0),
				Y: relativeTo.Y - int((3*e.hexSideSize)/2.0),
			}
		case MiddleRight:
			ConcreteTile = ConcreteHexagonTile{
				X: relativeTo.X + int(e.hexTotalWidth),
				Y: relativeTo.Y,
			}
		case MiddleLeft:
			ConcreteTile = ConcreteHexagonTile{
				X:        relativeTo.X - int(e.hexTotalWidth),
				Y:        relativeTo.Y,
				Resource: Sheep,
			}
		case BottomRight:
			ConcreteTile = ConcreteHexagonTile{
				X: relativeTo.X + int(math.Sqrt(3)*float64(e.hexSideSize)/2.0),
				Y: relativeTo.Y + int((3*e.hexSideSize)/2.0),
			}
		case BottomLeft:
			ConcreteTile = ConcreteHexagonTile{
				X: relativeTo.X - int(math.Sqrt(3)*float64(e.hexSideSize)/2.0),
				Y: relativeTo.Y + int((3*e.hexSideSize)/2.0),
			}
		}

		*concreteTiles = append(*concreteTiles, ConcreteTile)
		e.buildRelativeHexTiles(&directionalTile.relativeHexTile.AdjacentTiles, &ConcreteTile, concreteTiles)
	}
}
