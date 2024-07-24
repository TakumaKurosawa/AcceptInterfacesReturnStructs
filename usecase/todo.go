package usecase

import (
	"github.com/TakumaKurosawa/accept-interfaces-returns-structs/model"
	"github.com/TakumaKurosawa/accept-interfaces-returns-structs/pkg/uid"
)

type todoUseCase struct {
	ug uid.Generator
}

func NewTodoUseCase(ug uid.Generator) *todoUseCase {
	return &todoUseCase{
		ug: ug,
	}
}

func (u *todoUseCase) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	var err error
	todo.ID, err = u.ug.NewUUIDV7()
	if err != nil {
		return nil, err
	}

	return todo, nil
}
