package activity



import errs "go-tailwind-test/internal/util/err"


var (

	ErrIncomeCategoryRequired = &errs.AppError{
		Code: "INCOME_CATEGORY_REQUIRED",
		Message: "Income category is required",
		StatusCode: 400,
	}

	ErrIncomeTitleRequired = &errs.AppError{
		Code: "INCOME_TITLE_REQUIRED",
		Message: "Income title is required",
		StatusCode: 400,
	}

	ErrInvalidActivityType = &errs.AppError{
		Code: "INVALID_ACTIVITY_TYPE",
		Message: "Invalid activity type",
		StatusCode: 400,
	}
	ErrDatabaseFailure = &errs.AppError{
		Code: "DATABASE_FAILURE",
		Message: "Internal server error",
		StatusCode: 500,
	}

	ErrProjectNameRequired = &errs.AppError{
		Code: "PROJECT_NAME_REQUIRED",
		Message: "Project name is required",
		StatusCode: 400,
	}

	ErrVendorNameRequired = &errs.AppError{
		Code: "VENDOR_NAME_REQUIRED",
		Message: "Vendor name is required",
		StatusCode: 400,
	}

	ErrCostCategoryNameRequired = &errs.AppError{
		Code: "COST_CATEGORY_NAME_REQUIRED",
		Message: "Cost category name is required",
		StatusCode: 400,
	}
	
	ErrAmountInvalid = &errs.AppError{
		Code: "AMOUNT_INVALID",
		Message: "Amount must be greater than 0",
		StatusCode: 400,
	}
	
	ErrDateRequired = &errs.AppError{
		Code: "DATE_REQUIRED",
		Message: "Date is required",
		StatusCode: 400,
	}
	
	ErrEventTypeRequired = &errs.AppError{
		Code: "EVENT_TYPE_REQUIRED",
		Message: "Event type is required",
		StatusCode: 400,
	}
	
	ErrEventTitleRequired = &errs.AppError{
		Code: "EVENT_TITLE_REQUIRED",
		Message: "Event title is required",
		StatusCode: 400,
	}
	
	ErrNotFound = &errs.AppError{
		Code: "NOT_FOUND",
		Message: "Activity not found",
		StatusCode: 404,
	}

	


)
