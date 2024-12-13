package main

import (
	"fmt"
	"log"
	"net/http"

	mapping "smart-planner/mapping"

	"smart-planner/handlers"
)

// function main of the smart planners system
func main() {
	// The output of the map data.
	osmFile := "kisumu-data.xml"
	osm, err := mapping.ParseOSMFile(osmFile)
	if err != nil {
		fmt.Println("Error parsing OSM file:", err)
		return
	}
	// "This is the raw data"
	fmt.Println("This is my osm data\n", "<<<", osm)

	// Print the ralation details
	for _, relation := range osm.Relations {
		fmt.Printf("Relation ID: %d, Visible: %v, Version: %d, User: %s\n",
			relation.ID, relation.Visible, relation.Version, relation.User)
		fmt.Println("Members:")
		for _, member := range relation.Members {
			fmt.Printf(" - Type: %s, Ref: %d, Role: %s\n", member.Type, member.Ref, member.Role)
		}
		for _, tag := range relation.Tags {
			fmt.Printf(" - %s = %s\n", tag.Key, tag.Value)
		}
	}

	// functionality to handle the static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", http.StripPrefix("/static/", fs))

	// main page handler functions
	http.HandleFunc("/", handlers.IndexHandler)

	// localhost for the system testing
	log.Println("http://localhost:1234")
	http.ListenAndServe(":1234", nil)
}
