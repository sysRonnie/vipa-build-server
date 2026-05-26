package note

import "errors"

var (
	ErrNoteIDRequired = errors.New("note id is required")
	ErrNoteBodyRequired = errors.New("note body is required")
	ErrNoteNotFound = errors.New("note not found")
	ErrProjectNotFound = errors.New("project not found")
)