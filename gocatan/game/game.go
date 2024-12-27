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
		fmt.Printf("error parsing uuid in the delete Vertex func")
	}
	fmt.Printf("%d Vertices remaining", len(g.board.Vertices))
	delete(g.board.Vertices, u)

	fmt.Printf("in the delete !!! JVertex func with ws:%s\n", u)
}
