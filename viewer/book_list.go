package viewer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	ereader "epub_test/e-reader"

	"github.com/gdamore/tcell/v2"
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

		// r, err := regexp.MatchString(".epub", info.Name())
		sp := strings.Split(info.Name(), ".")

		if err != nil {
			return err
		}

		if !info.IsDir() && sp[len(sp)-1] == "epub" {
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
	table := tview.NewTable()
	fmt.Println("fff")

	provisional_dir := "/home/owner/go/src/e_reader/epubs/"

	book_paths, err := _getEpubPaths(provisional_dir)

	readers := []*ereader.EReader{}

	for _, bpath := range book_paths {
		reader := ereader.New()
		reader.OpenEpub(bpath)
		err := reader.InitContainer()
		if err != nil {
			continue
		}

		readers = append(readers, reader)
	}

	if err != nil {
		panic(err)
	}

	head := strings.Split("Title,Creater,Publisher,Date,Language", ",")
	rows := len(readers) + 1 // +1 head
	cols := len(head)

	for r := 0; r < rows; r++ {
		var param_list []string
		if r > 0 {
			param_list = append(param_list, readers[r-1].GetContent().Metadata.Title)
			param_list = append(param_list, readers[r-1].GetContent().Metadata.Creator)
			param_list = append(param_list, readers[r-1].GetContent().Metadata.Publisher)
			param_list = append(param_list, readers[r-1].GetContent().Metadata.Date)
			param_list = append(param_list, readers[r-1].GetContent().Metadata.Language)
		}

		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			if r == 0 {
				color = tcell.ColorYellow
				// Header
				table.SetCell(r, c,
					tview.NewTableCell(head[c]).SetTextColor(color).SetAlign(tview.AlignCenter))
				continue
			}

			table.SetCell(r, c,
				tview.NewTableCell(param_list[c]).
					SetTextColor(color).SetAlign(tview.AlignLeft))
		}
	}

	return table
}
