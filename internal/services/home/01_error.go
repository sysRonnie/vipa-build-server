package home 


import (
	errs "go-tailwind-test/internal/util/err"
)

var (
	ErrNoteIDRequired = &errs.AppError{
		Code: "NOTE_ID_REQUIRED",
		Message: "Note ID is required",
		StatusCode: 400,
	}
)	
	