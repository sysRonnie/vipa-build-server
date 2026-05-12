package api

import (
	"database/sql"
	"go-tailwind-test/internal/services/customer"
	"go-tailwind-test/internal/services/project"
	"go-tailwind-test/internal/services/user"
	"go-tailwind-test/internal/services/vendor"

	"github.com/labstack/echo/v4"
)


type API struct {
	db *sql.DB
}

func NewAPIServer(db *sql.DB) *API {
	return &API {
		db: db,
	}
}

func (s *API) APIService(e *echo.Echo) error {
	v1 := e.Group("/api/v1")

	userStore := user.NewUserStore(s.db)
	userService := user.NewService(userStore)
	userHandler := user.NewHandler(userService, userStore)
	userHandler.RegisterUserRoutes(v1)

	customerStore := customer.NewCustomerStore(s.db)
	customerService := customer.NewCustomerService(customerStore)
	customerHandler := customer.NewCustomerHandler(customerService, customerStore)
	customerHandler.RegisterCustomerRoutes(v1)

	projectStore := project.NewProjectStore(s.db)
	projectService := project.NewProjectService(projectStore)
	projectHandler := project.NewProjectHandler(projectService, projectStore)
	projectHandler.RegisterProjectRoutes(v1)

	vendorStore := vendor.NewVendorStore(s.db)
	vendorService := vendor.NewVendorService(vendorStore)
	vendorHandler := vendor.NewVendorHandler(vendorService, vendorStore)
	vendorHandler.RegisterVendorRoutes(v1)
	
	return nil
}