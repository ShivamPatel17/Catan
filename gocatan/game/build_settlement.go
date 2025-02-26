package game

import (
	"gocatan/api/wsmessages"
	"gocatan/board/models"

	"fmt"
)

func (g *Game) BuildSettlement(b messages.BuildSettlementMessage) error {
	vertex, ok := g.board.Vertices[b.Data.VertexUuid]
	if !ok {
		return fmt.Errorf("invalid Vertex provided to build a settlement")
	}

	vertex.Building = models.Settlement
	return nil
}
