package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/TakumaKurosawa/accept-interfaces-returns-structs/model"
)

type TodoUseCase interface {
	CreateTodo(todo *model.Todo) (*model.Todo, error)
}

type Handler struct {
	tu TodoUseCase
}

func NewHandler(tu TodoUseCase) *Handler {
	return &Handler{
		tu: tu,
	}
}

type createInput struct {
	Title string `json:"title"`
}

func (h *Handler) CreateTodo(ctx echo.Context) error {
	var input createInput
	if err := ctx.Bind(&input); err != nil {
		return err
	}

	created, err := h.tu.CreateTodo(&model.Todo{
		Title: input.Title,
	})
	if err != nil {
		return err
	}

	if err := ctx.JSON(http.StatusOK, created); err != nil {
		return err
	}

	return nil
}
