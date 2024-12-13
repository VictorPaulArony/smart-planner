package main

import (
	"log"
	"net/http"
	"os"

	"smart-planner/handlers"
)

// function main of the smart planners system
func main() {
	// functionality to handle the static files
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static", fs))

	// main page handler functions
	http.HandleFunc("/", handlers.IndexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default fallback for local testing
	}
	// localhost for the system testing
	log.Println("http://localhost:3000")
	http.ListenAndServe(":"+port, nil)
}
