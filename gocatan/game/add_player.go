package game

import (
	messages "gocatan/api/wsmessages"
	"gocatan/board/models"
	"golang.org/x/net/websocket"
)

// this should be called by the "lobby" at some point
// for now, it's called by the websocket handler
func (g *Game) AddPlayer(m messages.BaseMessage, c *websocket.Conn) {
	g.SendGameStateToConnection(c)
	g.pwc[m.GetPlayerUUID()] = PlayerWithConnection{
		player: &models.Player{
			Uuid: m.GetPlayerUUID(),
		},
		conn: c,
	}
}
