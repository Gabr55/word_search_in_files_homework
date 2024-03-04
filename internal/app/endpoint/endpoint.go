package endpoint

import (
	"fmt"
	"net/http"
	"unicode/utf8"

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
	if utf8.RuneCountInString(p) < 2 {
		err := c.String(http.StatusNotFound, "word is too short or empty") // Переделать на нормальный json ответ
		if err != nil {
			return err
		}
	}
	f, err := e.s.FindFiles(p)
	if err != nil {
		return err
	}
	if len(f) == 0 {
		fmt.Println("files not found")
		err = c.String(http.StatusNotFound, "files not found")
		if err != nil {
			return err
		}
		return nil
	}
	err = c.JSON(http.StatusOK, f)
	if err != nil {
		return err
	}

	return nil
}
