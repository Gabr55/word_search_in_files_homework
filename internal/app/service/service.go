package service

import (
	"strings"
	"word-search-in-files/pkg/searcher"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) FindFile(word string) string {
	w := strings.ToLower(word)
	search := searcher.New()
	search.Search(w)
	return w // Должно возвращаться имя файла
}
