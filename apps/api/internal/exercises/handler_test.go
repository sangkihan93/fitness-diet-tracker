package exercises

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateExerciseLogHandlerSuccess(t *testing.T) {
	repository := NewInMemoryRepository()
	service := NewService(repository)
	handler := NewHandler(service)

	requestBody := `{
		"user_id": "user-123",
		"name": "Running",
		"exercise_type": "cardio",
		"duration_minutes": 30
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/exercises",
		bytes.NewBufferString(requestBody),
	)

	response := httptest.NewRecorder()

	handler.CreateExerciseLog(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, response.Code)
	}

	var createdLog ExerciseLog
	err := json.NewDecoder(response.Body).Decode(&createdLog)
	if err != nil {
		t.Fatalf("expected valid JSON response, got error %v", err)
	}

	if createdLog.ID == "" {
		t.Fatal("expected ID to be assigned")
	}

	if createdLog.UserID != "user-123" {
		t.Errorf("expected user ID %q, got %q", "user-123", createdLog.UserID)
	}

	if createdLog.Name != "Running" {
		t.Errorf("expected name %q, got %q", "Running", createdLog.Name)
	}

	if createdLog.ExerciseType != "cardio" {
		t.Errorf("expected exercise type %q, got %q", "cardio", createdLog.ExerciseType)
	}

	if createdLog.DurationMinutes != 30 {
		t.Errorf("expected duration %d, got %d", 30, createdLog.DurationMinutes)
	}
}

func TestCreateExerciseLogHandlerInvalidJSON(t *testing.T) {
	repository := NewInMemoryRepository()
	service := NewService(repository)
	handler := NewHandler(service)

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/exercises",
		bytes.NewBufferString("{invalid-json"),
	)

	response := httptest.NewRecorder()

	handler.CreateExerciseLog(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}
}

func TestCreateExerciseLogHandlerMissingName(t *testing.T) {
	repository := NewInMemoryRepository()
	service := NewService(repository)
	handler := NewHandler(service)

	requestBody := `{
		"user_id": "user-123",
		"name": "",
		"exercise_type": "cardio",
		"duration_minutes": 30
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/exercises",
		bytes.NewBufferString(requestBody),
	)

	response := httptest.NewRecorder()

	handler.CreateExerciseLog(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}

	expectedBody := "exercise name is required\n"
	if response.Body.String() != expectedBody {
		t.Errorf("expected response body %q, got %q", expectedBody, response.Body.String())
	}
}

func TestCreateExerciseLogHandlerInvalidDuration(t *testing.T) {
	repository := NewInMemoryRepository()
	service := NewService(repository)
	handler := NewHandler(service)

	requestBody := `{
		"user_id": "user-123",
		"name": "Running",
		"exercise_type": "cardio",
		"duration_minutes": 0
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/exercises",
		bytes.NewBufferString(requestBody),
	)

	response := httptest.NewRecorder()

	handler.CreateExerciseLog(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}

	expectedBody := "duration must be greater than zero\n"
	if response.Body.String() != expectedBody {
		t.Errorf("expected response body %q, got %q", expectedBody, response.Body.String())
	}
}

func TestFindExerciseLogsByUserIDHandlerSuccess(t *testing.T) {
	repository := NewInMemoryRepository()
	service := NewService(repository)
	handler := NewHandler(service)

	_, err := repository.Create(context.Background(), ExerciseLog{
		UserID:          "user-123",
		Name:            "Running",
		ExerciseType:    "cardio",
		DurationMinutes: 30,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/exercises?user_id=user-123",
		nil,
	)

	response := httptest.NewRecorder()

	handler.FindExerciseLogsByUserID(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	var logs []ExerciseLog
	err = json.NewDecoder(response.Body).Decode(&logs)
	if err != nil {
		t.Fatalf("expected valid JSON response, got error %v", err)
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

func TestFindExerciseLogsByUserIDHandlerMissingUserID(t *testing.T) {
	repository := NewInMemoryRepository()
	service := NewService(repository)
	handler := NewHandler(service)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/exercises",
		nil,
	)

	response := httptest.NewRecorder()

	handler.FindExerciseLogsByUserID(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}

	expectedBody := "user id is required\n"
	if response.Body.String() != expectedBody {
		t.Errorf("expected response body %q, got %q", expectedBody, response.Body.String())
	}
}
