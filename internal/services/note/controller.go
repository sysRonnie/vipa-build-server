package note

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service NoteService
}

func NewNoteController(service NoteService) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) GetNoteList(
	ctx echo.Context,
) (*NoteRowList, error) {

	return c.service.GetNoteList(
		ctx.Request().Context(),
	)
}

func (c *Controller) GetNoteListRecycled(
	ctx echo.Context,
) (*NoteRowList, error) {

	return c.service.GetNoteListRecycled(
		ctx.Request().Context(),
	)
}

func (c *Controller) GetNoteByID(
	ctx echo.Context,
) (*NoteRow, error) {

	noteID, err := strconv.Atoi(
		ctx.Param("id"),
	)

	if err != nil {
		return nil, ErrNoteIDRequired
	}

	return c.service.GetNoteByID(
		ctx.Request().Context(),
		noteID,
	)
}

func (c *Controller) CreateNote(
	ctx echo.Context,
) error {

	var note NoteRow

	

	err := ctx.Bind(&note)
	if err != nil {
		return err
	}



	if note.NoteBody == "" {
		return ErrNoteBodyRequired
	}

	return c.service.CreateNote(
		ctx.Request().Context(),
		note,
	)
}

func (c *Controller) UpdateNote(
	ctx echo.Context,
) error {

	var note NoteRow

	err := ctx.Bind(&note)
	if err != nil {
		return err
	}

	if note.ID == 0 {
		return ErrNoteIDRequired
	}

	if note.NoteBody == "" {
		return ErrNoteBodyRequired
	}

	return c.service.UpdateNote(
		ctx.Request().Context(),
		note,
	)
}

func (c *Controller) RemoveNote(
	ctx echo.Context,
) error {

	var req NoteDeleteRequest

	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	if req.ID == 0 {
		return ErrNoteIDRequired
	}

	return c.service.RemoveNote(
		ctx.Request().Context(),
		req.ID,
	)
}

func (c *Controller) EraseNote(
	ctx echo.Context,
) error {

	var req NoteDeleteRequest

	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	if req.ID == 0 {
		return ErrNoteIDRequired
	}

	return c.service.EraseNote(
		ctx.Request().Context(),
		req.ID,
	)
}