package event

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service EventService
	store  EventStore
	controller EventController
}

func NewEventHandler(service EventService, store EventStore, controller EventController) *Handler {
	return &Handler{service: service, store: store, controller: controller}
}

func (h *Handler) RegisterEventRoutes(g *echo.Group) {
	g.GET("/event-read", h.GetEventList, auth.Middleware)
	g.GET("/event-read-names", h.GetEventListNames, auth.Middleware)
	g.GET("/event-read-recycled", h.GetEventListRecycled, auth.Middleware)
	g.GET("/event-read-by-id/:id", h.GetEventByID, auth.Middleware)
	g.POST("/event-create", h.InsertEvent, auth.Middleware)
	g.POST("/event-update", h.UpdateEvent, auth.Middleware)
	g.POST("/event-delete", h.DeleteEvent, auth.Middleware)


	g.POST("/event-activity-create", h.InsertEventActivity, auth.Middleware)
}

func (h *Handler) InsertEventActivity(c echo.Context) error {
	res, err := h.controller.CreateEventActivity(c)
	if err != nil {
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, res)
}

func (h *Handler) GetEventListNames(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get event list names request")
	eventNames, err := h.store.QueryEventListNames(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query event list names from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}

	advisor.Log("Successfully queried event list names from the database: "+strconv.Itoa(len(eventNames))+" names found")

	return network.BuildSuccessResponse(c, EventNameList{Names: eventNames})
}

func (h *Handler) GetEventByID(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get event by ID request")
	
	idParam := c.Param("id")
	eventID, err := strconv.Atoi(idParam)
	if err != nil {
		advisor.Error("invalid event ID parameter: ", err)
		return network.FailFromError(c, network.ErrInvalidRequest)
	}
	
	event, err := h.store.QueryEventByID(c.Request().Context(), eventID)
	if err != nil {
		advisor.Error("failed to query event by ID from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, event)
}

func (h *Handler) DeleteEvent(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing delete event request")
	
	var req struct {
		ID int `json:"id" validate:"required"`
	}
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.DeleteEvent(c.Request().Context(), req.ID)
	if err != nil {
		advisor.Error("failed to delete event in the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) UpdateEvent(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing update event request")
	
	var req EventRow
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.UpdateEvent(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to update event in the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) InsertEvent(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing insert event request")
	
	var req EventRow
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.InsertEvent(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to insert event into the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) GetEventList(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get event list request")
	eventList, err := h.store.QueryEventList(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query event list from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, EventRowList{Events: eventList})
}

func (h *Handler) GetEventListRecycled(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get recycled event list request")
	eventList, err := h.store.QueryEventListRecycled(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query recycled event list from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, EventRowList{Events: eventList})
}


