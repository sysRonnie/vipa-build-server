package user

import (
	"log"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service UserService
	store UserStore
}

func NewHandler(service UserService, store UserStore) *Handler {
	return &Handler{
		service: service,
		store: store,
	}
}

type LoginRequest struct {
	Token string `json:"token"`
}

func (h *Handler) RegisterUserRoutes(g *echo.Group) {
	g.POST("/login-user", h.HandleSubmitRegisterToken)
}

func (h *Handler) HandleSubmitRegisterToken(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, "invalid request")
	}
	log.Println("Accepting registration token:", req.Token)
	return c.JSON(200, "ok")
}