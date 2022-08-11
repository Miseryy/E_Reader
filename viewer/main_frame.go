package viewer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type mainFrame struct {
	app *tview.Application
}

const (
	p_book_list_name  = "BookPage"
	p_read_frame_name = "ReadPage"
)

func newMainFrame(app *tview.Application) *mainFrame {
	return &mainFrame{app: app}
}

func (m mainFrame) MakeFrame() tview.Primitive {
	pages = tview.NewPages()
	pages.SetBackgroundColor(tcell.ColorBlack)
	b_list := newBookList()
	t := tview.NewTextView()
	t.SetText("sfa")
	book_list := b_list.makeFrame()

	// app.SetFocus(book_list)
	pages.AddPage(p_book_list_name, book_list, true, false)
	pages.AddPage(p_read_frame_name, t, true, true)

	b_list.makeList()

	// main_frame.SetRows(0, 0).SetColumns(100, 0)
	// main_frame.AddItem(b_list.makeFrame(), 0, 0, 1, 1, 0, 0, true)

	return pages

}
