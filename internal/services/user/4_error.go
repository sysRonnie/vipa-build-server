package user

import errs "go-tailwind-test/internal/util/err"


var (
	ErrLoginFailed = &errs.AppError{
		Code: "LOGIN_FAILED",
		Message: "Login failed",
		StatusCode: 401,
	}
	ErrDatabaseFailure = &errs.AppError{
		Code: "DATABASE_FAILURE",
		Message: "Internal server error",
		StatusCode: 500,
	}

	ErrInvalidRefreshToken = &errs.AppError{
		Code: "INVALID_REFRESH_TOKEN",
		Message: "Invalid refresh token",
		StatusCode: 401,
	}

	ErrUserNotAuthorized = &errs.AppError{
		Code: "USER_NOT_AUTHORIZED",
		Message: "User not authorized",
		StatusCode: 403,
	}
)