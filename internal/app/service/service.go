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

func (s *Service) FindFiles(word string) (foundFiles []string, err error) {
	// Контрольный перевод в lowercase
	w := strings.ToLower(word)
	search := searcher.New()
	foundFiles, err = search.Search(w)
	if err != nil {
		return nil, err
	}
	return foundFiles, nil // Должен возвращаться слайс файлов
}
