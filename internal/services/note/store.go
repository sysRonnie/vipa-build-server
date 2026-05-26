package note

import (
	"context"
	"database/sql"
	"go-tailwind-test/internal/util/advisor"
)

type Store struct {
	db *sql.DB
}

func NewNoteStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) QueryNoteList(
	ctx context.Context,
	email string,
	includeDeleted bool,
) (*NoteRowList, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_query_note_list")

	rows, err := s.db.QueryContext(
		ctx,
		queryNoteList,
		email,
		includeDeleted,
	)
	if err != nil {
		advisor.Error("failed to query note list", err)
		return nil, err
	}
	defer rows.Close()

	notes := make([]NoteRow, 0)

	for rows.Next() {
		var note NoteRow

		err = rows.Scan(
			&note.ID,
			&note.UserID,
			&note.ProjectID,
			&note.ProjectName,
			&note.CustomerName,
			&note.NoteBody,
			&note.PhotoURL,
			&note.FlagIsDeleted,
			&note.CreatedAt,
			&note.UpdatedAt,
		)
		if err != nil {
			advisor.Error("failed to scan note row", err)
			return nil, err
		}

		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		advisor.Error("note rows iteration failed", err)
		return nil, err
	}

	return &NoteRowList{
		Notes: notes,
	}, nil
}

func (s *Store) QueryNoteByID(
	ctx context.Context,
	email string,
	noteID int,
) (*NoteRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_query_note_by_id")

	var note NoteRow

	err := s.db.QueryRowContext(
		ctx,
		queryNoteByID,
		email,
		noteID,
	).Scan(
		&note.ID,
		&note.UserID,
		&note.ProjectID,
		&note.ProjectName,
		&note.CustomerName,
		&note.NoteBody,
		&note.PhotoURL,
		&note.FlagIsDeleted,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		advisor.Error("failed to query note by id", err)
		return nil, err
	}

	return &note, nil
}

func (s *Store) InsertNote(
	ctx context.Context,
	email string,
	note NoteRow,
) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_insert_note")
	if note.PhotoURL != nil {
		advisor.Log("note_photo_url= "+ *note.PhotoURL)
	}

	_, err := s.db.ExecContext(
		ctx,
		queryInsertNote,
		email,
		note.NoteBody,
		note.PhotoURL,
		note.ProjectName,
	)
	if err != nil {
		advisor.Error("failed to insert note", err)
		return err
	}

	return nil
}

func (s *Store) UpdateNote(
	ctx context.Context,
	email string,
	note NoteRow,
) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_update_note")

	result, err := s.db.ExecContext(
		ctx,
		queryUpdateNote,
		email,
		note.ID,
		note.NoteBody,
		note.PhotoURL,
		note.ProjectName,
	)
	if err != nil {
		advisor.Error("failed to update note", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		advisor.Error("failed to read rows affected after updating note", err)
		return err
	}

	if rowsAffected == 0 {
		advisor.Error("no rows affected when updating note", sql.ErrNoRows)
		return sql.ErrNoRows
	}

	return nil
}

func (s *Store) RemoveNote(
	ctx context.Context,
	email string,
	noteID int,
) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_remove_note")

	result, err := s.db.ExecContext(
		ctx,
		queryRemoveNote,
		email,
		noteID,
	)
	if err != nil {
		advisor.Error("failed to remove note", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		advisor.Error("failed to read rows affected after removing note", err)
		return err
	}

	if rowsAffected == 0 {
		advisor.Error("no rows affected when removing note", sql.ErrNoRows)
		return sql.ErrNoRows
	}

	return nil
}

func (s *Store) EraseNote(
	ctx context.Context,
	email string,
	noteID int,
) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("store_attached_erase_note")

	result, err := s.db.ExecContext(
		ctx,
		queryEraseNote,
		email,
		noteID,
	)
	if err != nil {
		advisor.Error("failed to erase note", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		advisor.Error("failed to read rows affected after erasing note", err)
		return err
	}

	if rowsAffected == 0 {
		advisor.Error("no rows affected when erasing note", sql.ErrNoRows)
		return sql.ErrNoRows
	}

	return nil
}