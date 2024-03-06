package mw

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

// Инициализация слова из последнего запроса
var lastWord string

func WordCheck(cch cache.Cache) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// Слово из текущего запроса
			w := c.QueryParam("word")
			if w == "" || w != lastWord {
				lastWord = w
				err := next(c)
				if err != nil {
					fmt.Println("server error")
					return err
				}
				fmt.Println("response from server")
				return nil
			}

			// Возвращаем предыдущий запрос
			if cachedResponse, found := cch.Get(w); found {
				err := c.JSONBlob(http.StatusOK, cachedResponse.([]byte))
				if err != nil {
					fmt.Println("server error")
					return err
				}
				fmt.Println("response from cache")
			}
			return nil
		}
	}
}
