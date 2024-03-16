package main

import (
	"log"
	"net/http"

	"groupie-tracker/handlers"
	"groupie-tracker/render"
)

func main() {
	renderer, err := render.NewTemplateReader("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files (CSS, JS, etc.)
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	handlers.SetRenderer(renderer)

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/location/{id}", handlers.LocationHandler)

	// http.HandleFunc("/search", handlers.SearchHandler)

	// Start the server
	log.Println("Server is starting on port 8080...")
	http.ListenAndServe(":8080", mux)
}
