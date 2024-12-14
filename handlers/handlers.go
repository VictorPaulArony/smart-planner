package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Template loader
var templates = template.Must(template.New("").ParseFiles(
	"templates/index.html",
	"templates/case.html",
	"templates/contact.html",
	"templates/features.html",
	"template/maps.html",
))

// Handler serves requests for Vercel
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Serve static files
	if r.URL.Path == "/static/" || len(r.URL.Path) > len("/static/") {
		fs := http.FileServer(http.Dir("./static"))
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
		return
	}
	// Render the index page
	renderTemplate(w, "index.html", nil)
}

// renderTemplate is a helper to render templates
func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	err := templates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Println("Template error:", err)
		http.Error(w, "Template rendering error", http.StatusInternalServerError)
	}
}

// function to handle other pages of the website
func FeaturesHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "features.html", nil)
}

func CaseHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "case.html", nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html", nil)
}
func MapsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "maps.html", nil)
}