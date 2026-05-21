package auth

import (
	"context"
	"errors"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type ContextKey string

const ClaimsContextKey ContextKey = "auth_claims"

func GetClaimsFromContext(ctx echo.Context) (*Claims) {
	claims  := ctx.Get(string(ClaimsContextKey)).(*Claims)

	return claims
}
func GetAdvisorClaims(c echo.Context) (*advisor.Advisor, context.Context) {
	ctx := c.Request().Context()
	advisor := advisor.FromContext(ctx)

	return advisor, ctx
}

func GetClaims(ctx context.Context) (*Claims, error) {
	claims, ok := ctx.Value(ClaimsContextKey).(*Claims)

	if !ok || claims == nil {
		return nil, errors.New("claims not found in context")
	}

	return claims, nil
}

func Middleware(next echo.HandlerFunc,) echo.HandlerFunc {

	return func(c echo.Context) error {
		advisor := advisor.FromContext(c.Request().Context())
		advisor.Log("authentiation_middleware_engaged" )

		authHeader :=
			c.Request().Header.Get(
				"Authorization",
			)


		if authHeader == "" {

			advisor.Error("missing authorization header", network.ErrInvalidRequest)

			return c.JSON(
				http.StatusUnauthorized,
				map[string]any{
					"statusCode": 401,
					"message":    "Missing authorization header",
					"data":       map[string]any{},
				},
			)
		}

		splitToken :=
			strings.Split(
				authHeader,
				"Bearer ",
			)

		if len(splitToken) != 2 {


			

			advisor.Error("malformed bearer token", network.ErrInvalidRequest)
			return c.JSON(
				http.StatusUnauthorized,
				map[string]any{
					"statusCode": 401,
					"message":    "Invalid authorization header",
					"data":       map[string]any{},
				},
			)
		}

		tokenString := strings.TrimSpace(splitToken[1])


		

		advisor.Log("token string extracted from header: " + tokenString)

		claims, err :=
			ParseAccessToken(
				tokenString,
			)

		if err != nil {

			advisor.Error("invalid or expired token", network.ErrInvalidRequest)


			return c.JSON(
				http.StatusUnauthorized,
				map[string]any{
					"statusCode": 401,
					"message":    "Invalid or expired token",
					"data":       map[string]any{},
				},
			)
		}



		ctx := context.WithValue(
			c.Request().Context(),
			ClaimsContextKey,
			claims,
		)

		c.SetRequest(
			c.Request().WithContext(ctx),
		)

		c.Set(
			string(ClaimsContextKey),
			claims,
		)	

			return next(c)
		}
}
