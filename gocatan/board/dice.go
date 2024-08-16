package board

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

// Handler function to handle requests to /items endpoint
func RollHandler(w http.ResponseWriter, r *http.Request) {
	// Seed the random number generator to ensure different results each run
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
