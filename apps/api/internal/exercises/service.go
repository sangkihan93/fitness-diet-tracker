package exercises

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

func (s *Service) CreateExerciseLog(ctx context.Context, log ExerciseLog) (ExerciseLog, error) {
	if strings.TrimSpace(log.Name) == "" {
		return ExerciseLog{}, errors.New("exercise name is required")
	}

	if log.DurationMinutes <= 0 {
		return ExerciseLog{}, errors.New("duration must be greater than zero")
	}

	return s.repository.Create(ctx, log)
}
