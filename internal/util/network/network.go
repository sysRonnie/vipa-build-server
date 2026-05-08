package network

import (
	"errors"
	errs "go-tailwind-test/internal/util/err"
	"log"

	"github.com/labstack/echo/v4"
)

// Response package provides standard response structures for API responses.


type SuccessResponse struct {
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SandboxRoute struct {
	URL string 
	Method string
}

func FailFromError(
	c echo.Context,
	err error,
) error {

	var appErr *errs.AppError

	if errors.As(err, &appErr) {

		return c.JSON(
			appErr.StatusCode,
			SandboxResponse{
				StatusCode: appErr.StatusCode,
				Message: appErr.Message,
			},
		)
	}

	return c.JSON(
		500,
		SandboxResponse{
			StatusCode: 500,
			Message: "Internal server error",
		},
	)
}

// SandboxResponse standardizes API success and error responses.
// InternalMessage and Err are for internal use only (not serialized to JSON).
// internal/shared/response/response.go
type SandboxResponse struct {
	StatusCode      int         `json:"statusCode"`
	Message         string      `json:"message,omitempty"`
	ErrorMessage          string      `json:"error,omitempty"`
	Data            interface{} `json:"data"` // always included
	InternalMessage string      `json:"-"`
	InternalError error `json:"-"`
}
// MARK: - Core Respond Logic
func Respond(c echo.Context, r SandboxResponse) error {
	// Ensure Data is non-nil for consistent JSON structure
	if r.Data == nil {
		r.Data = map[string]any{}
	}

	// Internal logging
	if r.InternalMessage != "" {
		log.Println("[RESP]", r.LogText())
	}

	return c.JSON(r.StatusCode, r)
}

// MARK: - Helper Builders
func (r SandboxResponse) AttachError(err error) SandboxResponse {
	r.InternalError = err
	return r
}

func (r SandboxResponse) LogText() string {
	if r.InternalError != nil {
		return r.InternalMessage + ": " + r.InternalError.Error()
	}
	return r.InternalMessage
}

// MARK: - Success / Fail Shortcuts
// These allow concise one-line calls in handlers like:
// return response.Success(c, ForumSuccessMap.GetPosts, posts)
// return response.Fail(c, ForumErrorMap.DBQueryFailed)

func Success(c echo.Context, template SandboxResponse, payload ...any) error {
	if len(payload) > 0 {
		template.Data = payload[0]
	}
	return Respond(c, template)
}

func Fail(c echo.Context, template SandboxResponse, customMessage ...string) error {
	if len(customMessage) > 0 {
		template.Message = customMessage[0]
	}
	return Respond(c, template)
}
// BindAndValidate binds JSON and validates the struct in one call.
// Returns a non-nil error if either step fails.
func BindAndValidate(c echo.Context, v any) error {
	if err := c.Bind(v); err != nil {
		return err
	}
	if err := c.Validate(v); err != nil {
		return err
	}
	return nil
}
// Make SandboxResponse compatible with Go's error interface
func (r SandboxResponse) Error() string {
	if r.InternalError != nil {
		return r.InternalError.Error()
	}
	if r.Message != "" {
		return r.Message
	}
	return r.InternalMessage
}

func (r SandboxResponse) Err() error {
	return r
}