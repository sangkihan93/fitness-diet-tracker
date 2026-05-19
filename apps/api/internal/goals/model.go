package goals

import "time"

type GoalType string

const (
	GoalWeightLoss    GoalType = "weight_loss"
	GoalMuscleGain    GoalType = "muscle_gain"
	GoalMaintenance   GoalType = "maintenance"
	GoalGeneralHealth GoalType = "general_health"
)

type FitnessGoal struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	GoalType       GoalType  `json:"goal_type"`
	TargetWeight   *float64  `json:"target_weight,omitempty"`
	TargetCalories *int      `json:"target_calories,omitempty"`
	TargetProtein  *int      `json:"target_protein,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
