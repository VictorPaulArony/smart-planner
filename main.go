package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"smart-planner/handlers"
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

	// Send request
	resp, err := http.Post(OVERPASSURL, "text/plain", bytes.NewBufferString(query))
	if err != nil {
		log.Fatalf("Error making request to Overpass API: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Parse JSON response
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error parsing JSON response: %v", err)
	}

	// Print response or process further
	log.Printf("Overpass response: %s", body)

	// This handler serves the data from kisumu.
	http.HandleFunc("/map/kisumu-map", handlers.RegionHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("Server running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
