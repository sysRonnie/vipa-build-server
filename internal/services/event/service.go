package event

import (
	"context"
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/network"
)

type Service struct {
	store EventStore
}

func NewEventService(store EventStore) *Service {
	return &Service{store: store}
}

type EventService interface {
	CreateEventActivity(ctx context.Context, req EventActivityRow) error
}

func (s *Service) CreateEventActivity(ctx context.Context, req EventActivityRow) error {

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return err
	}


	email := claims.UserEmail
	if req.ProjectName == "" {
		return network.ErrInvalidRequest
	}

	if req.EventType == "" {
		return network.ErrInvalidRequest
	}

	if req.EventDate == "" {
		return network.ErrInvalidRequest
	}

	return s.store.InsertEventActivity(ctx, email, req)
}