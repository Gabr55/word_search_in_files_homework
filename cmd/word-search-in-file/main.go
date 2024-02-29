package main

import (
	"log"
	"word-search-in-files/internal/pkg/app"
)

func main() {
	// Инициализируем приложение
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	// Запускаем приложение
	err = a.Start()
	if err != nil {
		log.Fatal(err)
	}

}
