package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/TakumaKurosawa/accept-interfaces-returns-structs/handler"
	"github.com/TakumaKurosawa/accept-interfaces-returns-structs/pkg/uid"
	"github.com/TakumaKurosawa/accept-interfaces-returns-structs/usecase"
)

func main() {
	todoUseCase := usecase.NewTodoUseCase(uid.NewGenerator())
	todoHandler := handler.NewHandler(todoUseCase)

	e := echo.New()
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/todos", todoHandler.CreateTodo)

	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
