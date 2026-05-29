package home

import (
	"go-tailwind-test/internal/util/advisor"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service HomeService
}

func NewHomeController(service HomeService) *Controller {
	return &Controller{
		service: service,
	}
}

type HomeController interface {
	ControlHomeProjectCards(ctx echo.Context) (HomeDashboard, error)
}

func (c *Controller) ControlHomeProjectCards(ctx echo.Context) (HomeDashboard, error) {
	advisor := advisor.FromContext(ctx.Request().Context())
	advisor.Log("Processing get home dashboard request")
	return c.service.ServiceHomeProjectCards(ctx.Request().Context())
}