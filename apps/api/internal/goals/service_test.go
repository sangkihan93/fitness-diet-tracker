package goals

import (
	"context"
	"errors"
	"testing"
)

type fakeRepository struct {
	createFunc       func(ctx context.Context, goal FitnessGoal) (FitnessGoal, error)
	findByUserIDFunc func(ctx context.Context, userID string) ([]FitnessGoal, error)
}

func (f *fakeRepository) Create(ctx context.Context, goal FitnessGoal) (FitnessGoal, error) {
	if f.createFunc != nil {
		return f.createFunc(ctx, goal)
	}

	return goal, nil
}

func (f *fakeRepository) FindByUserID(ctx context.Context, userID string) ([]FitnessGoal, error) {
	if f.findByUserIDFunc != nil {
		return f.findByUserIDFunc(ctx, userID)
	}

	return []FitnessGoal{}, nil
}

func TestCreateFitnessGoalSuccess(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := FitnessGoal{
		UserID:      "user-123",
		GoalType:    GoalMuscleGain,
		TargetValue: 170,
		Unit:        "lbs",
	}

	createdGoal, err := service.CreateFitnessGoal(context.Background(), input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if createdGoal.GoalType != input.GoalType {
		t.Errorf("expected goal type %q, got %q", input.GoalType, createdGoal.GoalType)
	}

	if createdGoal.TargetValue != input.TargetValue {
		t.Errorf("expected target value %.1f, got %.1f", input.TargetValue, createdGoal.TargetValue)
	}

	if createdGoal.Unit != input.Unit {
		t.Errorf("expected unit %q, got %q", input.Unit, createdGoal.Unit)
	}
}

func TestCreateFitnessGoalRequiresGoalType(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := FitnessGoal{
		UserID:      "user-123",
		TargetValue: 170,
		Unit:        "lbs",
	}

	_, err := service.CreateFitnessGoal(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "goal type is required" {
		t.Errorf("expected goal type error, got %q", err.Error())
	}
}

func TestCreateFitnessGoalRejectsNegativeTargetValue(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := FitnessGoal{
		UserID:      "user-123",
		GoalType:    GoalWeightLoss,
		TargetValue: -170,
		Unit:        "lbs",
	}

	_, err := service.CreateFitnessGoal(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "target value must be greater than zero" {
		t.Errorf("expected target value error, got %q", err.Error())
	}
}

func TestCreateFitnessGoalRejectsZeroTargetValue(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := FitnessGoal{
		UserID:      "user-123",
		GoalType:    GoalWeightLoss,
		TargetValue: 0,
		Unit:        "lbs",
	}

	_, err := service.CreateFitnessGoal(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "target value must be greater than zero" {
		t.Errorf("expected target value error, got %q", err.Error())
	}
}

func TestCreateFitnessGoalRequiresUnit(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := FitnessGoal{
		UserID:      "user-123",
		GoalType:    GoalMaintenance,
		TargetValue: 170,
		Unit:        "   ",
	}

	_, err := service.CreateFitnessGoal(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "unit is required" {
		t.Errorf("expected unit error, got %q", err.Error())
	}
}

func TestCreateFitnessGoalReturnsRepositoryError(t *testing.T) {
	expectedErr := errors.New("repository failed")

	repository := &fakeRepository{
		createFunc: func(ctx context.Context, goal FitnessGoal) (FitnessGoal, error) {
			return FitnessGoal{}, expectedErr
		},
	}

	service := NewService(repository)

	input := FitnessGoal{
		UserID:      "user-123",
		GoalType:    GoalMuscleGain,
		TargetValue: 170,
		Unit:        "lbs",
	}

	_, err := service.CreateFitnessGoal(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected repository error, got %v", err)
	}
}
