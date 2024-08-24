package main

import (
	"gocatan/api"
	board "gocatan/board"
	"gocatan/config"
	"log"
	"net/http"
)

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

		next.ServeHTTP(w, r)
	})
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register handlers with the ServeMux
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/roll", board.RollHandler)
	mux.HandleFunc("/board", api.BoardHandler)
	mux.HandleFunc("/config", config.GetConfigHandler)

	// Wrap the ServeMux with the CORS middleware
	handler := enableCors(mux)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		log.Fatal(err)
	}
}
