package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	p := "ид"
	fmt.Println(utf8.RuneCountInString(p))
}
