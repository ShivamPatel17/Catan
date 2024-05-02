package main

// https://www.alexedwards.net/blog/serving-static-sites-with-go
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/catan.html")
	})

	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "healthy2")
	})

	http.HandleFunc("/items", getItemsHandler)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Handler function to handle requests to /items endpoint
func getItemsHandler(w http.ResponseWriter, r *http.Request) {

	// Define a struct to represent your data
	type Item struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	// In-memory storage for the example
	var items = []Item{
		{1, "Item 1"},
		{2, "Item 2"},
		{3, "Item 3"},
	}

	// Convert items to JSON
	jsonData, err := json.Marshal(items)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set content type and send the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
