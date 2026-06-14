package meals

import "context"

type Repository interface {
	Create(ctx context.Context, mealLog MealLog) (MealLog, error)
	FindByUserID(ctx context.Context, userID string) ([]MealLog, error)
}
