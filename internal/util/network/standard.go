package network

import (
	errs "go-tailwind-test/internal/util/err"

	"github.com/labstack/echo/v4"
)


var (
	ErrInvalidPayload = &errs.AppError{
		Code:       "INVALID_PAYLOAD",
		Message:    "Invalid payload",
		StatusCode: 400,
	}
	ErrValidationFailed = &errs.AppError{
		Code:       "VALIDATION_FAILED",
		Message:    "Data validation failed",
		StatusCode: 400,
	}
	ErrInvalidRequest = &errs.AppError{
		Code:       "INVALID_REQUEST",
		Message:    "Invalid request",
		StatusCode: 400,
	}
	ErrUnauthorized = &errs.AppError{
		Code:       "UNAUTHORIZED",
		Message:    "Unauthorized",
		StatusCode: 401,
	}
	ErrDatabaseFailure = &errs.AppError{
		Code:       "DATABASE_FAILURE",
		Message:    "Internal server error",
		StatusCode: 500,
	}

)

var (
	SuccessResponseOK = SandboxResponse{
		StatusCode: 200,
		Message:    "ok",
	}
)

func BuildSuccessResponse(c echo.Context, data any) error {
	return Respond(c, SandboxResponse{
		StatusCode: 200,
		Message: "Success!",
		Data: data,
	})
}

func BuildSuccessResponseOK(c echo.Context) error {
	return Respond(c, SandboxResponse{
		StatusCode: 200,
		Message: "Success!",
	})
}

