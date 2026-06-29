package exercises

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type createExerciseLogRequest struct {
	UserID          string   `json:"user_id"`
	Name            string   `json:"name"`
	ExerciseType    string   `json:"exercise_type"`
	DurationMinutes int      `json:"duration_minutes"`
	Sets            *int     `json:"sets,omitempty"`
	Reps            *int     `json:"reps,omitempty"`
	WeightLbs       *float64 `json:"weight_lbs,omitempty"`
	CaloriesBurned  *int     `json:"calories_burned,omitempty"`
	Notes           string   `json:"notes,omitempty"`
}

func (h *Handler) CreateExerciseLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload createExerciseLogRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	log := ExerciseLog{
		UserID:          payload.UserID,
		Name:            payload.Name,
		ExerciseType:    payload.ExerciseType,
		DurationMinutes: payload.DurationMinutes,
		Sets:            payload.Sets,
		Reps:            payload.Reps,
		WeightLbs:       payload.WeightLbs,
		CaloriesBurned:  payload.CaloriesBurned,
		Notes:           payload.Notes,
	}

	createdLog, err := h.service.CreateExerciseLog(r.Context(), log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdLog); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) FindExerciseLogsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	logs, err := h.service.FindExerciseLogsByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(logs); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
