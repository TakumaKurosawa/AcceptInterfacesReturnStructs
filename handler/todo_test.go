package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TakumaKurosawa/accept-interfaces-returns-structs/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockTodoUseCase struct {
	CreateTodoFunc func(todo *model.Todo) (*model.Todo, error)
}

func (m *MockTodoUseCase) CreateTodo(todo *model.Todo) (*model.Todo, error) {
	return m.CreateTodoFunc(todo)
}

func TestHandler_CreateTodo(t *testing.T) {
	tests := []struct {
		name       string
		input      createInput
		mockReturn *model.Todo
		mockError  error
		wantStatus int
	}{
		{
			name: "success",
			input: createInput{
				Title: "Test Todo",
			},
			mockReturn: &model.Todo{
				ID:    "1",
				Title: "Test Todo",
			},
			mockError:  nil,
			wantStatus: http.StatusOK,
		},
		{
			name: "usecase error",
			input: createInput{
				Title: "Test Todo",
			},
			mockReturn: nil,
			mockError:  errors.New("hoge"), // Assume errSomeError is defined elsewhere
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			reqBody, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockUseCase := &MockTodoUseCase{
				CreateTodoFunc: func(todo *model.Todo) (*model.Todo, error) {
					return tt.mockReturn, tt.mockError
				},
			}

			h := NewHandler(mockUseCase)
			if assert.NoError(t, h.CreateTodo(c)) {
				assert.Equal(t, tt.wantStatus, rec.Code)
			}
		})
	}
}
