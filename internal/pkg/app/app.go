package app

import (
	"fmt"
	"time"
	"word-search-in-files/internal/app/endpoint"
	"word-search-in-files/internal/app/mw"
	"word-search-in-files/internal/app/service"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

type App struct {
	s    *service.Service
	e    *endpoint.Endpoint
	echo *echo.Echo
	cch  *cache.Cache
}

func New() (*App, error) {
	a := &App{}

	// Инициализация нашей службы
	a.s = service.New()

	// Инициализация кэша
	a.cch = cache.New(5*time.Minute, 10*time.Minute)

	// Инициализация нашего эндпоинта status
	// Передаем службу и указатель на кэш
	a.e = endpoint.New(a.s, *a.cch)

	// Инициализация веб сервера
	a.echo = echo.New()

	// Для передачи middleware
	a.echo.Use(mw.WordCheck(*a.cch))
	// Маршруты
	a.echo.GET("/find", a.e.Find)

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
