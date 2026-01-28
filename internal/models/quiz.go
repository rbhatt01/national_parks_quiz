package models

// QuizState tracks the current state of a user's quiz session
type QuizState struct {
	CurrentQuestion int               // 1-based question number
	Answers         map[string]string // questionID -> optionID (e.g., "Q1" -> "A")
}

// QuizResult represents the outcome of the quiz
type QuizResult struct {
	Park           *Park
	UserTraits     map[string]float64 // Normalized user trait scores
	MatchScore     float64            // How closely the park matched (lower is better for distance-based)
	TopTraits      []string           // The traits that contributed most to the match
}
