package searcher

import (
	"io/fs"
)

type Searcher struct {
	FS fs.FS
}

func New() *Searcher {
	return &Searcher{}
}

func (s *Searcher) Search(word string) (files []string, err error) {
	//dir.FilesFS(s.FS, ".")?
	return
}
