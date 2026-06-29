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

func TestFindExerciseLogsByUserIDSuccess(t *testing.T) {
	repository := &fakeRepository{
		findByUserIDFunc: func(ctx context.Context, userID string) ([]ExerciseLog, error) {
			if userID != "user-123" {
				t.Errorf("expected user ID %q, got %q", "user-123", userID)
			}

			return []ExerciseLog{
				{
					ID:              "exercise-1",
					UserID:          "user-123",
					Name:            "Running",
					ExerciseType:    "cardio",
					DurationMinutes: 30,
				},
			}, nil
		},
	}

	service := NewService(repository)

	logs, err := service.FindExerciseLogsByUserID(context.Background(), "user-123")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(logs) != 1 {
		t.Fatalf("expected 1 log, got %d", len(logs))
	}

	if logs[0].UserID != "user-123" {
		t.Errorf("expected user ID %q, got %q", "user-123", logs[0].UserID)
	}

	if logs[0].Name != "Running" {
		t.Errorf("expected exercise name %q, got %q", "Running", logs[0].Name)
	}
}

func TestFindExerciseLogsByUserIDRequiresUserID(t *testing.T) {
	repositoryCalled := false

	repository := &fakeRepository{
		findByUserIDFunc: func(ctx context.Context, userID string) ([]ExerciseLog, error) {
			repositoryCalled = true
			return []ExerciseLog{}, nil
		},
	}

	service := NewService(repository)

	_, err := service.FindExerciseLogsByUserID(context.Background(), "   ")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "user id is required" {
		t.Errorf("expected user ID error, got %q", err.Error())
	}

	if repositoryCalled {
		t.Error("expected repository not to be called")
	}
}
