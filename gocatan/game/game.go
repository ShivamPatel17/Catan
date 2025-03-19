package game

import (
	"context"
	"fmt"
	"gocatan/api"
	"gocatan/api/wsmessages"
	"gocatan/board/models"
	"gocatan/config"
	"log"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type PlayerWithConnection struct {
	player *models.Player
	conn   *websocket.Conn
}

type Game struct {
	board models.GameBoard
	mutex sync.Mutex
	// clients can be anyone.. consider people "spectating"
	clients map[*websocket.Conn]bool
	// for now, any new client will be considered a new player. Need the front end to have some way of identifying itself to limit this map
	pwc map[uuid.UUID]PlayerWithConnection
}

func NewGame(ctx context.Context, cfg config.Config) *Game {
	return &Game{
		board:   api.BuildBoard(ctx, cfg),
		clients: make(map[*websocket.Conn]bool),
		pwc:     make(map[uuid.UUID]PlayerWithConnection),
	}
}

// Broadcast updated game state to all connected clients
func (g *Game) BroadcastGameState() {
	fmt.Println("BroadcastGameState()")
	g.mutex.Lock()
	defer g.mutex.Unlock()
	fmt.Printf("number of clients is %d\n", len(g.pwc))

	for _, p := range g.pwc {
		err := g.SendGameStateToConnection(&p)
		if err != nil {
			log.Println(err)
			delete(g.pwc, p.player.Uuid) // remove disconnected clients
			p.conn.Close()
		}
	}
}

// TODO: refactor to accept the map instead of relying on the internals of the game
func (g *Game) SendGameStateToConnection(p *PlayerWithConnection) error {
	err := websocket.JSON.Send(p.conn, messages.GameStateMessage{
		EmbeddedBaseMessage: messages.EmbeddedBaseMessage{
			MessageType: "gameState",
			PlayerUuid:  p.player.Uuid,
		},
		Board: g.board,
	})
	if err != nil {
		return fmt.Errorf("error sending game state to client: %w", err)
	}
	return nil
}

func (g *Game) DeleteVertex(v messages.VertexClickedMessage) {
	u, err := uuid.Parse(v.Data.Id)
	if err != nil {
		log.Printf("error parsing uuid in the delete Vertex func")
	}

	delete(g.board.Vertices, u)
}
