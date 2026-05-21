package event

import (
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service EventService
}

func NewEventController(service EventService) *Controller {
	return &Controller{service: service}
}

type EventController interface {
	CreateEventActivity(c echo.Context) (*network.EmptyResponse, error)
}

func (ctr *Controller) CreateEventActivity(c echo.Context) (*network.EmptyResponse, error) {
	ctx := c.Request().Context()


	advisor := advisor.FromContext(ctx)
	advisor.Log("processing_event_activity_create_request")
	var req EventActivityRow
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("bind_and_validate_error_on_request_body: ", err)
		return nil, err
	}

	if err := ctr.service.CreateEventActivity(ctx, req); err != nil {
		advisor.Error("create_event_activity_service_error: ", err)
		return nil, err
	}

	return &network.EmptyResponse{}, nil
}