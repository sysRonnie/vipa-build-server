package note

import (
	"context"

	"github.com/labstack/echo/v4"
)

type NoteController interface {
	GetNoteList(c echo.Context) (*NoteRowList, error)
	GetNoteListRecycled(c echo.Context) (*NoteRowList, error)
	GetNoteByID(c echo.Context) (*NoteRow, error)
	CreateNote(c echo.Context) error
	UpdateNote(c echo.Context) error
	RemoveNote(c echo.Context) error
	EraseNote(c echo.Context) error
}

type NoteService interface {
	GetNoteList(ctx context.Context) (*NoteRowList, error)
	GetNoteListRecycled(ctx context.Context) (*NoteRowList, error)
	GetNoteByID(ctx context.Context, noteID int) (*NoteRow, error)
	CreateNote(ctx context.Context, note NoteRow) error
	UpdateNote(ctx context.Context, note NoteRow) error
	RemoveNote(ctx context.Context, noteID int) error
	EraseNote(ctx context.Context, noteID int) error
}

type NoteStore interface {
	QueryNoteList(ctx context.Context, email string, includeDeleted bool) (*NoteRowList, error)
	QueryNoteByID(ctx context.Context, email string, noteID int) (*NoteRow, error)
	InsertNote(ctx context.Context, email string, note NoteRow) error
	UpdateNote(ctx context.Context, email string, note NoteRow) error
	RemoveNote(ctx context.Context, email string, noteID int) error
	EraseNote(ctx context.Context, email string, noteID int) error
}