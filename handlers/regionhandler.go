package handlers

import (
	"html/template"
	"net/http"
)

func RegionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/kisumu-map.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil) // Pass data as the second argument if needed
}
