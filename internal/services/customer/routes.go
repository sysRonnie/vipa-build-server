package customer

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)


type Handler struct {
	service CustomerService
	store CustomerStore
}


func NewCustomerHandler(service CustomerService, store CustomerStore) *Handler {
	return &Handler{
		service: service,
		store: store,
	}
}

func (h *Handler) RegisterCustomerRoutes(g *echo.Group) {
	g.POST("/customer-create", h.InsertCustomer, auth.Middleware)
	g.GET("/customer-read", h.GetCustomerList, auth.Middleware)
	g.POST("/customer-update", h.UpdateCustomer, auth.Middleware) 
	g.POST("/customer-delete", h.DeleteCustomer, auth.Middleware) 

	g.GET("/customer-names", h.GetCustomerNames, auth.Middleware)
	g.GET("/customer-name-latest", h.GetCustomerNamesLatest, auth.Middleware) // New route for latest customer names
	g.GET("/customer-detail/:id", h.GetCustomerDetailScreen, auth.Middleware)
	g.GET("/customer-read-recycled", h.GetCustomerListRecycled, auth.Middleware)
}

func (h *Handler) GetCustomerNamesLatest(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())	
	advisor.Log("advisor successfully attached to the request context in the GetCustomerNamesLatest handler")
	
	customerName, err := h.store.QueryCustomerNameLatest(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query latest customer names from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, customerName)
}

func (h *Handler) GetCustomerNames(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())	
	advisor.Log("advisor successfully attached to the request context in the GetCustomerNames handler")
	
	customerNames, err := h.store.QueryCustomerNames(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query customer names from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponse(c, customerNames)
}

func (h *Handler) DeleteCustomer(c echo.Context) error {
	// For simplicity, we will just mark the customer as deleted in the database instead of actually deleting the record
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("Processing delete customer request")
	
	var req DeleteCustomerRequest
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.store.DeleteCustomer(c.Request().Context(), req.ID)
	if err != nil {
		advisor.Error("failed to delete customer in the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	
	return network.BuildSuccessResponseOK(c)
}

func (h *Handler) UpdateCustomer(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())	
	advisor.Log("advisor successfully attached to the request context in the UpdateCustomer handler")
	var req CustomerInsertRequest
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}
	
	err := h.service.ProcessUpdateCustomer(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to update customer in the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.BuildSuccessResponseOK(c)
}	


func (h *Handler) GetCustomerDetailScreen(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())
	advisor.Log("customer_detail_request_started")

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		advisor.Error("invalid_customer_id_route_param", err)
		return network.FailFromError(c, network.ErrInvalidRequest)
	}

	customer, err := h.store.QueryCustomerById(c.Request().Context(), id)

	if err != nil {
		advisor.Error("failed_to_query_customer_by_id", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}

	return network.BuildSuccessResponse(c, customer)
}

func (h *Handler) InsertCustomer(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())	
	advisor.Log("advisor successfully attached to the request context in the InsertCustomer handler")
	var req CustomerInsertRequest
	if err := network.BindAndValidate(c, &req); err != nil {
		advisor.Error("failed to bind and validate request body: ", err)
		return network.FailFromError(c, err)
	}


	err := h.service.ProcessCreateCustomer(c.Request().Context(), req)
	if err != nil {
		advisor.Error("failed to insert customer into the database: ", err)
		return network.FailFromError(c, err)
	}
	
	return network.Success(c, network.SuccessResponseOK)	
}

func (h *Handler) GetCustomerListRecycled(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())	
	advisor.Log("advisor successfully attached to the request context in the GetCustomerListRecycledScreen handler")
	customerList, err := h.store.QueryCustomerListRecycled(c.Request().Context())
	if err != nil {
		advisor.Error("failed to query customer list from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}
	response := CustomerListResponse{Customers: customerList}

	return network.BuildSuccessResponse(c, response)
}

func (h *Handler) GetCustomerList(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())	
	advisor.Log("advisor successfully attached to the request context in the GetCustomerListScreen handler")

	customerList, err := h.store.QueryCustomerList(c.Request().Context())

	if err != nil {
		log.Println("error internal")
		advisor.Error("failed to query customer list from the database: ", err)
		return network.FailFromError(c, network.ErrDatabaseFailure)
	}

	response := CustomerListResponse{Customers: customerList}

	return network.BuildSuccessResponse(c, response)
}