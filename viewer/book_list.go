package viewer

import (
	"os"
	"path/filepath"
	"strings"

	ereader "epub_test/e-reader"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type bookList struct {
	table   *tview.Table
	readers []*ereader.EReader
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

func (b *bookList) makeFrame() tview.Primitive {
	frame := tview.NewGrid()
	command_texts := tview.NewTextView()
	b.table = tview.NewTable()
	b.table.SetSelectable(true, false)

	b.table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyTab {
		}

		if key == tcell.KeyEnter {
		}

	}).SetSelectedFunc(func(row int, colm int) {
		if row < 1 {
			return
		}

		e_reader = b.readers[row]
		read_book_ele.text_view.Clear()

		pages.SwitchToPage(p_read_frame_name)
		b.table.SetSelectable(true, false)

	})

	frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			pages.SwitchToPage(p_read_frame_name)

		}
		return event
	})

	command_texts.SetDynamicColors(true).SetRegions(true)
	command_string := "[red]<Tab>[white]::GoToRead"
	command_texts.SetText(command_string)

	frame.SetRows(0, 1).SetColumns(0)
	frame.AddItem(b.table, 0, 0, 1, 1, 0, 0, true)
	frame.AddItem(command_texts, 1, 0, 1, 1, 0, 0, true)

	return frame
}

func (b *bookList) makeList() {
	provisional_dir := "/home/owner/go/src/e_reader/epubs/"
	book_paths, err := _getEpubPaths(provisional_dir)
	b.readers = []*ereader.EReader{}

	for _, bpath := range book_paths {
		reader := ereader.New()
		reader.OpenEpub(bpath)
		err := reader.InitContainer()
		if err != nil {
			continue
		}

		b.readers = append(b.readers, reader)
	}

	if err != nil {
		panic(err)
	}

	head := strings.Split("Title,Creater,Publisher,Date,Language", ",")
	rows := len(b.readers) + 1 // +1 head
	cols := len(head)

	for r := 0; r < rows; r++ {
		var param_list []string
		if r > 0 {
			param_list = append(param_list, b.readers[r-1].GetContent().Metadata.Title)
			param_list = append(param_list, b.readers[r-1].GetContent().Metadata.Creator)
			param_list = append(param_list, b.readers[r-1].GetContent().Metadata.Publisher)
			param_list = append(param_list, b.readers[r-1].GetContent().Metadata.Date)
			param_list = append(param_list, b.readers[r-1].GetContent().Metadata.Language)
		}

		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			if r == 0 {
				color = tcell.ColorYellow
				// Header
				b.table.SetCell(r, c,
					tview.NewTableCell(head[c]).SetTextColor(color).SetAlign(tview.AlignCenter))
				continue
			}

			b.table.SetCell(r, c,
				tview.NewTableCell(param_list[c]).
					SetTextColor(color).SetAlign(tview.AlignLeft))
		}
	}

	app.SetFocus(b.table)

}
