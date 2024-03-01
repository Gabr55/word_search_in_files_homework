package searcher

import (
	"bufio"
	"io/fs"
	"strings"
	"word-search-in-files/pkg/internal/dir"
)

type Searcher struct {
	FS fs.FS
}

func New() *Searcher {
	return &Searcher{}
}

func (s *Searcher) Search(word string) (files []string, err error) {
	fileNames, err := dir.FilesFS(s.FS, "./examples") // Проблема в нем

	if err != nil {
		return nil, err
	}
	// Проверка наличия файлов
	if fileNames == nil {
		return nil, nil
	}
	// Алоцируем слайс размером с количесво файлов
	files, err = make([]string, 0, len(fileNames)), nil
	w := strings.ToLower(word)
	for _, f := range fileNames {
		file, err := s.FS.Open(f)
		if err != nil {
			return nil, err
		}
		// Освобождаем ресурс после отработки функции
		defer file.Close()
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			line := scan.Text()
			if strings.Contains(strings.ToLower(line), w) { // Найти как получить слово игорируя только знаки. Например Word1 это другое слово
				files = append(files, strings.Split(f, ".")[0]) // Найти как получить имя без расширения
			}
		}

	}
	if err != nil {
		return nil, err
	}
	return files, nil
}
