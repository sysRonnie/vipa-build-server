package project

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service ProjectService
	store ProjectStore
}

func NewProjectHandler(service ProjectService, store ProjectStore) *Handler{
	return &Handler{
		service: service,
		store: store,
	}
}


func (h *Handler) RegisterProjectRoutes(g *echo.Group) {
	g.POST("/project-create", h.InsertProject, auth.Middleware)
	g.GET("/project-read", h.GetProjectList, auth.Middleware)
	g.GET("/project-read-recycled", h.GetProjectListRecycled, auth.Middleware)
	g.GET("/project-read-by-id/:id", h.GetProjectByID, auth.Middleware)
	g.POST("/project-update", h.UpdateProject, auth.Middleware)
	g.POST("/project-delete", h.DeleteProject, auth.Middleware)
}


func (h *Handler) DeleteProject(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing delete project request")
	
	var req struct {
		ID int `json:"id" validate:"required"`
	}
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.DeleteProject(c.Request().Context(), req.ID)
	if err != nil {
		advisor.Error("failed to delete project in the database: ", err)
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponseOK(c)

}


func (h *Handler) UpdateProject(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing update project request")
	var req ProjectRow
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}

	advisor.Log(" ----- JSON STRUCT ------")
	advisor.Log("req.ID" + strconv.Itoa(req.ID))
	advisor.Log("req.Name"+  req.Name)
	advisor.Log("req.CustomerName"+  req.CustomerName)
	advisor.Log("req.StartDate"+  req.StartDate)
	advisor.Log("req.EndDateEst"+  req.EndDateEst)
	advisor.Log("req.EndDateActual"+  req.EndDateActual)
	advisor.Log("req.IsDeleted"+  strconv.FormatBool(req.IsDeleted))
	advisor.Log("req.CreatedAt"+  strconv.FormatInt(req.CreatedAt.Unix(), 10))
	advisor.Log("req.UpdatedAt"+  strconv.FormatInt(req.UpdatedAt.Unix(), 10))
	
	err := h.store.UpdateProject(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to update project in the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) GetProjectByID(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get project by ID request")
	
	idParam := c.Param("id")
	if idParam == "" {
		advisor.Error("missing project ID in request path", network.ErrInvalidRequest)
		return network.FailFromError(c, network.ErrInvalidRequest)
	}
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		advisor.Error("invalid project ID format: ", err)
		return network.FailFromError(c, network.ErrInvalidRequest)
	}
	
	project, err := h.store.QueryProjectByID(c.Request().Context(), id)
	if err != nil {
		advisor.Error("failed to query project by ID from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, project)
}

func (h *Handler) InsertProject(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing insert project request")
	var req ProjectRow
	if err := network.BindAndValidate(c, &req); err != nil {
		return network.FailFromError(c, err)
	}

	advisor.Log(" ----- JSON STRUCT ------")
	advisor.Log("req.ID" + strconv.Itoa(req.ID))
	advisor.Log("req.Name"+  req.Name)
	advisor.Log("req.CustomerName"+  req.CustomerName)
	advisor.Log("req.StartDate"+  req.StartDate)
	advisor.Log("req.EndDateEst"+  req.EndDateEst)
	advisor.Log("req.EndDateActual"+  req.EndDateActual)
	advisor.Log("req.Price"+  strconv.FormatFloat(req.Price, 'f', 2, 64))
	advisor.Log("req.Budget"+  strconv.FormatFloat(req.Budget, 'f', 2, 64))
	advisor.Log("req.Note"+  req.Note)
	advisor.Log("req.IsDeleted"+  strconv.FormatBool(req.IsDeleted))
	advisor.Log("req.CreatedAt"+  strconv.FormatInt(req.CreatedAt.Unix(), 10))
	advisor.Log("req.UpdatedAt"+  strconv.FormatInt(req.UpdatedAt.Unix(), 10))
	
	
	err := h.store.InsertProject(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to insert project into the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) GetProjectList(c echo.Context) error {
	// For simplicity, we will just return an empty list of projects for now
	// In a real implementation, you would query the database for the list of projects and return that data
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get project list request")
	projects, err := h.store.QueryProjectList(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query project list from the database: ", err)
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, ProjectRowList{
		Projects: projects,
	})
}


func (h *Handler) GetProjectListRecycled(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get recycled project list request")
	projects, err := h.store.QueryProjectListRecycled(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query recycled project list from the database: ", err)
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, ProjectRowList{
		Projects: projects,
	})
}


