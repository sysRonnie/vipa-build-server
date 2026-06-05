package vendor

import errs "go-tailwind-test/internal/util/err"

var (
	ErrDuplicateVendorNotDeleted = &errs.AppError{
		Code:       "DUPLICATE_VENDOR",
		Message:    "Vendor already exists in your collection",
		StatusCode: 400,
	}
	ErrDuplicateVendorInRecycleBin = &errs.AppError{
		Code:       "DUPLICATE_VENDOR",
		Message:    "Vendor already exists in your recycle bin.",
		StatusCode: 400,
	}
)