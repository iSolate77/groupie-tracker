package main

import (
	"net/http"

	"groupie-tracker/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.ListenAndServe(":8080", nil)
}
