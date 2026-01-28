package services

import (
	"fmt"
	"math"
	"national-parks-quiz/internal/models"
	"sort"
)

// CalculateMatch finds the best matching park based on user's answers
func CalculateMatch(answers map[string]string) (*models.QuizResult, error) {
	// Get all parks
	parks := GetParks()

	if len(parks) == 0 {
		return nil, fmt.Errorf("no parks loaded")
	}

	// Step 1: Accumulate trait scores from user's answers
	traitScores := make(map[string]float64)
	traitCounts := make(map[string]int) // Track how many questions contributed to each trait

	for questionID, optionID := range answers {
		question, err := GetQuestionByID(questionID)
		if err != nil {
			continue // Skip invalid questions
		}

		// Find the selected option
		var selectedOption *models.Option
		for i := range question.Options {
			if question.Options[i].ID == optionID {
				selectedOption = &question.Options[i]
				break
			}
		}

		if selectedOption == nil {
			continue // Skip invalid options
		}

		// Add scores from this option
		for trait, score := range selectedOption.Scores {
			traitScores[trait] += score
			traitCounts[trait]++
		}
	}

	// Step 2: Normalize user trait scores to 0-1 scale
	// Assuming max score per question is 5, and each trait appears in multiple questions
	userTraits := make(map[string]float64)
	for trait, totalScore := range traitScores {
		count := traitCounts[trait]
		if count > 0 {
			// Average score normalized to 0-1 scale
			avgScore := totalScore / float64(count)
			// Scores can range from 0 to 5, so normalize to [0,1]
			userTraits[trait] = avgScore / 5.0
		}
	}

	// Step 3: Find the park with the smallest distance to user's trait profile
	var bestPark *models.Park
	bestDistance := math.MaxFloat64

	for i := range parks {
		park := &parks[i]
		distance := calculateTraitDistance(userTraits, park.Traits)

		if distance < bestDistance {
			bestDistance = distance
			bestPark = park
		}
	}

	if bestPark == nil {
		return nil, fmt.Errorf("no matching park found")
	}

	// Step 4: Identify top contributing traits
	topTraits := getTopTraits(userTraits, bestPark.Traits, 3)

	return &models.QuizResult{
		Park:       bestPark,
		UserTraits: userTraits,
		MatchScore: bestDistance,
		TopTraits:  topTraits,
	}, nil
}

// calculateTraitDistance calculates Euclidean distance between user and park trait profiles
func calculateTraitDistance(userTraits, parkTraits map[string]float64) float64 {
	var sumSquares float64
	allTraits := make(map[string]bool)

	// Collect all unique traits
	for trait := range userTraits {
		allTraits[trait] = true
	}
	for trait := range parkTraits {
		allTraits[trait] = true
	}

	// Calculate distance for each trait
	for trait := range allTraits {
		userValue := userTraits[trait] // Defaults to 0 if not present
		parkValue := parkTraits[trait] // Defaults to 0 if not present
		diff := userValue - parkValue
		sumSquares += diff * diff
	}

	return math.Sqrt(sumSquares)
}

// getTopTraits identifies the traits that contributed most to the match
func getTopTraits(userTraits, parkTraits map[string]float64, limit int) []string {
	type traitMatch struct {
		name       string
		similarity float64 // How close user and park values are (1.0 = identical)
	}

	var matches []traitMatch
	for trait, userValue := range userTraits {
		if parkValue, exists := parkTraits[trait]; exists {
			// Calculate similarity (inverse of distance, 0-1 scale)
			distance := math.Abs(userValue - parkValue)
			similarity := 1.0 - distance
			matches = append(matches, traitMatch{
				name:       trait,
				similarity: similarity,
			})
		}
	}

	// Sort by similarity (highest first)
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].similarity > matches[j].similarity
	})

	// Return top N trait names
	var topTraits []string
	for i := 0; i < limit && i < len(matches); i++ {
		topTraits = append(topTraits, matches[i].name)
	}

	return topTraits
}
