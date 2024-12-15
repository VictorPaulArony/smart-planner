package main

import (
	"log"
	"net/http"
	"os"

	"smart-planner/handlers"
)

// function main of the smart planners system
func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/features", handlers.FeaturesHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	http.HandleFunc("/case", handlers.CaseHandler)
	http.HandleFunc("/maps", handlers.MapsHandler)

	// seting the environment for easy deployment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default fallback for local testing
	}
	// localhost for the system testing
	log.Println("http://localhost:3000")
	http.ListenAndServe(":"+port, nil)
}
