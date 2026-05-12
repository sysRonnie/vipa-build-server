package project

import (
	"context"
	"go-tailwind-test/internal/util/network"
	"time"
)

type Service struct {
	store ProjectStore
}

func NewProjectService(store ProjectStore) *Service {
	return &Service{
		store: store,
	}
}

type ProjectService interface {
	ProcessCreateProject(ctx context.Context, newProject ProjectRow) error
}

func (s *Service) ProcessCreateProject(
	ctx context.Context,
	newProject ProjectRow,
) error {

	const layout = "2006-01-02"

	if err := checkDateFormat(newProject.StartDate); err != nil {
		return network.ErrInvalidRequest
	}

	if err := checkDateFormat(newProject.EndDateEst); err != nil {
		return network.ErrInvalidRequest
	}

	if err := checkDateFormat(newProject.EndDateActual); err != nil {
		return network.ErrInvalidRequest
	}



	return s.store.InsertProject(ctx, newProject)
}

func checkDateFormat(dateStr string) error {
	const layout = "2006-01-02"
	_, err := time.Parse(layout, dateStr)
	return err
}
