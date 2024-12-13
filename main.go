package main

import (
	"log"
	"net/http"

	"smart-planner/handlers"
)

// function main of the smart planners system
func main() {
	// functionality to handle the static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", http.StripPrefix("/static/", fs))

	// main page handler functions
	http.HandleFunc("/", handlers.IndexHandler)

	// localhost for the system testing
	log.Println("http://localhost:1234")
	http.ListenAndServe(":1234", nil)
}
