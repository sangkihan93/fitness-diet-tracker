package exercises

import "context"

type Repository interface {
	Create(ctx context.Context, log ExerciseLog) (ExerciseLog, error)
	FindByUserID(ctx context.Context, userID string) ([]ExerciseLog, error)
}
