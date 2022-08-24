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
	p_toc_name        = "ToCPage"
)

func newMainFrame() *mainFrame {
	return &mainFrame{}
}

func (m mainFrame) MakeFrame() tview.Primitive {
	pages = tview.NewPages()
	frames = &view_frames{}
	read_book_ele = &read_book_element{}
	pages.SetBackgroundColor(tcell.ColorBlack)
	frame_objects.book_list = newBookList()
	frame_objects.read_book = newReadBook()
	frame_objects.toc = newToC()
	frames.book_list = frame_objects.book_list.makeFrame()
	frames.read_book = frame_objects.read_book.makeFrame()
	frames.toc = frame_objects.toc.makeFrame()

	frame_objects.book_list.makeList()

	pages.AddPage(p_book_list_name, frames.book_list, true, true)
	pages.AddPage(p_read_frame_name, frames.read_book, true, false)
	pages.AddPage(p_toc_name, frames.toc, true, false)

	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		name, _ := pages.GetFrontPage()
		switch event.Key() {
		case tcell.KeyTab:
			switch name {
			case p_read_frame_name:
				pages.SwitchToPage(p_book_list_name)
			case p_book_list_name:
				pages.SwitchToPage(p_read_frame_name)
			case p_toc_name:
				pages.SwitchToPage(p_read_frame_name)
			}

		case tcell.KeyEnter:
			switch name {
			case p_read_frame_name:

			}
		}

		switch event.Rune() {
		case 'p':
			pages.SwitchToPage(p_toc_name)
			app.SetFocus(frames.toc)

		case 'Q':
			app.Stop()
		}

		return event
	})

	return pages

}
