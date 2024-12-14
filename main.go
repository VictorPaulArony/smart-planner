package main

import (
	"log"
	"net/http"
	"os"

	"smart-planner/handlers"
)

// function main of the smart planners system
func main() {
	// The output of the map data.
	// osmFile := "kisumu-data.xml"

	// functionality to handle the static files
	http.HandleFunc("/", handlers.IndexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default fallback for local testing
	}
	// localhost for the system testing
	log.Println("http://localhost:3000")
	http.ListenAndServe(":"+port, nil)
}
