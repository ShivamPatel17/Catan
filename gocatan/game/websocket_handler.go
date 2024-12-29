package game

import (
	"encoding/json"
	"fmt"
	types "gocatan/api/wsmessages"
	"golang.org/x/net/websocket"
	"log"
)

func (g *Game) WsHandler(ws *websocket.Conn) {
	defer ws.Close()
	g.AddClient(ws)
	g.BroadcastGameState()

	for {
		// Step 1: Read message from WebSocket
		var msg []byte
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			g.BroadcastGameState()
			log.Println("Error reading message:", err)
			break
		}

		// Step 2: Parse the message
		parsedMessage, err := parseWebSocketMessage(msg)
		if err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		switch m := parsedMessage.(type) {
		case types.BuildSettlementMessage:
			g.BuildSettlement(m)
			g.BroadcastGameState()
		case types.VertexClickedMessage:
			g.DeleteVertex(m)
			g.BroadcastGameState()
		default:
			log.Println("Unknown message type")
		}
	}
}

func parseWebSocketMessage(data []byte) (interface{}, error) {
	// Step 1: Unmarshal into BaseMessage to extract the MessageType
	var baseMsg types.BaseMessage
	if err := json.Unmarshal(data, &baseMsg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base message: %w", err)
	}

	// Step 2: Unmarshal into the appropriate specific struct
	switch baseMsg.MessageType {
	case "gameState":
		var msg types.GameStateMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal gameState message: %w", err)
		}

		return msg, nil
	case "buildSettlement":
		var msg types.BuildSettlementMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal buildSettlement message: %w", err)
		}
		return msg, nil
	case "vertexClicked":
		var msg types.VertexClickedMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal vertexClicked message: %w", err)
		}
		return msg, nil
	default:
		return nil, fmt.Errorf("unknown message type: %s", baseMsg.MessageType)
	}
}
