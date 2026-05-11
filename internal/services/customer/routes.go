package customer

import (
	"go-tailwind-test/internal/services/auth"
	"go-tailwind-test/internal/util/advisor"
	"go-tailwind-test/internal/util/network"
	"log"

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
	g.GET("/customer-list", h.GetCustomerListScreen, auth.Middleware)
	g.POST("/customer-insert", h.InsertCustomer, auth.Middleware)
}

func (h *Handler) InsertCustomer(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())	
	advisor.Log("advisor successfully attached to the request context in the InsertCustomer handler")
	var req CustomerInsertRequest
	if err := c.Bind(&req); err != nil {
		advisor.Log("failed to bind request body to CustomerInsertRequest struct: " + err.Error())
		return network.Fail(c, network.SandboxResponse{
			StatusCode: 400,
			Message: "invalid request body",
		})
	}

	if err := c.Validate(&req); err != nil {
		advisor.Log("validation failed for CustomerInsertRequest: " + err.Error())
		return network.Fail(c, network.SandboxResponse{
			StatusCode: 400,
			Message: "validation failed: " + err.Error(),
		})
	}

	log.Println("customer name received:", req.Name)

	err := h.store.InsertCustomer(c.Request().Context(), req.Name, req.Phone, req.Email, req.Comment)
	if err != nil {
		advisor.Log("failed to insert customer into the database: " + err.Error())
		return network.Fail(c, network.SandboxResponse{
			StatusCode: 500,
			Message: "failed to insert customer",
		})
	}
	
	
	return network.Success(c, network.SandboxResponse{
		StatusCode: 200,
		Message: "ok",
	})
}


func (h *Handler) GetCustomerListScreen(c echo.Context) error {
	advisor := advisor.FromContext(c.Request().Context())	
	advisor.Log("advisor successfully attached to the request context in the GetCustomerListScreen handler")

	customerList, err := h.store.QueryCustomerList(c.Request().Context())
	if err != nil {
		advisor.Log("failed to query customer list from the database: " + err.Error())
		return network.Fail(c, network.SandboxResponse{
			StatusCode: 500,
			Message: "failed to query customer list",
		})
	}

	response := CustomerListResponse{
		Customers: customerList,
	}


	return network.Success(c, network.SandboxResponse{
		StatusCode: 200,
		Message: "ok",
		Data: response,
	})
}