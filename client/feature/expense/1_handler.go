package expense

import (
	"log"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service ExpenseService
	store ExpenseStore
}

func NewHandler(service ExpenseService, store ExpenseStore) *Handler {
	return &Handler{
		service: service,
		store: store,
	}
}

func (h *Handler) RegisterExpenseRoutes(g *echo.Group) {
	g.GET("/expense/:viewType", h.ShowExpensePage)
	g.GET("/expense", h.ShowExpensePage)
	g.POST("/expense/insert-new-record", h.HandleInsertNewRecord)
	g.GET("/expense/edit-record/:rowId", h.HandleEditRecord)
	g.GET("/expense/update-record/:rowId", h.HandleEditRecord)
}

func (h *Handler) HandleUpdateRecord(c echo.Context) error {
	param, err := strconv.Atoi(c.Param("rowId"))
	if err != nil {
		return c.JSON(400, "invalid rowId")
	}

	log.Println("Updating Recording Hit")
	return nil



}

func (h *Handler) HandleEditRecord(c echo.Context) error {
	log.Println("HandleEditRecord")
	param, err := strconv.Atoi(c.Param("rowId"))
	if err != nil {
		return c.JSON(400, "invalid rowId")
	}

	row, err := h.service.FetchExpenseTableRow(ViewTypeClient, param)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return render(c, EditClientModal(row))
}

func (h *Handler) HandleInsertNewRecord(c echo.Context) error {
    log.Println("HandleInsertNewRecord")

    if err := c.Request().ParseForm(); err != nil {
        return err
    }

    raw := map[string][]string(c.Request().PostForm)

	viewType, ok := ParseViewType(c.FormValue("viewType"))
    if !ok {
        return c.JSON(400, "invalid viewType")
    }
	

    err := h.service.ProcessNewRecordInsertion(viewType, raw)
    if err != nil {
        return c.JSON(400, err.Error())
    }

    return c.JSON(200, "ok")
}
type ExpensePageProps struct {
    TableData ExpenseTableModel
    ViewType ExpenseViewType
}


type PageHeader struct {
    Title string
}
func (h *Handler) ShowExpensePage(c echo.Context) error {
    param := c.Param("viewType")

    viewType, ok := ParseViewType(param)
    if !ok {
        viewType = ViewTypeClient
    }

	viewType = ViewTypeClient

	
	tableModel, err := h.service.ProcessExpenseTablePayload(viewType)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	

    props := ExpensePageProps{
        TableData: tableModel,
        ViewType:  viewType,
    }

    return render(c, ExpensePage(props))
}

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}