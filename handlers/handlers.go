package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// function to handle the index page of the system
// IndexHandler serves the signup page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

// renderTemplate is a helper to render templates
func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	err := templates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Println("Template error:", err)
	}
}

// Template loader
var templates = template.Must(template.New("").ParseFiles(

	"templates/index.html",
))
