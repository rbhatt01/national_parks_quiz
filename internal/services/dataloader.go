package services

import (
	"encoding/json"
	"fmt"
	"national-parks-quiz/internal/models"
	"os"
)

var (
	parks     []models.Park
	questions []models.Question
)

// ParksData represents the structure of parks.json
type ParksData struct {
	Parks []models.Park `json:"parks"`
}

// QuestionsData represents the structure of questions.json
type QuestionsData struct {
	Questions []models.Question `json:"questions"`
}

// LoadData loads both parks and questions from JSON files
func LoadData() error {
	if err := loadParks(); err != nil {
		return fmt.Errorf("failed to load parks: %w", err)
	}
	if err := loadQuestions(); err != nil {
		return fmt.Errorf("failed to load questions: %w", err)
	}
	return nil
}

// loadParks reads and parses parks.json
func loadParks() error {
	data, err := os.ReadFile("data/parks.json")
	if err != nil {
		return fmt.Errorf("error reading parks.json: %w", err)
	}

	var parksData ParksData
	if err := json.Unmarshal(data, &parksData); err != nil {
		return fmt.Errorf("error parsing parks.json: %w", err)
	}

	parks = parksData.Parks
	fmt.Printf("Loaded %d parks\n", len(parks))
	return nil
}

// loadQuestions reads and parses questions.json
func loadQuestions() error {
	data, err := os.ReadFile("data/questions.json")
	if err != nil {
		return fmt.Errorf("error reading questions.json: %w", err)
	}

	var questionsData QuestionsData
	if err := json.Unmarshal(data, &questionsData); err != nil {
		return fmt.Errorf("error parsing questions.json: %w", err)
	}

	questions = questionsData.Questions
	fmt.Printf("Loaded %d questions\n", len(questions))
	return nil
}

// GetParks returns all loaded parks
func GetParks() []models.Park {
	return parks
}

// GetQuestions returns all loaded questions
func GetQuestions() []models.Question {
	return questions
}

// GetQuestionByID returns a specific question by its ID
func GetQuestionByID(id string) (*models.Question, error) {
	for i := range questions {
		if questions[i].ID == id {
			return &questions[i], nil
		}
	}
	return nil, fmt.Errorf("question not found: %s", id)
}

// GetQuestionByIndex returns a question by its index (0-based)
func GetQuestionByIndex(index int) (*models.Question, error) {
	if index < 0 || index >= len(questions) {
		return nil, fmt.Errorf("question index out of range: %d", index)
	}
	return &questions[index], nil
}

// GetTotalQuestions returns the total number of questions
func GetTotalQuestions() int {
	return len(questions)
}
