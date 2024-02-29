package main

import (
	"log"
	"word-search-in-files/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	err = a.Start()
	if err != nil {
		log.Fatal(err)
	}

}
