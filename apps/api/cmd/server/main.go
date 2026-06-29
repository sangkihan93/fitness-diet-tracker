package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sangkihan93/fitness-diet-tracker/apps/api/internal/exercises"
	"github.com/sangkihan93/fitness-diet-tracker/apps/api/internal/health"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	exerciseRepository := exercises.NewInMemoryRepository()
	exerciseService := exercises.NewService(exerciseRepository)
	exerciseHandler := exercises.NewHandler(exerciseService)

	mux.HandleFunc("/api/exercises", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			exerciseHandler.CreateExerciseLog(w, r)
		case http.MethodGet:
			exerciseHandler.FindExerciseLogsByUserID(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/health", health.Handler)

	log.Printf("API server running on port %s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
