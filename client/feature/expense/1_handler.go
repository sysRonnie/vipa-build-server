package expense

import (
	"go-tailwind-test/client/ui"
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
	g.GET("/expense/modal-edit-record/:viewType/:rowId", h.GetModalEditRecord)
	g.GET("/expense/modal-remove-record/:viewType/:rowId", h.GetModalRemoveRecord)
	g.POST("/expense/update-record/:viewType/:rowId", h.HandleUpdateRecord)
	g.POST("/expense/remove-record/:viewType/:rowId", h.HandleRemoveRecord)

}



func (h *Handler) HandleRemoveRecord(c echo.Context) error {
	rowId, err := strconv.Atoi(c.Param("rowId"))
	if err != nil {
		log.Println("rowId error", err)
		
		return c.JSON(400, "invalid rowId")
	}
	viewType, ok := ParseViewType(c.Param("viewType"))
	if !ok {
		log.Println("viewType is not okay")
		log.Println(c.Param("viewType"))
		return c.JSON(400, "invalid rowId")
	}

	err = h.service.ProcessRemoveRecordRequest(viewType, rowId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	log.Println("Here are the params for the handleRemoveRecord", viewType, rowId)

	return c.JSON(200, "ok")
}

func (h *Handler) GetModalRemoveRecord(c echo.Context) error {
	rowId, err := strconv.Atoi(c.Param("rowId"))
	if err != nil {
		log.Println("rowId error", err)
		
		return c.JSON(400, "invalid rowId")
	}
	viewType, ok := ParseViewType(c.Param("viewType"))
	if !ok {
		log.Println("viewType is not okay")
		log.Println(c.Param("viewType"))
		return c.JSON(400, "invalid rowId")
	}	


	row, err := h.service.FetchExpenseTableRow(viewType, rowId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return render(c, ModalRemoveClient(row))
}

func (h *Handler) HandleUpdateRecord(c echo.Context) error {
	rowId, err := strconv.Atoi(c.Param("rowId"))
	if err != nil {
		log.Println("rowId error", err)
		
		return c.JSON(400, "invalid rowId")
	}
	viewType, ok := ParseViewType(c.Param("viewType"))
	if !ok {
		log.Println("viewType is not okay")
		log.Println(c.Param("viewType"))
		return c.JSON(400, "invalid rowId")
	}
	log.Println("Here are the params for the handleEditRecord", viewType, rowId)

	return c.JSON(200, "ok")



}

func (h *Handler) GetModalEditRecord(c echo.Context) error {
	rowId, err := strconv.Atoi(c.Param("rowId"))
	if err != nil {
		log.Println("rowId error", err)
		
		return c.JSON(400, "invalid rowId")
	}
	viewType, ok := ParseViewType(c.Param("viewType"))
	if !ok {
		log.Println("viewType is not okay")
		log.Println(c.Param("viewType"))
		return c.JSON(400, "invalid rowId")
	}
	log.Println("Here are the params for the handleEditRecord", viewType, rowId)

	row, err := h.service.FetchExpenseTableRow(viewType, rowId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return render(c, ModalEditClient(row))
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
		if dupErr, ok := err.(*DuplicateClientNameError); ok {
			return render(c, ui.SandboxErrorMessage(dupErr.Error()))
		} else {
			return render(c, ui.SandboxErrorMessage("Something went wrong htmx"))
		}
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