package meals

import "context"

type Repository interface {
	Create(ctx context.Context, MealLog MealLog) (MealLog, error)
	FindByUserID(ctx context.Context, userID string) ([]MealLog, error)
}
