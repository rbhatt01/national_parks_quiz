package handlers

import (
	"html/template"
	"net/http"
	"national-parks-quiz/internal/services"
	"strconv"
	"strings"
)

// QuizHandler handles both GET (display question) and POST (submit answer)
func QuizHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handleQuizGet(w, r, tmpl)
		} else if r.Method == http.MethodPost {
			handleQuizPost(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// handleQuizGet displays a specific question
func handleQuizGet(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	// Get question number from query parameter
	questionNumStr := r.URL.Query().Get("question")
	questionNum, err := strconv.Atoi(questionNumStr)
	if err != nil || questionNum < 1 {
		questionNum = 1
	}

	totalQuestions := services.GetTotalQuestions()
	if questionNum > totalQuestions {
		http.Redirect(w, r, "/quiz?question="+strconv.Itoa(totalQuestions), http.StatusSeeOther)
		return
	}

	// Get the question (0-indexed)
	question, err := services.GetQuestionByIndex(questionNum - 1)
	if err != nil {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	// Collect previous answers from query parameters
	previousAnswers := make(map[string]string)
	for key, values := range r.URL.Query() {
		if strings.HasPrefix(key, "prev_") && len(values) > 0 {
			questionID := strings.TrimPrefix(key, "prev_")
			previousAnswers[questionID] = values[0]
		}
	}

	// Build previous answers excluding the last one (for back button)
	previousAnswersExceptLast := make(map[string]string)
	if questionNum > 1 {
		lastQuestionID := "Q" + strconv.Itoa(questionNum-1)
		for k, v := range previousAnswers {
			if k != lastQuestionID {
				previousAnswersExceptLast[k] = v
			}
		}
	}

	data := map[string]interface{}{
		"Question":                   question,
		"CurrentQuestion":            questionNum,
		"TotalQuestions":             totalQuestions,
		"ProgressPercent":            (questionNum * 100) / totalQuestions,
		"IsLastQuestion":             questionNum == totalQuestions,
		"PreviousAnswers":            previousAnswers,
		"PreviousAnswersExceptLast": previousAnswersExceptLast,
		"PrevQuestion":               questionNum - 1,
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

// handleQuizPost processes the answer and redirects to next question or results
func handleQuizPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get current question number
	currentQuestionStr := r.FormValue("current_question")
	currentQuestion, err := strconv.Atoi(currentQuestionStr)
	if err != nil {
		http.Error(w, "Invalid question number", http.StatusBadRequest)
		return
	}

	// Get the selected answer
	answer := r.FormValue("answer")
	if answer == "" {
		http.Error(w, "No answer selected", http.StatusBadRequest)
		return
	}

	// Collect all previous answers
	allAnswers := make(map[string]string)
	for key, values := range r.Form {
		if strings.HasPrefix(key, "prev_") && len(values) > 0 {
			questionID := strings.TrimPrefix(key, "prev_")
			allAnswers[questionID] = values[0]
		}
	}

	// Add current answer
	currentQuestionID := "Q" + strconv.Itoa(currentQuestion)
	allAnswers[currentQuestionID] = answer

	totalQuestions := services.GetTotalQuestions()

	// If this is the last question, redirect to results
	if currentQuestion >= totalQuestions {
		// Build form data for POST to results
		redirectURL := "/results?"
		params := []string{}
		for qid, ans := range allAnswers {
			params = append(params, "answer_"+qid+"="+ans)
		}
		redirectURL += strings.Join(params, "&")
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
		return
	}

	// Otherwise, redirect to next question with all answers as query params
	nextQuestion := currentQuestion + 1
	redirectURL := "/quiz?question=" + strconv.Itoa(nextQuestion)
	for qid, ans := range allAnswers {
		redirectURL += "&prev_" + qid + "=" + ans
	}

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
