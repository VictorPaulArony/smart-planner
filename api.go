package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Location struct {
	Coords  [2]float64 `json:"coords"`
	Title   string     `json:"title"`
	Density int        `json:"density"`
}

type Resource struct {
	Schools   []Location `json:"schools"`
	Hospitals []Location `json:"hospitals"`
}

type (
	Ward      map[string]Resource
	SubCounty map[string]Ward
	County    map[string]SubCounty
)

var data County

// Middleware to handle CORS
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from frontend origin (adjust as needed)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Handler for the /api/data endpoint
func handleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Load JSON data from the file
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("Failed to open data.json: %v", err)
	}
	defer file.Close()

	// Decode JSON data into the `data` variable
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatalf("Failed to decode data.json: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/data", handleData)

	// Wrap the mux with the CORS middleware
	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", enableCors(mux))
}
