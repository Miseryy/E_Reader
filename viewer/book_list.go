package viewer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/rivo/tview"
)

type bookList struct {
}

func newBookList() *bookList {
	return &bookList{}
}

func _getEpubPaths(root string) ([]string, error) {
	var book_paths []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		r, err := regexp.MatchString("epub", info.Name())
		if err != nil {
			return err
		}

		if !info.IsDir() && r {
			book_paths = append(book_paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return book_paths, nil
}

func (s bookList) makeFrame() tview.Primitive {
	list_frame := tview.NewTable()

	provisional_dir := "/home/owner/go/src/e_reader/epubs/"

	book_paths, err := _getEpubPaths(provisional_dir)

	if err != nil {
		panic(err)
	}

	fmt.Println(book_paths)

	rows := len(book_paths)
	cols := 5

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

		}
	}

	return list_frame
}
