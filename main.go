package main

import (
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.ListenAndServe(":8080", nil)
}
