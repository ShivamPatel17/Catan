package main

import (
	"context"
	"gocatan/api"
	"gocatan/board"
	"gocatan/board/models"
	"gocatan/config"
	"log"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

// WebSocket handler function

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
type WSMessage struct {
	MessageType string      `json:"messageType"`
	Data        interface{} `json:"data,omitempty"`
}

// WebSocket handler function
func wsHandler(ws *websocket.Conn) {
	defer ws.Close()
	clients[ws] = true // Add new client

	// Send the current game state to the new client
	sendGameState(ws)

	for {
		var message WSMessage
		// Receive a message from the client
		err := websocket.JSON.Receive(ws, &message)
		if err != nil {
			log.Println("Error receiving message:", err)
			delete(clients, ws)
			break
		}

		log.Printf("Received message: %v\n", message)

		// Handle the incoming message based on the message type
		switch message.MessageType {
		case "gameStateRequest":
			// Client requests the current game state
			sendGameState(ws)
		case "action":
			// Example: Handle a player action, update the game state
			handlePlayerAction(message)
		case "vertexClicked":
			deleteVertex(ws)
		default:
			log.Println("Unknown message type:", message.MessageType)
		}
	}
}

// Send the current game state to the client
func sendGameState(ws *websocket.Conn) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	msg := WSMessage{
		MessageType: "gameState",
		Data:        gameState,
	}

	err := websocket.JSON.Send(ws, msg)
	if err != nil {
		log.Println("Error sending game state:", err)
	}
}

// Example function to handle player actions
func handlePlayerAction(message WSMessage) {
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
		err := websocket.JSON.Send(client, WSMessage{
			MessageType: "gameState",
			Data:        gameState,
		})
		if err != nil {
			log.Println("Error sending game state to client:", err)
			client.Close()
			delete(clients, client) // Remove disconnected clients
		}
	}
}
