package main

import (
	"fmt"
	"log"
	"net/http"

	"smart-planner/handlers"
	"smart-planner/mapping"
)

const OVERPASSURL = "https://overpass-api.de/api/interpreter"

// function main of the smart planners system
func main() {
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

	// Start the server
	port := ":8080"
	fmt.Printf("Server running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
