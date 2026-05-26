package note

import (
	"go-tailwind-test/internal/services/auth"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	controller NoteController
}

func NewNoteHandler(controller NoteController) *Handler {
	return &Handler{
		controller: controller,
	}
}

func (h *Handler) RegisterNoteRoutes(g *echo.Group) {
	g.GET("/note-read", h.GetNoteList, auth.Middleware)
	g.GET("/note-read-recycled", h.GetNoteListRecycled, auth.Middleware)
	g.GET("/note-read-by-id/:id", h.GetNoteByID, auth.Middleware)

	g.POST("/note-create", h.CreateNote, auth.Middleware)
	g.POST("/note-update", h.UpdateNote, auth.Middleware)

	g.POST("/note-remove", h.RemoveNote, auth.Middleware)
	g.POST("/note-erase", h.EraseNote, auth.Middleware)
}