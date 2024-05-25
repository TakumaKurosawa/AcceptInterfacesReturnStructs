package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/TakumaKurosawa/accept-interfaces-returns-structs/model"
	"github.com/TakumaKurosawa/accept-interfaces-returns-structs/usecase"
)

type Handler struct {
	tu usecase.TodoUseCase
}

func NewHandler(tu usecase.TodoUseCase) *Handler {
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
