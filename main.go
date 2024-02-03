package main

// https://www.alexedwards.net/blog/serving-static-sites-with-go
import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/catan.html")
	})

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
