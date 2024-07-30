package handlers

import (
	"html/template"
	"net/http"
)

// type AsciiData struct {
// 	Text   string
// 	Banner string
// 	Result string
// }

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Render the landing page
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// data := AsciiData{}
	err = tmpl.Execute(w, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
