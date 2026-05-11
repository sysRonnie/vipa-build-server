package advisor

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		advisor := NewAdvisor()

		advisor.Log("request_started")
		advisor.Log(
			"method=" + c.Request().Method +
				" path=" + c.Request().URL.Path,
		)

		ctx := WithAdvisor(
			c.Request().Context(),
			advisor,
		)

		c.SetRequest(
			c.Request().WithContext(ctx),
		)

		err := next(c)

		if err != nil {

			advisor.Error(
				"request_failed",
				err,
			)

			c.Error(err)

		} else {

			status := c.Response().Status

			if status >= http.StatusBadRequest {

				advisor.Log(
					"request_completed_with_error_status",
				)

			} else {

				advisor.Log(
					"request_completed_successfully",
				)
			}
		}

		advisor.Flush()

		return err
	}
}
