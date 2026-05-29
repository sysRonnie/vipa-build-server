package home

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/network"

	"github.com/labstack/echo/v4"
)


type Handler struct {
	controller HomeController
}


func NewHomeHandler(controller HomeController) *Handler {
	return &Handler{
		controller: controller,
	}
}


func (h *Handler) RegisterHomeRoutes(g *echo.Group) {
	g.GET("/home-read", h.GetHomeDashboard, auth.Middleware)
}

func (h *Handler) GetHomeDashboard(c echo.Context) error {
	res, err := h.controller.ControlHomeProjectCards(c)
	if err != nil {
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponse(c, res)
}