package board

import "math"

type Engine struct {
	hexSideSize    int
	hexTotalWidth  int
	hexTotalHeight int
}

func (e *Engine) BuildMap(relativeTiles *RelativeHexagonTile) ([](ConcreteHexagonTile), error) {
	concreteTiles := make([]ConcreteHexagonTile, 0)
	concreteTile := relativeTiles.concrete
	if concreteTile != nil {
		concreteTiles = append(concreteTiles, ConcreteHexagonTile{
			X: concreteTile.X,
			Y: concreteTile.Y,
		})
	}
	e.buildRelativeHexTiles(&relativeTiles.adjacentTiles, &concreteTiles[0], &concreteTiles)
	return concreteTiles, nil
}

func (e *Engine) buildRelativeHexTiles(tiles *[]DirectionalHexagonTile, relativeTo *ConcreteHexagonTile, concreteTiles *[]ConcreteHexagonTile) {
	if relativeTo == nil {
		return
	}

	for _, directionalTile := range *tiles {
		var concreteTile ConcreteHexagonTile
		switch directionalTile.direction {
		case TopRight:
			concreteTile = ConcreteHexagonTile{
				X: relativeTo.X + int(math.Sqrt(3)*float64(e.hexSideSize)/2.0),
				Y: relativeTo.Y - (3*e.hexSideSize)/2.0,
			}
		case TopLeft:
			concreteTile = ConcreteHexagonTile{
				X: relativeTo.X - int(math.Sqrt(3)*float64(e.hexSideSize)/2.0),
				Y: relativeTo.Y - (3*e.hexSideSize)/2.0,
			}
		case MiddleRight:
			concreteTile = ConcreteHexagonTile{
				X: relativeTo.X + e.hexTotalWidth,
				Y: relativeTo.Y,
			}
		case MiddleLeft:
			concreteTile = ConcreteHexagonTile{
				X: relativeTo.X - e.hexTotalWidth,
				Y: relativeTo.Y,
			}
		case BottomRight:
			concreteTile = ConcreteHexagonTile{
				X: relativeTo.X + int(math.Sqrt(3)*float64(e.hexSideSize)/2.0),
				Y: relativeTo.Y + (3*e.hexSideSize)/2.0,
			}
		case BottomLeft:
			concreteTile = ConcreteHexagonTile{
				X: relativeTo.X - int(math.Sqrt(3)*float64(e.hexSideSize)/2.0),
				Y: relativeTo.Y + (3*e.hexSideSize)/2.0,
			}
		}

		*concreteTiles = append(*concreteTiles, concreteTile)
		e.buildRelativeHexTiles(&directionalTile.relativeHexTile.adjacentTiles, &concreteTile, concreteTiles)
	}
}
