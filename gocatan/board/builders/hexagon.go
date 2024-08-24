package builders

import (
	"gocatan/board/models"
	"gocatan/config"
	"math"
	"math/rand"
)

type HexagonEngine struct {
	HexSideSize    float32
	HexTotalWidth  float32
	HexTotalHeight float32
}

func NewHexagonEngine(cfg config.Config) HexagonEngine {
	return HexagonEngine{
		HexSideSize:    cfg.HexHeight / 2,
		HexTotalWidth:  cfg.HexWidth,
		HexTotalHeight: cfg.HexHeight,
	}
}

func (e *HexagonEngine) BuildHexagons(relativeTiles *models.RelativeHexagonTile) ([](models.ConcreteHexagonTile), error) {
	concreteTiles := make([]models.ConcreteHexagonTile, 0)
	ConcreteTile := relativeTiles.Concrete
	if ConcreteTile != nil {
		concreteTiles = append(concreteTiles, models.ConcreteHexagonTile{
			X: ConcreteTile.X,
			Y: ConcreteTile.Y,
		})
	}
	e.buildRelativeHexTiles(&relativeTiles.AdjacentTiles, &concreteTiles[0], &concreteTiles)
	assignResources(&concreteTiles)
	return concreteTiles, nil
}

func (e *HexagonEngine) buildRelativeHexTiles(tiles *[]models.DirectionalHexagonTile, relativeTo *models.ConcreteHexagonTile, concreteTiles *[]models.ConcreteHexagonTile) {
	if relativeTo == nil {
		return
	}

	for _, directionalTile := range *tiles {
		var ConcreteTile models.ConcreteHexagonTile
		switch directionalTile.Direction {
		case models.TopRight:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X + int(math.Sqrt(3)*float64(e.HexSideSize)/2.0),
				Y: relativeTo.Y - int((3*e.HexSideSize)/2.0),
			}
		case models.TopLeft:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X - int(math.Sqrt(3)*float64(e.HexSideSize)/2.0),
				Y: relativeTo.Y - int((3*e.HexSideSize)/2.0),
			}
		case models.MiddleRight:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X + int(e.HexTotalWidth),
				Y: relativeTo.Y,
			}
		case models.MiddleLeft:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X - int(e.HexTotalWidth),
				Y: relativeTo.Y,
			}
		case models.BottomRight:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X + int(math.Sqrt(3)*float64(e.HexSideSize)/2.0),
				Y: relativeTo.Y + int((3*e.HexSideSize)/2.0),
			}
		case models.BottomLeft:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X - int(math.Sqrt(3)*float64(e.HexSideSize)/2.0),
				Y: relativeTo.Y + int((3*e.HexSideSize)/2.0),
			}
		}

		*concreteTiles = append(*concreteTiles, ConcreteTile)
		e.buildRelativeHexTiles(&directionalTile.RelativeHexTile.AdjacentTiles, &ConcreteTile, concreteTiles)
	}
}

func assignResources(concreteTiles *[]models.ConcreteHexagonTile) {
	for i := range *concreteTiles {
		(*concreteTiles)[i].Resource = getRandomResource()
	}
}
func getRandomResource() models.Resource {
	// Seed the random number generator to ensure different results on each run

	// Create a slice of all possible resources
	resources := []models.Resource{models.Sheep, models.Wheat, models.Ore, models.Wood, models.Brick}

	// Select a random index
	randomIndex := rand.Intn(len(resources))

	// Return the randomly selected resource
	return resources[randomIndex]
}
