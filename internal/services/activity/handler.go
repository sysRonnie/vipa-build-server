package activity

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/network"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	controller ActivityController
	service    ActivityService
	store      ActivityStore
}

func NewActivityHandler(controller ActivityController, service ActivityService, store ActivityStore) *Handler {
	return &Handler{controller: controller, service: service, store: store}
}

func (h *Handler) RegisterActivityRoutes(g *echo.Group) {
	g.GET("/activity-read", h.GetActivityList, auth.Middleware)
	g.GET("/activity-read-recycled", h.GetActivityListRecycled, auth.Middleware)
	g.GET("/activity-read-by-id/:id", h.GetActivityByID, auth.Middleware)

	g.POST("/activity-create", h.CreateActivity, auth.Middleware)
	g.POST("/activity-update", h.UpdateActivity, auth.Middleware)

	g.POST("/activity-remove", h.RemoveActivitySoft, auth.Middleware)
	g.POST("/activity-erase", h.RemoveActivityErase, auth.Middleware)

	g.GET("/activity-dropdown-data", h.GetActivityDropdownData, auth.Middleware)
}

func (h *Handler) GetActivityDropdownData(c echo.Context) error {
	res, err := h.controller.GetActivityDropdownData(c)
	if err != nil {
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponse(c, res)
}

func (h *Handler) UpdateActivity(c echo.Context) error {
	advisor, ctx := auth.GetAdvisorClaims(c)
	var req ActivityRow
	if err := c.Bind(&req); err != nil {
		advisor.Error("failed to bind create activity request", err)
		return network.FailFromError(c, err)
	}
	updatedActivity, err := h.controller.UpdateActivity(ctx, req)
	if err != nil {
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponse(c, updatedActivity)
}

func (h *Handler) CreateActivity(c echo.Context) error {
	advisor, ctx := auth.GetAdvisorClaims(c)

	var req ActivityRow
	if err := c.Bind(&req); err != nil {
		advisor.Error("failed to bind create activity request", err)
		return network.FailFromError(c, err)
	}

	if err := h.controller.CreateActivity(ctx, req); err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, nil)
}


func (h *Handler) RemoveActivitySoft(c echo.Context) error {
	err := h.controller.ControllerDeleteActivitySoft(c)
	if err != nil {
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponse(c, nil)
}

func (h *Handler) RemoveActivityErase(c echo.Context) error {
	err := h.controller.ControllerDeleteActivityErase(c)
	if err != nil {
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponse(c, nil)
}

func (h *Handler) UpdateActivityExpense(c echo.Context) error {
	err := h.controller.ControllerUpdateActivityExpense(c)
	if err != nil {
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponse(c, nil)
}

func (h *Handler) GetActivityByID(c echo.Context) error {
	res, err := h.controller.ControllerGetActivityById(c)
	if err != nil {
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponse(c, res)
}


func (h *Handler) InsertActivityExpense(c echo.Context) error {
	err := h.controller.ControllerInsertActivityExpense(c)
	if err != nil {
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponse(c, nil)
}

func (h *Handler) GetActivityList(c echo.Context) error {
	res, err := h.controller.ControllerGetActivityList(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, res)
}

func (h *Handler) GetActivityListRecycled(c echo.Context) error {
	res, err := h.controller.ControllerGetActivityListRecycled(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, res)
}