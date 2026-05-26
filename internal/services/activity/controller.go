package activity

import (
	"context"
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"strconv"

	"github.com/labstack/echo/v4"
)


type Controller struct {
	service ActivityService
}

func NewActivityController(service ActivityService) *Controller {
	return &Controller{service: service}
}

type ActivityController interface {
	ControllerGetActivityList(c echo.Context) (*ActivityRowList, error)
	ControllerGetActivityListRecycled(c echo.Context) (*ActivityRowList, error)
	ControllerInsertActivityExpense(c echo.Context)  error
	ControllerUpdateActivityExpense(c echo.Context)  error
	ControllerGetActivityById(c echo.Context) (*ActivityRow, error)
	ControllerDeleteActivitySoft(c echo.Context) error
	ControllerDeleteActivityErase(c echo.Context) error
	CreateActivity(ctx context.Context, activity ActivityRow) error
	UpdateActivity(ctx context.Context, activity ActivityRow) (*ActivityRow, error)
	GetActivityDropdownData(c echo.Context) (*ActivityDropdownData, error)
}

func (ctr *Controller) GetActivityDropdownData(c echo.Context) (*ActivityDropdownData, error) {
	advisor, ctx := auth.GetAdvisorClaims(c)
	
	advisor.Log("controller_attached_get_activity_dropdown_data")

	data, err := ctr.service.GetActivityDropdownData(ctx)
	if err != nil {
		advisor.Error("failed to get activity dropdown data", err)
		return nil, err
	}
	
	return data, nil
}


func (ctr *Controller) UpdateActivity(ctx context.Context, activity ActivityRow) (*ActivityRow, error) {
	advisor := advisor.FromContext(ctx)
	advisor.Log("controller_update_activity")
	
	if err := activity.ValidateForCreate(); err != nil {
		advisor.Error("activity validation failed", err)
		return nil, err
	}

	updatedActivity, err := ctr.service.UpdateActivity(ctx, activity)
	if err != nil {
		advisor.Error("failed to update activity", err)
		return nil, err
	}

	return updatedActivity, nil
}

func (ctr *Controller) CreateActivity(ctx context.Context, activity ActivityRow) error {
	advisor := advisor.FromContext(ctx)
	advisor.Log("controller_create_activity")

	if err := activity.ValidateForCreate(); err != nil {
		advisor.Error("activity validation failed", err)
		return err
	}

	if err := ctr.service.CreateActivity(ctx, activity); err != nil {
		advisor.Error("failed to create activity", err)
		return err
	}

	return nil
}

func (ctr *Controller) ControllerDeleteActivityErase(c echo.Context) error {
	advisor, ctx := auth.GetAdvisorClaims(c)
	var updatedActivity ActivityDeleteRequest
	if err := c.Bind(&updatedActivity); err != nil {
		advisor.Error("failed to bind request body to ActivityDeleteRequest", err)
		return err
	}
	advisor.Log("=================")
	advisor.Log("controller_attached_delete_activity_erase: " + strconv.Itoa(updatedActivity.ID))
	advisor.Log("=================")
	err := ctr.service.DeleteActivityErase(ctx, updatedActivity.ID)
	if err != nil {
		advisor.Error("failed to delete activity erase", err)
		return err
	}
	
	return nil
}


func (ctr *Controller) ControllerDeleteActivitySoft(c echo.Context) error {
	advisor, ctx := auth.GetAdvisorClaims(c)
	var updatedActivity ExpenseActivityRow
	if err := c.Bind(&updatedActivity); err != nil {
		advisor.Error("failed to bind request body to ExpenseActivityRow", err)
		return err
	}
	
	err := ctr.service.DeleteActivitySoft(ctx, updatedActivity.ID)
	if err != nil {
		advisor.Error("failed to delete activity soft", err)
		return err
	}
	
	return nil
}

func (ctr *Controller) ControllerUpdateActivityExpense(c echo.Context) error {
	advisor, ctx := auth.GetAdvisorClaims(c)
	
	advisor.Log("controller_attached_update_activity_expense")
	var updatedActivity ExpenseActivityRow
	if err := c.Bind(&updatedActivity); err != nil {
		advisor.Error("failed to bind request body to ExpenseActivityRow", err)
		return err
	}
	
	err := ctr.service.UpdateActivityExpense(ctx, updatedActivity)
	if err != nil {
		advisor.Error("failed to update activity expense", err)
		return err
	}
	
	return nil
}

func (ctr *Controller) ControllerGetActivityById(c echo.Context) (*ActivityRow, error) {
	advisor, ctx := auth.GetAdvisorClaims(c)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		advisor.Error("invalid id parameter", err)
		return nil, err
	}

	advisor.Log("controller_attached_get_activity_list_by_id"+ strconv.Itoa(id))

	row, err := ctr.service.GetActivityById(ctx, id)
	if err != nil {
		advisor.Error("failed to get activity list by id", err)
		return nil, err
	}
	
	return row, nil
}	

func (ctr *Controller) ControllerInsertActivityExpense(c echo.Context) error {
	advisor, ctx := auth.GetAdvisorClaims(c)

	advisor.Log("controller_attached_insert_activity_expense")
	
	
	var newActivity ExpenseActivityRow
	if err := c.Bind(&newActivity); err != nil {
		advisor.Error("failed to bind request body to ExpenseActivityRow", err)
		return err
	}
	
	err := ctr.service.InsertActivityExpense(ctx, newActivity)
	if err != nil {
		advisor.Error("failed to insert activity expense", err)
		return err
	}
	
	return nil
}



func (ctr *Controller) ControllerGetActivityList(c echo.Context) (*ActivityRowList, error) {
	advisor, ctx := auth.GetAdvisorClaims(c)

	advisor.Log("controller_attached_get_activity_list")

	activities, err := ctr.service.GetActivityList(ctx)
	if err != nil {
		advisor.Error("failed to get activity list", err)
		return nil, err
	}

	return &ActivityRowList{
		Activities: activities,
	}, nil
}

func (ctr *Controller) ControllerGetActivityListRecycled(c echo.Context) (*ActivityRowList, error) {
	advisor, ctx := auth.GetAdvisorClaims(c)

	advisor.Log("controller_attached_get_activity_list_recycled")

	activities, err := ctr.service.GetActivityListRecycled(ctx)
	if err != nil {
		advisor.Error("failed to get activity list recycled", err)
		return nil, err
	}

	return &ActivityRowList{
		Activities: activities,
	}, nil
}