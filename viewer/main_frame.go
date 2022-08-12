package viewer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type mainFrame struct {
}

const (
	p_book_list_name  = "BookPage"
	p_read_frame_name = "ReadPage"
)

func newMainFrame() *mainFrame {
	return &mainFrame{}
}

func (m mainFrame) MakeFrame() tview.Primitive {
	pages = tview.NewPages()
	frames := &view_frames{}
	read_book_ele = &read_book_element{}
	pages.SetBackgroundColor(tcell.ColorBlack)
	frame_objects.book_list = newBookList()
	frame_objects.read_book = newReadBook()
	frames.book_list = frame_objects.book_list.makeFrame()
	frames.read_book = frame_objects.read_book.makeFrame()

	frame_objects.book_list.makeList()

	// app.SetFocus(book_list)
	pages.AddPage(p_book_list_name, frames.book_list, true, true)
	pages.AddPage(p_read_frame_name, frames.read_book, true, false)

	// main_frame.SetRows(0, 0).SetColumns(100, 0)
	// main_frame.AddItem(b_list.makeFrame(), 0, 0, 1, 1, 0, 0, true)

	return pages

}
