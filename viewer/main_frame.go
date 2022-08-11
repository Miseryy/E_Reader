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
	main_frame := tview.NewPages()
	main_frame.SetBackgroundColor(tcell.ColorBlack)
	b_list := newBookList()
	main_frame.AddPage(p_book_list_name, b_list.makeFrame(), true, true)

	// main_frame.SetRows(0, 0).SetColumns(100, 0)
	// main_frame.AddItem(b_list.makeFrame(), 0, 0, 1, 1, 0, 0, true)

	return main_frame

}
