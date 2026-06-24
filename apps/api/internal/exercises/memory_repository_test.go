package exercises

import (
	"context"
	"testing"
)

func TestInMemoryRepositoryCreate(t *testing.T) {
	repository := NewInMemoryRepository()

	input := ExerciseLog{
		UserID:          "user-123",
		Name:            "Running",
		DurationMinutes: 30,
	}

	createdLog, err := repository.Create(context.Background(), input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if createdLog.ID == "" {
		t.Fatal("expected ID to be assigned")
	}

	if createdLog.CreatedAt.IsZero() {
		t.Fatal("expected CreatedAt to be set")
	}

	if createdLog.PerformedAt.IsZero() {
		t.Fatal("expected PerformedAt to be set")
	}

	if createdLog.UserID != input.UserID {
		t.Errorf("expected user ID %q, got %q", input.UserID, createdLog.UserID)
	}

	if createdLog.Name != input.Name {
		t.Errorf("expected name %q, got %q", input.Name, createdLog.Name)
	}

	if createdLog.DurationMinutes != input.DurationMinutes {
		t.Errorf("expected duration %d, got %d", input.DurationMinutes, createdLog.DurationMinutes)
	}
}

func TestInMemoryRepositoryFindByUserID(t *testing.T) {
	repository := NewInMemoryRepository()

	_, err := repository.Create(context.Background(), ExerciseLog{
		UserID:          "user-123",
		Name:            "Running",
		DurationMinutes: 30,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = repository.Create(context.Background(), ExerciseLog{
		UserID:          "user-456",
		Name:            "Cycling",
		DurationMinutes: 45,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	logs, err := repository.FindByUserID(context.Background(), "user-123")
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

func TestInMemoryRepositoryFindByUserIDReturnsEmptyListWhenNoLogsMatch(t *testing.T) {
	repository := NewInMemoryRepository()

	_, err := repository.Create(context.Background(), ExerciseLog{
		UserID:          "user-456",
		Name:            "Cycling",
		DurationMinutes: 45,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	logs, err := repository.FindByUserID(context.Background(), "user-123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(logs) != 0 {
		t.Fatalf("expected 0 logs, got %d", len(logs))
	}
}
