package handlers

import (
	"html/template"
	"net/http"
)

// HomeHandler renders the landing page
func HomeHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "base.html", nil); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}
