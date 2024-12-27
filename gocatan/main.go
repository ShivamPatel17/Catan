package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gocatan/api"
	"gocatan/board"
	"gocatan/board/models"
	"gocatan/config"
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

func main() {
	// Static file server for serving files from the ./static directory
	fs := http.FileServer(http.Dir("./static"))

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register handlers with the ServeMux
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/roll", board.RollHandler)
	mux.HandleFunc("/board", api.BoardHandler)
	mux.HandleFunc("/config", api.GetConfigHandler)

	// Register WebSocket handler
	mux.Handle("/ws", websocket.Handler(wsHandler))

	// Wrap the ServeMux with the CORS middleware
	handler := enableCors(mux)

	cfg := config.NewConfig()
	ctx := context.Background()
	gameState = api.BuildBoard(ctx, cfg)

	// Start the server
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Define the persistent game state structure
var gameState models.GameBoard

var clients = make(map[*websocket.Conn]bool) // Track all connected clients
var stateMutex sync.Mutex                    // Mutex to handle concurrent access

// WSMessage represents a WebSocket message
type BaseMessage struct {
	MessageType string `json:"messageType"`
}

type VertexClickedMessage struct {
	BaseMessage
	Data VertexClickedMessageData `json:"Data"`
}

type VertexClickedMessageData struct {
	Id string `json:"Id"`
}

type GameStateMessage struct {
	BaseMessage
	Data interface{} `json:"data"`
}

func wsHandler(ws *websocket.Conn) {
	defer ws.Close()

	for {
		// Step 1: Read message from WebSocket
		var msg []byte
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Step 2: Parse the message
		parsedMessage, err := parseWebSocketMessage(msg)
		if err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		// Step 3: Handle the parsed message
		switch m := parsedMessage.(type) {
		case GameStateMessage:
			log.Printf("Game State: %s\n", m.Data)
			sendGameState(ws)
		case VertexClickedMessage:
			deleteVertex(m)
		default:
			log.Println("Unknown message type")
		}
	}
}

func parseWebSocketMessage(data []byte) (interface{}, error) {
	// Step 1: Unmarshal into BaseMessage to extract the MessageType
	var baseMsg BaseMessage
	if err := json.Unmarshal(data, &baseMsg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base message: %w", err)
	}

	// Step 2: Unmarshal into the appropriate specific struct
	switch baseMsg.MessageType {
	case "gameState":
		var msg GameStateMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal gameState message: %w", err)
		}

		return msg, nil
	case "vertexClicked":
		var msg VertexClickedMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal gameState message: %w", err)
		}
		return msg, nil
	default:
		return nil, fmt.Errorf("unknown message type: %s", baseMsg.MessageType)
	}
}

// Send the current game state to the client
func sendGameState(ws *websocket.Conn) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	msg := GameStateMessage{
		BaseMessage: BaseMessage{
			MessageType: "gameState",
		},
		Data: gameState,
	}

	err := websocket.JSON.Send(ws, msg)
	if err != nil {
		log.Println("Error sending game state:", err)
	}
}

// Example function to handle player actions
func handlePlayerAction(message BaseMessage) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	// Parse the message and update the game state
	// For example, modifying the vertices or edges based on player action
	log.Printf("Handling player action: %v\n", message)

	// After updating the game state, broadcast the updated state to all clients
	broadcastGameState()
}

// Broadcast updated game state to all connected clients
func broadcastGameState() {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	for client := range clients {
		err := websocket.JSON.Send(client, GameStateMessage{
			BaseMessage: BaseMessage{
				MessageType: "gameState",
			},
			Data: gameState,
		})
		if err != nil {
			log.Println("Error sending game state to client:", err)
			client.Close()
			delete(clients, client) // Remove disconnected clients
		}
	}
}

func deleteVertex(v VertexClickedMessage) {
	fmt.Printf("in the deleteVertex func with ws:%s\n", v.Data.Id)
}

// CORS middleware function
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Handle WebSocket request: Upgrade request method is "GET"
		if r.Header.Get("Upgrade") == "websocket" {
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
