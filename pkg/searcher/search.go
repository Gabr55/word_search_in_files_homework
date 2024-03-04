package searcher

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"
	"word-search-in-files/pkg/internal/dir"
)

type Searcher struct {
	FS fs.FS
}

// Конструктор поисковика
func New() *Searcher {
	return &Searcher{}
}

// Структура слова и подходящего списка файлов
type WordFiles struct {
	word string
	file string
}

func (s *Searcher) Search(word string) (files []string, err error) {
	if s.FS == nil {
		return nil, errors.New("nil file system")
	}
	// Проверка наличия cлова для поиска
	if utf8.RuneCountInString(word) < 2 {
		return nil, errors.New("word is too short or empty")
	}
	// Инициалицаия файловой системы
	fileNames, err := dir.FilesFS(s.FS, "")
	if err != nil {
		return nil, err
	}
	// Иницаилизируем слайс c набором файлов
	files = make([]string, 0, len(fileNames))
	w := strings.ToLower(word)

	// Иницаилизируем WaitGroup для всех горутин
	var wg sync.WaitGroup
	// Инициализируем буферизированный канал с результатами WordFiles
	r := make(chan WordFiles, len(fileNames))

	// Запускаем горутины для каждого файл
	for _, f := range fileNames {
		wg.Add(1)
		go func(f string) error {
			defer wg.Done()
			file, err := s.FS.Open(f)
			if err != nil {
				return err
			}
			// Освобождаем ресурс после отработки функции
			defer file.Close()
			// Создаем текстовый буфер
			scan := bufio.NewScanner(file)
			for scan.Scan() {
				line := strings.Split(scan.Text(), " ")
				// Проверяем наличие слова
				for _, i := range line {

					if strings.EqualFold(strings.ToLower(
						strings.Trim(i, ".,!?-\n\t")), w) {
						r <- WordFiles{word: w, file: strings.Split(f, ".")[0]}
						break
					}
				}
			}
			if err := scan.Err(); err != nil {
				fmt.Println("Error reading file:", err)
			}
			return nil
		}(f)
	}
	// Ждем завершения всех горутин
	wg.Wait()
	close(r)
	for v := range r {
		fmt.Println(v)
		files = append(files, v.file)
	}
	// Финальная сортировка
	sort.Strings(files)
	return files, nil
}
