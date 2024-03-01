package endpoint

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	FindFiles(word string) ([]string, error)
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
	// Добавить проверку на пустой или отсутсвующий параметр
	if p == "" {
		err := c.String(http.StatusOK, "Параметры не найдены") // Переделать на нормальный json ответ
		if err != nil {
			return err
		}
	}
	f, err := e.s.FindFiles(p)
	if err != nil {
		return err
	}
	err = c.JSON(http.StatusOK, f) // Переделать на JSON с полем founded
	if err != nil {
		return err
	}

	return nil
}
