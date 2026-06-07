package event

import errs "go-tailwind-test/internal/util/err"


var (
	ErrDuplicateEventNotDeleted = &errs.AppError{
		Code:       "DUPLICATE_EVENT",
		Message:    "Event category exists in your collection",
		StatusCode: 400,
	}

	ErrDuplicateEventInRecycleBin = &errs.AppError{
		Code:       "DUPLICATE_EVENT",
		Message:    "Event category exists in your recycle bin",
		StatusCode: 400,
	}
)