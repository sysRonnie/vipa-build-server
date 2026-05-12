package project

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"log"
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
	g.GET("/project-read", h.GetProjectList, auth.Middleware)
	g.GET("/project-read-by-id/:id", h.GetProjectByID, auth.Middleware)
	g.POST("/project-create", h.InsertProject, auth.Middleware)
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

	log.Println(" ----- JSON STRUCT ------")
	log.Println("req.Name", req.Name)
	log.Println("req.CustomerName", req.CustomerName)
	log.Println("req.StartDate", req.StartDate)
	log.Println("req.EndDateEst", req.EndDateEst)
	log.Println("req.EndDateActual", req.EndDateActual)
	log.Println("req.IsDeleted", req.IsDeleted)
	log.Println("req.CreatedAt", req.CreatedAt)
	log.Println("req.UpdatedAt", req.UpdatedAt)
	
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
	projects, err := h.store.QueryProjectList(c.Request().Context())
	if err != nil {
		log.Println("Error querying project list: ", err)
		return network.FailFromError(c, err)
	}

	return network.BuildSuccessResponse(c, ProjectRowList{
		Projects: projects,
	})
}



