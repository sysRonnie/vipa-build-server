package note

import (
	"go-tailwind-test/internal/util/network"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetNoteList(c echo.Context) error {
	res, err := h.controller.GetNoteList(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, res)
}

func (h *Handler) GetNoteListRecycled(c echo.Context) error {
	res, err := h.controller.GetNoteListRecycled(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, res)
}

func (h *Handler) GetNoteByID(c echo.Context) error {
	res, err := h.controller.GetNoteByID(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, res)
}

func (h *Handler) CreateNote(c echo.Context) error {
	err := h.controller.CreateNote(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, nil)
}

func (h *Handler) UpdateNote(c echo.Context) error {
	err := h.controller.UpdateNote(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, nil)
}

func (h *Handler) RemoveNote(c echo.Context) error {
	err := h.controller.RemoveNote(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, nil)
}

func (h *Handler) EraseNote(c echo.Context) error {
	err := h.controller.EraseNote(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, nil)
}