package viewer

import "github.com/rivo/tview"

type mainFrame struct {
	app *tview.Application
}

func newMainFrame(app *tview.Application) *mainFrame {
	return &mainFrame{app: app}
}

func (m mainFrame) MakeFrame() tview.Primitive {
	main_frame := tview.NewGrid()
	main_frame.SetBorder(true)
	return main_frame

}
