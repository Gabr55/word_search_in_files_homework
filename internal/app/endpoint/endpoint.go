package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	FindFile(word string) string
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Find(c echo.Context) error {
	p := c.QueryParam("word") // Достаем значение параетра word
	if p == "" {
		p = "" // Добавить условие пустого параметра word
	}
	f := e.s.FindFile(p)

	err := c.String(http.StatusOK, f) // Переделать на JSON
	if err != nil {
		return err
	}

	return nil
}
