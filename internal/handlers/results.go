package handlers

import (
	"html/template"
	"net/http"
	"national-parks-quiz/internal/services"
	"strings"
)

// ResultsHandler calculates and displays the quiz result
func ResultsHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Collect all answers from query parameters
		answers := make(map[string]string)
		for key, values := range r.URL.Query() {
			if strings.HasPrefix(key, "answer_") && len(values) > 0 {
				questionID := strings.TrimPrefix(key, "answer_")
				answers[questionID] = values[0]
			}
		}

		// Validate that we have answers
		if len(answers) == 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Calculate the match
		result, err := services.CalculateMatch(answers)
		if err != nil {
			http.Error(w, "Error calculating match: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Prepare template data
		data := map[string]interface{}{
			"Result": result,
		}

		if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}
