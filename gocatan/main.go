package main

import (
	"gocatan/api"
	"gocatan/board"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

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

		// Echo the message back to the client
		err = websocket.Message.Send(ws, message)
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

	// Start the server
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
