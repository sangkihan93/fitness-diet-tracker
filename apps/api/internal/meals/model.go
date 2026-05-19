package meals

import "time"

type MealLog struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	MealType  string    `json:"meal_type"`
	Calories  int       `json:"calories"`
	ProteinG  float64   `json:"protein_g"`
	CarbsG    float64   `json:"carbs_g"`
	FatG      float64   `json:"fat_g"`
	Notes     string    `json:"notes,omitempty"`
	EatenAt   time.Time `json:"eaten_at"`
	CreatedAt time.Time `json:"created_at"`
}
