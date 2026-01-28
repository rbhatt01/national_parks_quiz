package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"national-parks-quiz/internal/handlers"
	"national-parks-quiz/internal/middleware"
	"national-parks-quiz/internal/services"
)

func main() {
	// Load data at startup
	log.Println("Loading quiz data...")
	if err := services.LoadData(); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	// Parse templates separately for each handler to avoid namespace collisions
	log.Println("Parsing templates...")
	funcMap := template.FuncMap{
		"mult": func(a, b float64) float64 {
			return a * b
		},
	}

	homeTmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles(
		"templates/base.html",
		"templates/home.html",
	)
	if err != nil {
		log.Fatalf("Failed to parse home templates: %v", err)
	}

	quizTmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles(
		"templates/base.html",
		"templates/quiz.html",
	)
	if err != nil {
		log.Fatalf("Failed to parse quiz templates: %v", err)
	}

	resultsTmpl, err := template.New("base.html").Funcs(funcMap).ParseFiles(
		"templates/base.html",
		"templates/results.html",
	)
	if err != nil {
		log.Fatalf("Failed to parse results templates: %v", err)
	}

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler(homeTmpl))
	mux.HandleFunc("/quiz", handlers.QuizHandler(quizTmpl))
	mux.HandleFunc("/results", handlers.ResultsHandler(resultsTmpl))

	// Wrap with middleware
	loggedMux := middleware.Logging(mux)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, loggedMux))
}
