package auth

import (
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type ContextKey string

const ClaimsContextKey ContextKey = "auth_claims"

func GetClaimsFromContext(ctx echo.Context) (*Claims) {
	claims  := ctx.Get(string(ClaimsContextKey)).(*Claims)

	return claims
}

func Middleware(next echo.HandlerFunc,) echo.HandlerFunc {

	return func(c echo.Context) error {
		advisor := advisor.FromContext(c.Request().Context())
		advisor.Log("authentiation_middleware_engaged" )

		authHeader :=
			c.Request().Header.Get(
				"Authorization",
			)

		log.Println(
			"AUTH HEADER:",
			authHeader,
		)

		if authHeader == "" {

			log.Println(
				"AUTH ERROR: missing authorization header",
			)

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

			log.Println(
				"AUTH ERROR: malformed bearer token",
			)

			

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

		tokenString :=
			strings.TrimSpace(
				splitToken[1],
			)

		log.Println(
			"AUTH TOKEN:",
			tokenString,
		)

		

		advisor.Log("token string extracted from header: " + tokenString)

		claims, err :=
			ParseAccessToken(
				tokenString,
			)

		if err != nil {

			log.Println(
				"AUTH ERROR: invalid or expired token",
			)

			log.Println(
				"TOKEN PARSE ERROR:",
				err,
			)

			return c.JSON(
				http.StatusUnauthorized,
				map[string]any{
					"statusCode": 401,
					"message":    "Invalid or expired token",
					"data":       map[string]any{},
				},
			)
		}

		log.Println(
			"TOKEN ISSUED AT:",
			claims.IssuedAt.Time.UTC(),
		)

		log.Println(
			"TOKEN EXPIRES:",
			claims.ExpiresAt.Time.UTC(),
		)

		log.Println(
			"CURRENT UTC TIME:",
			time.Now().UTC(),
		)

		log.Println(
			"TOKEN EXPIRED:",
			time.Now().UTC().After(
				claims.ExpiresAt.Time.UTC(),
			),
		)

		c.Set(
			string(ClaimsContextKey),
			claims,
		)

		return next(c)
	}
}
