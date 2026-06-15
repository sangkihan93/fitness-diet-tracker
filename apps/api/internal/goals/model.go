package goals

import "time"

type GoalType string

const (
	GoalWeightLoss    GoalType = "weight_loss"
	GoalMuscleGain    GoalType = "muscle_gain"
	GoalMaintenance   GoalType = "maintenance"
	GoalGeneralHealth GoalType = "general_health"
)

type GoalStatus string

const (
	GoalStatusActive    GoalStatus = "active"
	GoalStatusCompleted GoalStatus = "completed"
	GoalStatusPaused    GoalStatus = "paused"
)

type FitnessGoal struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	GoalType    GoalType   `json:"goal_type"`
	TargetValue float64    `json:"target_value"`
	Unit        string     `json:"unit"`
	StartDate   time.Time  `json:"start_date"`
	TargetDate  *time.Time `json:"target_date,omitempty"`
	Status      GoalStatus `json:"status"`
	Notes       string     `json:"notes,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
