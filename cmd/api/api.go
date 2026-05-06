package api

import (
	"database/sql"
	"go-tailwind-test/internal/services/user"

	"github.com/labstack/echo/v4"
)


type API struct {
	db *sql.DB
}

func NewAPIServer(db *sql.DB) *API {
	return &API {
		db: db,
	}
}

func (s *API) APIService(e *echo.Echo) error {
	v1 := e.Group("/api/v1")

	userStore := user.NewUserStore(s.db)
	userService := user.NewService(userStore)
	userHandler := user.NewHandler(userService, userStore)
	userHandler.RegisterUserRoutes(v1)

	
	return nil
}