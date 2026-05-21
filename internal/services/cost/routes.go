package cost

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service CostService
	store  CostStore
}

func NewCostHandler(service CostService, store CostStore) *Handler {
	return &Handler{service: service, store: store}
}

func (h *Handler) RegisterCostRoutes(g *echo.Group) {
	g.GET("/cost-read", h.GetCostList, auth.Middleware)
	g.GET("/cost-read-names", h.GetCostListNames, auth.Middleware)
	g.GET("/cost-read-recycled", h.GetCostListRecycled, auth.Middleware)
	g.GET("/cost-read-by-id/:id", h.GetCostByID, auth.Middleware)
	g.POST("/cost-create", h.InsertCost, auth.Middleware)
	g.POST("/cost-update", h.UpdateCost, auth.Middleware)
	g.POST("/cost-delete", h.DeleteCost, auth.Middleware)
}

func (h *Handler) GetCostListNames(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get cost list names request")
	
	costNames, err := h.store.QueryCostListNames(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query cost list names from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}


	return network.BuildSuccessResponse(c, costNames)
}

func (h *Handler) GetCostByID(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get vendor by ID request")
	
	idParam := c.Param("id")
	vendorID, err := strconv.Atoi(idParam)
	if err != nil {
		advisor.Error("invalid vendor ID parameter: ", err)
		return network.FailFromError(c, network.ErrInvalidRequest)
	}
	
	vendor, err := h.store.QueryCostByID(c.Request().Context(), vendorID)
	if err != nil {
		advisor.Error("failed to query cost by ID from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, vendor)
}

func (h *Handler) DeleteCost(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing delete cost request")
	
	var req struct {
		ID int `json:"id" validate:"required"`
	}
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.DeleteCost(c.Request().Context(), req.ID)
	if err != nil {
		advisor.Error("failed to delete cost in the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) UpdateCost(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing update cost request")
	
	var req CostRow
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.UpdateCost(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to update cost in the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) InsertCost(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing insert vendor request")
	
	var req CostRow
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.InsertCost(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to insert cost into the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) GetCostList(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get cost list request")
	costList, err := h.store.QueryCostList(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query cost list from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, CostRowList{Costs: costList})
}

func (h *Handler) GetCostListRecycled(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get recycled cost list request")
	costList, err := h.store.QueryCostListRecycled(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query recycled cost list from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, CostRowList{Costs: costList})
}

