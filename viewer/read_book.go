package viewer

import (
	"github.com/rivo/tview"
)

type readBook struct {
	frame *tview.Grid
}

func newReadBook() *readBook {
	return &readBook{}
}

func (r *readBook) makeFrame() tview.Primitive {
	r.frame = tview.NewGrid()
	read_book_ele.table_contents = tview.NewTreeView()
	read_book_ele.table_contents.SetBorder(true)

	read_book_ele.text_view = tview.NewTextView()
	read_book_ele.text_view.SetBorder(true)
	read_book_ele.text_view.SetScrollable(true)

	command_text_view := tview.NewTextView()
	command_text_view.SetDynamicColors(true).SetRegions(true)
	command_string := "[red]<Tab>[white]::GoToBookList"
	command_text_view.SetText(command_string)

	r.frame.SetRows(0, 1).SetColumns(0)
	r.frame.AddItem(read_book_ele.text_view, 0, 0, 1, 2, 0, 0, false)
	// under
	r.frame.AddItem(command_text_view, 1, 0, 1, 2, 0, 0, false)

	return r.frame
}
