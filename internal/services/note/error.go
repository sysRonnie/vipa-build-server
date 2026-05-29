package note

import (
	errs "go-tailwind-test/internal/util/err"
)

var (
	ErrNoteIDRequired = &errs.AppError{
		Code: "NOTE_ID_REQUIRED",
		Message: "Note ID is required",
		StatusCode: 400,
	}
	ErrNoteBodyRequired = &errs.AppError{
		Code: "NOTE_BODY_REQUIRED",
		Message: "Note body is required",
		StatusCode: 400,
	}
	ErrNoteNotFound = &errs.AppError{
		Code: "NOTE_NOT_FOUND",
		Message: "Note not found",
		StatusCode: 404,
	}
	ErrProjectNotFound = &errs.AppError{
		Code: "PROJECT_NOT_FOUND",
		Message: "Project not found",
		StatusCode: 404,
	}
)	
	