package main

import (
	"context"
	"gocatan/api"
	"gocatan/board"
	"gocatan/config"
	"gocatan/game"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	mux.HandleFunc("/roll", board.RollHandler)
	mux.HandleFunc("/config", api.GetConfigHandler)

	ctx := context.Background()
	cfg := config.NewConfig()
	game := game.NewGame(ctx, cfg)

	// Register WebSocket handler
	mux.Handle("/ws", websocket.Handler(game.WsHandler))

	// Wrap the ServeMux with the CORS middleware
	handler := enableCors(mux)

	// Start the server
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
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
