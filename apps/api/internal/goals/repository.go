package goals

import "context"

type Repository interface {
	Create(ctx context.Context, goal FitnessGoal) (FitnessGoal, error)
	FindByUserID(ctx context.Context, userID string) ([]FitnessGoal, error)
}
