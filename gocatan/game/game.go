package game

import (
	"context"
	"fmt"
	"gocatan/api"
	types "gocatan/api/wsmessages"
	"gocatan/board/models"
	"gocatan/config"
	"log"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type Game struct {
	board   models.GameBoard
	mutex   sync.Mutex
	clients map[*websocket.Conn]bool
}

func NewGame(ctx context.Context, cfg config.Config) *Game {
	return &Game{
		board:   api.BuildBoard(ctx, cfg),
		clients: make(map[*websocket.Conn]bool),
	}
}

func (g *Game) AddClient(ws *websocket.Conn) {
	g.clients[ws] = true
}

// Broadcast updated game state to all connected clients
func (g *Game) BroadcastGameState() {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	fmt.Printf("number of clients is %d\n", len(g.clients))

	for ws := range g.clients {
		fmt.Println(ws)
	}
	fmt.Println()

	for client := range g.clients {
		err := websocket.JSON.Send(client, types.GameStateMessage{
			BaseMessage: types.BaseMessage{
				MessageType: "gameState",
			},
			Board: g.board,
		})
		if err != nil {
			log.Println("Error sending game state to client:", err)
			client.Close()
			delete(g.clients, client) // Remove disconnected clients
		}
	}
}

func (g *Game) DeleteVertex(v types.VertexClickedMessage) {
	u, err := uuid.Parse(v.Data.Id)
	if err != nil {
		log.Printf("error parsing uuid in the delete Vertex func")
	}

	delete(g.board.Vertices, u)
}

func (g *Game) BuildSettlement(b types.BuildSettlementMessage) error {
	vertex, ok := g.board.Vertices[b.Data.VertexUuid]
	if !ok {
		return fmt.Errorf("invalid Vertex provided to build a settlement")
	}

	vertex.Building = models.Settlement
	return nil
}
