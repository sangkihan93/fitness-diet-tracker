package meals

import (
	"context"
	"errors"
	"testing"
)

type fakeRepository struct {
	createFunc       func(ctx context.Context, mealLog MealLog) (MealLog, error)
	findByUserIDFunc func(ctx context.Context, userID string) ([]MealLog, error)
}

func (f *fakeRepository) Create(ctx context.Context, mealLog MealLog) (MealLog, error) {
	if f.createFunc != nil {
		return f.createFunc(ctx, mealLog)
	}

	return mealLog, nil
}

func (f *fakeRepository) FindByUserID(ctx context.Context, userID string) ([]MealLog, error) {
	if f.findByUserIDFunc != nil {
		return f.findByUserIDFunc(ctx, userID)
	}

	return []MealLog{}, nil
}

func TestCreateMealLogSuccess(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := MealLog{
		UserID:   "user-123",
		Name:     "Chicken Bowl",
		MealType: "dinner",
		Calories: 1200,
		ProteinG: 150,
		CarbsG:   50,
		FatG:     20,
	}

	createdLog, err := service.CreateMealLog(context.Background(), input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if createdLog.Name != input.Name {
		t.Errorf("expected name %q, got %q", input.Name, createdLog.Name)
	}

	if createdLog.Calories != input.Calories {
		t.Errorf("expected calories %d, got %d", input.Calories, createdLog.Calories)
	}

	if createdLog.CarbsG != input.CarbsG {
		t.Errorf("expected carbs %.1f, got %.1f", input.CarbsG, createdLog.CarbsG)
	}

	if createdLog.ProteinG != input.ProteinG {
		t.Errorf("expected protein %.1f, got %.1f", input.ProteinG, createdLog.ProteinG)
	}

	if createdLog.FatG != input.FatG {
		t.Errorf("expected fat %.1f, got %.1f", input.FatG, createdLog.FatG)
	}
}

func TestCreateMealLogRequiresName(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := MealLog{
		UserID:   "user-123",
		Name:     "   ",
		MealType: "dinner",
		Calories: 1200,
		ProteinG: 150,
		CarbsG:   50,
		FatG:     20,
	}

	_, err := service.CreateMealLog(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "meal name is required" {
		t.Errorf("expected meal name error, got %q", err.Error())
	}
}

func TestCreateMealLogRejectsNegativeCalories(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := MealLog{
		UserID:   "user-123",
		Name:     "Chicken Bowl",
		MealType: "dinner",
		Calories: -1200,
		ProteinG: 150,
		CarbsG:   50,
		FatG:     20,
	}

	_, err := service.CreateMealLog(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "calories must be greater than or equal to zero" {
		t.Errorf("expected calories error, got %q", err.Error())
	}
}

func TestCreateMealLogRejectsNegativeCarbs(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := MealLog{
		UserID:   "user-123",
		Name:     "Chicken Bowl",
		MealType: "dinner",
		Calories: 1200,
		ProteinG: 150,
		CarbsG:   -50,
		FatG:     20,
	}

	_, err := service.CreateMealLog(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "carbs must be greater than or equal to zero" {
		t.Errorf("expected carbs error, got %q", err.Error())
	}
}

func TestCreateMealLogRejectsNegativeProtein(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := MealLog{
		UserID:   "user-123",
		Name:     "Chicken Bowl",
		MealType: "dinner",
		Calories: 1200,
		ProteinG: -150,
		CarbsG:   50,
		FatG:     20,
	}

	_, err := service.CreateMealLog(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "protein must be greater than or equal to zero" {
		t.Errorf("expected protein error, got %q", err.Error())
	}
}

func TestCreateMealLogRejectsNegativeFat(t *testing.T) {
	repository := &fakeRepository{}
	service := NewService(repository)

	input := MealLog{
		UserID:   "user-123",
		Name:     "Chicken Bowl",
		MealType: "dinner",
		Calories: 1200,
		ProteinG: 150,
		CarbsG:   50,
		FatG:     -20,
	}

	_, err := service.CreateMealLog(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "fat must be greater than or equal to zero" {
		t.Errorf("expected fat error, got %q", err.Error())
	}
}

func TestCreateMealLogReturnsRepositoryError(t *testing.T) {
	expectedErr := errors.New("repository failed")

	repository := &fakeRepository{
		createFunc: func(ctx context.Context, mealLog MealLog) (MealLog, error) {
			return MealLog{}, expectedErr
		},
	}

	service := NewService(repository)

	input := MealLog{
		UserID:   "user-123",
		Name:     "Chicken Bowl",
		MealType: "dinner",
		Calories: 1200,
		ProteinG: 150,
		CarbsG:   50,
		FatG:     20,
	}

	_, err := service.CreateMealLog(context.Background(), input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected repository error, got %v", err)
	}
}
