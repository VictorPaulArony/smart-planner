package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"smart-planner/handlers"
	mapping "smart-planner/mapping"
)

// function main of the smart planners system
func main() {
	// The output of the map data.
	filename := "kisumu-data.xml"
	osmData, err := mapping.ParseOSMFile(filename)
	if err != nil {
		log.Println("Error parsing osm file", err)
		return
	}

	// Convert to GeoJSON
	geoJSON := mapping.ConvertOSMToGeoJSON(osmData)

	// Save GeoJSON to file
	file, err := os.Create("kisumu-data.geojson")
	if err != nil {
		fmt.Println("Error creating GeoJSON file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(geoJSON)
	if err != nil {
		fmt.Println("Error writing GeoJSON file:", err)
		return
	}

	fmt.Println("GeoJSON data saved to kisumu-data.geojson")

	//This handler serves the data from kisumu.
	http.HandleFunc("/map/kisumu-map", handlers.RegionHandler)


	// Start the server
	port := ":8080"
	fmt.Printf("Server running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
