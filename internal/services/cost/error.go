package cost

import errs "go-tailwind-test/internal/util/err"

var (
	ErrDuplicateCostNotDeleted = &errs.AppError{
		Code:       "DUPLICATE_COST",
		Message:    "Cost category exists in your collection",
		StatusCode: 400,
	}

	ErrDuplicateCostInRecycleBin = &errs.AppError{
		Code:       "DUPLICATE_COST",
		Message:    "Cost category exists in your recycle bin",
		StatusCode: 400,
	}
)