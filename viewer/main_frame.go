package viewer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type mainFrame struct {
	app *tview.Application
}

func newMainFrame(app *tview.Application) *mainFrame {
	return &mainFrame{app: app}
}

func (m mainFrame) MakeFrame() tview.Primitive {
	main_frame := tview.NewGrid()
	main_frame.SetBackgroundColor(tcell.ColorBlack)
	b_list := newBookList()
	main_frame.SetRows(30, 0).SetColumns(0, 0)
	main_frame.AddItem(b_list.makeFrame(), 0, 0, 1, 1, 0, 0, true)

	return main_frame

}
