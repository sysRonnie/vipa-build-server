package advisor

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type contextKey string

const AdvisorContextKey contextKey = "advisor"

type Advisor struct {
	RequestID string
	StartedAt time.Time
	Logs      []AdvisorLog
	mu        sync.Mutex
}

type AdvisorLog struct {
	Timestamp time.Time
	Message   string
	Error     string
}

func NewAdvisor() *Advisor {
	return &Advisor{
		RequestID: uuid.NewString(),
		StartedAt: time.Now().UTC(),
		Logs:      []AdvisorLog{},
	}
}

func (a *Advisor) Log(message string) {

	if a == nil {
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.Logs = append(
		a.Logs,
		AdvisorLog{
			Timestamp: time.Now().UTC(),
			Message:   message,
		},
	)
}

func (a *Advisor) Error(
	message string,
	err error,
) {

	if a == nil {
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	errorMessage := ""

	if err != nil {
		errorMessage = err.Error()
	}

	a.Logs = append(
		a.Logs,
		AdvisorLog{
			Timestamp: time.Now().UTC(),
			Message:   message,
			Error:     errorMessage,
		},
	)
}

func (a *Advisor) Flush() {
	a.mu.Lock()
	defer a.mu.Unlock()

	duration := time.Since(a.StartedAt)

	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Println("REQUEST ID:", a.RequestID)
	log.Println("DURATION:", duration)

	for _, entry := range a.Logs {

		if entry.Error != "" {

			log.Printf(
				"[ADVISOR] ts=%s msg=%s error=%s",
				entry.Timestamp.Format(time.RFC3339),
				entry.Message,
				entry.Error,
			)

			continue
		}

		log.Printf(
			"[ADVISOR] ts=%s msg=%s",
			entry.Timestamp.Format(time.RFC3339),
			entry.Message,
		)
	}

	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func WithAdvisor(
	ctx context.Context,
	advisor *Advisor,
) context.Context {
	return context.WithValue(
		ctx,
		AdvisorContextKey,
		advisor,
	)
}

func FromContext(
	ctx context.Context,
) *Advisor {

	advisor, ok := ctx.Value(
		AdvisorContextKey,
	).(*Advisor)

	if ok {
		return advisor
	}

	return &Advisor{}
}