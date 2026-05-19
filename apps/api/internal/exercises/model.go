package exercises

import "time"

type ExerciseLog struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	Name            string    `json:"name"`
	ExerciseType    string    `json:"exercise_type"`
	DurationMinutes int       `json:"duration_minutes"`
	Sets            *int      `json:"sets,omitempty"`
	Reps            *int      `json:"reps,omitempty"`
	WeightLbs       *float64  `json:"weight_lbs,omitempty"`
	CaloriesBurned  *int      `json:"calories_burned,omitempty"`
	Notes           string    `json:"notes,omitempty"`
	PerformedAt     time.Time `json:"performed_at"`
	CreatedAt       time.Time `json:"created_at"`
}
