package main

import (
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	log.Println("Server is starting on port 8080...")
	http.HandleFunc("/", handlers.IndexHandler)
	// http.HandleFunc("/search", handlers.SearchHandler)

	// Serve static files (CSS, JS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	http.ListenAndServe(":8080", nil)
}
