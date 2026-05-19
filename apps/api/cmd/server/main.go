package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sangkihan93/fitness-diet-tracker/apps/api/internal/health"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.Handler)

	log.Printf("API server running on port %s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
