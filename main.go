package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"smart-planner/handlers"
	"smart-planner/mapping"
)

const OVERPASSURL = "https://overpass-api.de/api/interpreter"

// function main of the smart planners system
func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/features", handlers.FeaturesHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	http.HandleFunc("/case", handlers.CaseHandler)
	http.HandleFunc("/maps", handlers.MapsHandler)
	// Overpass API query
	query := `
    [out:json];
    area[name="Kisumu"]->.searchArea;
    (
      relation["boundary"="administrative"](area.searchArea);
      way["boundary"="administrative"](area.searchArea);
    );
    out body;
    >;
    out skel qt;
    `
	data, err := mapping.KisumuMap(OVERPASSURL, query)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(">>>> ", data)

	// Serve the HTML template for the map
	http.Handle("/geojson/", http.StripPrefix("/geojson", http.FileServer(http.Dir("."))))
	// server the HML files from the templates folder
	http.Handle("/kisumu", http.FileServer(http.Dir("templates")))

	// This handler serves the data from kisumu.
	http.HandleFunc("/kisumu-map", handlers.RegionHandler)

	// seting the environment for easy deployment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default fallback for local testing
	}
	// localhost for the system testing
	log.Println("http://localhost:3000")
	http.ListenAndServe(":"+port, nil)
}
