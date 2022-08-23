package viewer

import (
	ereader "epub_test/e-reader"
	"os"

	"github.com/rivo/tview"
)

type read_book_element struct {
	text_view      *tview.TextView
	table_contents *tview.TreeView
}

type view_frames struct {
	read_book tview.Primitive
	book_list tview.Primitive
	toc       tview.Primitive
}

type frameObjects struct {
	read_book *readBook
	book_list *bookList
	toc       *tableOfContents
}

var (
	current_dir, _    = os.Getwd()
	e_reader          = &ereader.EReader{}
	read_book_ele     = &read_book_element{}
	pages             = &tview.Pages{}
	app               = &tview.Application{}
	frame_objects     = &frameObjects{}
	frames            = &view_frames{}
	chapter_path_list = []string{}
)

func Run() {
	app = tview.NewApplication()
	main_f := newMainFrame()

	if err := app.SetRoot(main_f.MakeFrame(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
