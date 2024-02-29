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

func (e *Endpoint) Root(c echo.Context) error {
	// r := c.Response() // Достать слово из респонса
	r := "" // Заглушка
	f := e.s.FindFile(r)

	err := c.String(http.StatusOK, f) // Переделать на JSON
	if err != nil {
		return err
	}

	return nil
}
