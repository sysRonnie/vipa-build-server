package project

import errs "go-tailwind-test/internal/util/err"

var (
	ErrDuplicateProjectNotDeleted = &errs.AppError{
		Code:       "DUPLICATE_PROJECT",
		Message:    "Project already exists in your collection",
		StatusCode: 400,
	}

	ErrDuplicateProjectInRecycleBin = &errs.AppError{
		Code:       "DUPLICATE_PROJECT",
		Message:    "Project already exists in your recycle bin",
		StatusCode: 400,
	}
)
