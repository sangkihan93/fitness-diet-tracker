package exercises

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type InMemoryRepository struct {
	mu     sync.Mutex
	logs   []ExerciseLog
	nextID int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		logs:   []ExerciseLog{},
		nextID: 1,
	}
}

func (r *InMemoryRepository) Create(ctx context.Context, log ExerciseLog) (ExerciseLog, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now().UTC()

	if log.ID == "" {
		log.ID = fmt.Sprintf("exercise-%d", r.nextID)
		r.nextID++
	}

	if log.CreatedAt.IsZero() {
		log.CreatedAt = now
	}

	if log.PerformedAt.IsZero() {
		log.PerformedAt = now
	}

	r.logs = append(r.logs, log)

	return log, nil
}

func (r *InMemoryRepository) FindByUserID(ctx context.Context, userID string) ([]ExerciseLog, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	matchingLogs := []ExerciseLog{}

	for _, log := range r.logs {
		if log.UserID == userID {
			matchingLogs = append(matchingLogs, log)
		}
	}

	return matchingLogs, nil
}
