package models

// Question represents a quiz question with multiple options
type Question struct {
	ID      string   `json:"id"`
	Prompt  string   `json:"prompt"`
	Options []Option `json:"options"`
}

// Option represents an answer choice with trait scores
type Option struct {
	ID     string            `json:"id"`     // e.g., "A", "B", "C"
	Text   string            `json:"text"`
	Scores map[string]float64 `json:"scores"` // trait name -> score (0-5)
}
