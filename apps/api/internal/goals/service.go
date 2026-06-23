package goals

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

func (s *Service) CreateFitnessGoal(ctx context.Context, goal FitnessGoal) (FitnessGoal, error) {
	if strings.TrimSpace(string(goal.GoalType)) == "" {
		return FitnessGoal{}, errors.New("goal type is required")
	}

	if goal.TargetValue <= 0 {
		return FitnessGoal{}, errors.New("target value must be greater than zero")
	}

	if strings.TrimSpace(goal.Unit) == "" {
		return FitnessGoal{}, errors.New("unit is required")
	}

	return s.repository.Create(ctx, goal)
}
