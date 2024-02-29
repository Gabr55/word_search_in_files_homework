package app

import (
	"fmt"
	"word-search-in-files/internal/app/endpoint"
	"word-search-in-files/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	s    *service.Service
	e    *endpoint.Endpoint
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	// Маршруты
	a.echo.GET("/", a.e.Root) // Переделать на POST с передачей слова

	return a, nil
}

func (a *App) Start() error {
	err := a.echo.Start(":8000")
	if err != nil {
		return err
	}
	fmt.Println("server is started...")

	return nil

}
