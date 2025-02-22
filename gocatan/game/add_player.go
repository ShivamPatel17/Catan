package game

import (
	messages "gocatan/api/wsmessages"
	"gocatan/board/models"

	"golang.org/x/net/websocket"
)

// this should be called by the "lobby" at some point
// for now, it's called by the websocket handler
func (g *Game) AddPlayer(m messages.BaseMessage, c *websocket.Conn) {
	g.pwc[m.PlayerUuid] = PlayerWithConnection{
		player: &models.Player{},
		conn:   c,
	}
}
