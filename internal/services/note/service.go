package note

import (
	"context"

	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
)

type Service struct {
	store NoteStore
}

func NewNoteService(store NoteStore) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) GetNoteList(
	ctx context.Context,
) (*NoteRowList, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_get_note_list")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims from context", err)
		return nil, err
	}

	return s.store.QueryNoteList(
		ctx,
		claims.UserEmail,
		false,
	)
}

func (s *Service) GetNoteListRecycled(
	ctx context.Context,
) (*NoteRowList, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_get_note_list_recycled")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims from context", err)
		return nil, err
	}

	return s.store.QueryNoteList(
		ctx,
		claims.UserEmail,
		true,
	)
}

func (s *Service) GetNoteByID(
	ctx context.Context,
	noteID int,
) (*NoteRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_get_note_by_id")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims from context", err)
		return nil, err
	}

	return s.store.QueryNoteByID(
		ctx,
		claims.UserEmail,
		noteID,
	)
}

func (s *Service) CreateNote(
	ctx context.Context,
	note NoteRow,
) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_create_note")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims from context", err)
		return err
	}

	return s.store.InsertNote(
		ctx,
		claims.UserEmail,
		note,
	)
}

func (s *Service) UpdateNote(
	ctx context.Context,
	note NoteRow,
) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_update_note")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims from context", err)
		return err
	}

	return s.store.UpdateNote(
		ctx,
		claims.UserEmail,
		note,
	)
}

func (s *Service) RemoveNote(
	ctx context.Context,
	noteID int,
) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_remove_note")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims from context", err)
		return err
	}

	return s.store.RemoveNote(
		ctx,
		claims.UserEmail,
		noteID,
	)
}

func (s *Service) EraseNote(
	ctx context.Context,
	noteID int,
) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("service_attached_erase_note")

	claims, err := auth.GetClaims(ctx)
	if err != nil {
		advisor.Error("failed to get claims from context", err)
		return err
	}

	return s.store.EraseNote(
		ctx,
		claims.UserEmail,
		noteID,
	)
}