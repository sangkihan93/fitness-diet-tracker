package meals

import (
	"context"
	"errors"
	"strings"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateMealLog(ctx context.Context, mealLog MealLog) (MealLog, error) {
	if strings.TrimSpace(mealLog.Name) == "" {
		return MealLog{}, errors.New("meal name is required")
	}

	if mealLog.Calories < 0 {
		return MealLog{}, errors.New("calories must be greater than or equal to zero")
	}

	if mealLog.CarbsG < 0 {
		return MealLog{}, errors.New("carb must be greater than or equal to zero")
	}

	if mealLog.FatG < 0 {
		return MealLog{}, errors.New("fat must be greater than or equal to zero")
	}

	if mealLog.ProteinG < 0 {
		return MealLog{}, errors.New("protein must be greater than or equal to zero")
	}

	return s.repository.Create(ctx, mealLog)
}
