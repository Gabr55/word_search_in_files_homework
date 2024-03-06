package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
)

// Интерфейс основной службы
type Service interface {
	FindFiles(word string) ([]string, error)
}

// Создаем структуру обработчика
type Endpoint struct {
	s   Service
	cch cache.Cache
}

// Конструктор
func New(s Service, cch cache.Cache) *Endpoint {
	return &Endpoint{
		s:   s,
		cch: cch,
	}
}

// Хендлер поиска слова
func (e *Endpoint) Find(c echo.Context) error {
	// Создаем структуру ответа
	type WordFiles struct {
		Word  string   `json:"word"`
		Files []string `json:"files"`
	}
	// Инициализируем структуру
	wf := &WordFiles{
		Word:  c.QueryParam("word"),
		Files: []string{},
	}
	// Добавить проверку на пустой или отсутсвующий параметр
	if utf8.RuneCountInString(wf.Word) < 2 {
		err := c.String(http.StatusNotFound, "word is too short or empty") // Переделать на нормальный json ответ
		if err != nil {
			return err
		}
	}
	f, err := e.s.FindFiles(wf.Word)
	if err != nil {
		return err
	}
	// Проверка на наличие файлов
	if len(f) == 0 {
		fmt.Println("files not found")
		err = c.String(http.StatusNotFound, "files not found")
		if err != nil {
			return err
		}
		return nil
	}

	// Сохраняем список найденных файлов
	wf.Files = f

	// Приводим к json
	jsonBytes, err := json.Marshal(wf)
	if err != nil {
		return err
	}

	// Сохраняем в кэш
	e.cch.Set(wf.Word, jsonBytes, cache.DefaultExpiration)

	// Отправляем ответ
	err = c.JSONPretty(http.StatusOK, wf, "")
	if err != nil {
		return err
	}

	return nil
}
