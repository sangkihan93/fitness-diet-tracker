package exercises

import (
	"context"
	"errors"
	"testing"
)

type fakeRepository struct {
	createFunc       func(ctx context.Context, log ExerciseLog) (ExerciseLog, error)
	findByUserIDFunc func(ctx context.Context, userID string) ([]ExerciseLog, error)
}

func (f *fakeRepository) Create(ctx context.Context, log ExerciseLog) (ExerciseLog, error) {
	if f.createFunc != nil {
		return f.createFunc(ctx, log)
	}

	return log, nil
}

func (f *fakeRepository) FindByUserID(ctx context.Context, userID string) ([]ExerciseLog, error) {
	if f.findByUserIDFunc != nil {
		return f.findByUserIDFunc(ctx, userID)
	}

	return []ExerciseLog{}, nil
}

func TestCreateExerciseLogSuccess(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := ExerciseLog{
		UserID:          "user-123",
		Name:            "Running",
		ExerciseType:    "cardio",
		DurationMinutes: 30,
	}

	createdLog, err := service.CreateExerciseLog(context.Background(), input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if createdLog.Name != input.Name {
		t.Errorf("expected name %q, got %q", input.Name, createdLog.Name)
	}

	if createdLog.DurationMinutes != input.DurationMinutes {
		t.Errorf("expected duration %d, got %d", input.DurationMinutes, createdLog.DurationMinutes)
	}
}

func TestCreateExerciseLogRequiresName(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := ExerciseLog{
		UserID:          "user-123",
		Name:            "   ",
		ExerciseType:    "cardio",
		DurationMinutes: 30,
	}

	_, err := service.CreateExerciseLog(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "exercise name is required" {
		t.Errorf("expected exercise name error, got %q", err.Error())
	}
}

func TestCreateExerciseLogRequiresPositiveDuration(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := ExerciseLog{
		UserID:          "user-123",
		Name:            "Running",
		ExerciseType:    "cardio",
		DurationMinutes: 0,
	}

	_, err := service.CreateExerciseLog(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "duration must be greater than zero" {
		t.Errorf("expected duration error, got %q", err.Error())
	}
}

func TestCreateExerciseLogReturnsRepositoryError(t *testing.T) {
	expectedErr := errors.New("repository failed")

	repository := &fakeRepository{
		createFunc: func(ctx context.Context, log ExerciseLog) (ExerciseLog, error) {
			return ExerciseLog{}, expectedErr
		},
	}

	service := NewService(repository)

	input := ExerciseLog{
		UserID:          "user-123",
		Name:            "Running",
		ExerciseType:    "cardio",
		DurationMinutes: 30,
	}

	_, err := service.CreateExerciseLog(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected repository error, got %v", err)
	}
}
