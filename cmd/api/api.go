package api

import (
	"database/sql"
	"log"

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
	if v1 != nil {
		log.Println("API service registered")

	}
	return nil
}