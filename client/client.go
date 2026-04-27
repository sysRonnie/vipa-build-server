package client

import (
	"database/sql"
	"go-tailwind-test/client/feature/expense"
	"go-tailwind-test/client/page"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	db *sql.DB
}

func NewClientServer(db *sql.DB) *ClientHandler {
	return &ClientHandler{
		db: db,
	}
}

func (h *ClientHandler) ClientService(e *echo.Echo) error {
	e.GET("/", ShowLandingPage)
	e.GET("/home", ShowHomePage)


	v1 := e.Group("")
	expenseStore := expense.NewExpenseStore(h.db)
	expenseService := expense.NewService(expenseStore)
	expenseHandler := expense.NewHandler(expenseService, expenseStore)
	expenseHandler.RegisterExpenseRoutes(v1)

	return nil
}

func ShowLandingPage(c echo.Context) error {
	return Render(c, page.LandingPage())
}

func ShowHomePage(c echo.Context) error {
	return Render(c, page.HomePage())
}

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}