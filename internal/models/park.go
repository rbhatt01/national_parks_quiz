package models

// Park represents a national park with its characteristics and trait profile
type Park struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	State           string          `json:"state"`
	PrimaryGroup    string          `json:"primary_group"`
	SecondaryGroups []string        `json:"secondary_groups"`
	Traits          map[string]float64 `json:"traits"` // e.g., {"energy": 0.6, "social": 0.8}
	Tags            []string        `json:"tags"`
	ImageURL        string          `json:"image_url"`
	Description     string          `json:"description"`
}
