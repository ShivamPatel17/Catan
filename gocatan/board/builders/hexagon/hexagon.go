package hexagon

import (
	"fmt"
	"gocatan/board/models"
	mathhelper "gocatan/internal/math"
	"math/rand"

	"github.com/google/uuid"
)

func (e *HexagonEngine) BuildHexagons(relativeTiles *models.RelativeHexagonTile) ([]*models.ConcreteHexagonTile, error) {
	concreteTiles := make([]*models.ConcreteHexagonTile, 0)
	ConcreteTile := relativeTiles.Concrete
	if ConcreteTile != nil {
		concreteTiles = append(concreteTiles, &models.ConcreteHexagonTile{
			X: ConcreteTile.X,
			Y: ConcreteTile.Y,
		})
	}
	e.buildRelativeHexTiles(relativeTiles.AdjacentTiles, concreteTiles[0], &concreteTiles)
	assignResources(concreteTiles)
	return concreteTiles, nil
}

func (e *HexagonEngine) buildRelativeHexTiles(tiles []*models.DirectionalHexagonTile, relativeTo *models.ConcreteHexagonTile, concreteTiles *[]*models.ConcreteHexagonTile) {
	if relativeTo == nil {
		return
	}

	for _, directionalTile := range tiles {
		var ConcreteTile models.ConcreteHexagonTile
		switch directionalTile.Direction {
		case models.TopRight:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X + mathhelper.HeightOfEqualateralTriangle(e.HexSideSize),
				Y: relativeTo.Y - ((3 * e.HexSideSize) / 2.0),
			}
		case models.TopLeft:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X - mathhelper.HeightOfEqualateralTriangle(e.HexSideSize),
				Y: relativeTo.Y - ((3 * e.HexSideSize) / 2.0),
			}
		case models.MiddleRight:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X + (e.HexTotalWidth),
				Y: relativeTo.Y,
			}
		case models.MiddleLeft:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X - (e.HexTotalWidth),
				Y: relativeTo.Y,
			}
		case models.BottomRight:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X + mathhelper.HeightOfEqualateralTriangle(e.HexSideSize),
				Y: relativeTo.Y + ((3 * e.HexSideSize) / 2.0),
			}
		case models.BottomLeft:
			ConcreteTile = models.ConcreteHexagonTile{
				X: relativeTo.X - mathhelper.HeightOfEqualateralTriangle(e.HexSideSize),
				Y: relativeTo.Y + ((3 * e.HexSideSize) / 2.0),
			}
		}

		*concreteTiles = append(*concreteTiles, &ConcreteTile)
		e.buildRelativeHexTiles(directionalTile.RelativeHexTile.AdjacentTiles, &ConcreteTile, concreteTiles)
	}
}

func assignResources(concreteTiles []*models.ConcreteHexagonTile) {
	for _, tile := range concreteTiles {
		u, _ := uuid.NewV7()
		tile.Uuid = u
		tile.Resource = getRandomResource()
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
