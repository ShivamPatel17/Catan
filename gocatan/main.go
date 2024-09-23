package main

import (
	"context"
	"encoding/json"
	"gocatan/api"
	"gocatan/board"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/websocket"
)

// WSMessage represents a message structure for WebSocket communication
type WSMessage struct {
	MessageType string `json:"messageType"`
	Content     string `json:"content,omitempty"`
}

// WebSocket handler function
func wsHandler(ws *websocket.Conn) {
	defer ws.Close()
	log.Println("New WebSocket connection established")

	for {
		var message string
		// Receive a message from the client
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			log.Println("Error receiving message:", err)
			break
		}

		log.Printf("Received message: %s\n", message)

		// Decode the received message
		var receivedMsg WSMessage
		err = json.Unmarshal([]byte(message), &receivedMsg)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		// Handle message types
		switch receivedMsg.MessageType {
		case "rollDice":
			// Perform dice roll or relevant game action here
			log.Println("Roll dice action received")
		default:
			log.Println("Unknown message type:", receivedMsg.MessageType)
		}

		// Create response message
		responseMsg := WSMessage{MessageType: "gameState"}

		// Marshal the response into JSON
		msg, err := json.Marshal(responseMsg)
		if err != nil {
			log.Println("Error marshalling message:", err)
			continue
		}

		// Send the message back to the client
		err = websocket.Message.Send(ws, msg)
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
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

	// Create an HTTP server with timeout settings and graceful shutdown support
	server := &http.Server{
		Addr:    ":3000",
		Handler: handler,
	}

	// Start the server in a goroutine
	go func() {
		log.Println("Server is listening on :3000...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// Graceful shutdown setup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Block until we receive a signal
	<-stop

	// Gracefully shut down the server, waiting 5 seconds for active connections to close
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
