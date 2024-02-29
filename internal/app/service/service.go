package service

import (
	"word-search-in-files/pkg/searcher"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) FindFile(word string) string {
	search := searcher.New()
	search.Search(word)
	return ""
}
