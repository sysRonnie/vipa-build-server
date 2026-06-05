package vendor

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service VendorService
	store   VendorStore
}

func NewVendorHandler(service VendorService, store VendorStore) *Handler {
	return &Handler{
		service: service,
		store:   store,
	}
}

func (h *Handler) RegisterVendorRoutes(g *echo.Group) {
	g.POST("/vendor-create", h.InsertVendor, auth.Middleware)
	g.GET("/vendor-read", h.GetVendorList, auth.Middleware)	
	g.GET("/vendor-read-names", h.GetVendorListNames, auth.Middleware)
	g.GET("/vendor-read-by-id/:id", h.GetVendorByID, auth.Middleware)
	g.GET("/vendor-read-recycled", h.GetVendorListRecycled, auth.Middleware)
	g.POST("/vendor-update", h.UpdateVendor, auth.Middleware)
	g.POST("/vendor-delete", h.DeleteVendor, auth.Middleware)
	g.GET("/vendor-name-latest", h.GetVendorNameLatest, auth.Middleware) // New route for latest vendor name
}

func (h *Handler) GetVendorNameLatest(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get latest vendor name request")
	
	vendorName, err := h.store.QueryVendorNameLatest(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query latest vendor name from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, vendorName)
}

func (h *Handler) GetVendorListNames(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get vendor list names request")

	vendorNames, err := h.store.QueryVendorListNames(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query vendor list names from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}

	advisor.Log("Successfully queried vendor list names from the database: " + strconv.Itoa(len(vendorNames.Names)) + " names found")

	return network.BuildSuccessResponse(c, vendorNames)
}


func (h *Handler) GetVendorByID(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get vendor by ID request")
	
	idParam := c.Param("id")
	vendorID, err := strconv.Atoi(idParam)
	if err != nil {
		advisor.Error("invalid vendor ID parameter: ", err)
		return network.FailFromError(c, network.ErrInvalidRequest)
	}
	
	vendor, err := h.store.QueryVendorByID(c.Request().Context(), vendorID)
	if err != nil {
		advisor.Error("failed to query vendor by ID from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, vendor)
}

func (h *Handler) DeleteVendor(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing delete vendor request")
	
	var req struct {
		ID int `json:"id" validate:"required"`
	}
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.DeleteVendor(c.Request().Context(), req.ID)
	if err != nil {
		advisor.Error("failed to delete vendor in the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) UpdateVendor(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing update vendor request")
	
	var req VendorRow
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.UpdateVendor(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to update vendor in the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) InsertVendor(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing insert vendor request")
	
	var req VendorRow
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.InsertVendor(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to insert vendor into the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) GetVendorList(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get vendor list request")
	vendorList, err := h.store.QueryVendorList(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query vendor list from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, VendorRowList{Vendors: vendorList})
}

func (h *Handler) GetVendorListRecycled(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing get recycled vendor list request")
	vendorList, err := h.store.QueryVendorListRecycled(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query recycled vendor list from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, VendorRowList{Vendors: vendorList})
}

