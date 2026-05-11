package customer

import errs "go-tailwind-test/internal/util/err"


var (

	ErrDatabaseFailure = &errs.AppError{
		Code: "DATABASE_FAILURE",
		Message: "Internal server error",
		StatusCode: 500,
	}
	ErrFullNameRequired = &errs.AppError{
		Code: "FULL_NAME_REQUIRED",
		Message: "Client first name and last name are required",
		StatusCode: 400,
	}
	ErrCustomerAlreadyExists = &errs.AppError{
		Code: "CUSTOMER_ALREADY_EXISTS",
		Message: "Client already exists",
		StatusCode: 400,
	}
	ErrCustomerAlreadyExistsRecycled = &errs.AppError{
		Code: "CUSTOMER_ALREADY_EXISTS_RECYCLED",
		Message: "Client exists in the recycle bin",
		StatusCode: 400,
	}

)