package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
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
	mux.HandleFunc("/roll", rollHandler)

	// Wrap the ServeMux with the CORS middleware
	handler := enableCors(mux)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		log.Fatal(err)
	}
}

// Handler function to handle requests to /items endpoint
func rollHandler(w http.ResponseWriter, r *http.Request) {
	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 6
	roll := rand.Intn(6) + 1

	// Convert items to JSON
	jsonData, err := json.Marshal(roll)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set content type and send the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
