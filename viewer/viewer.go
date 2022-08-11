package viewer

import (
	ereader "epub_test/e-reader"
	"os"

	"github.com/rivo/tview"
)

var (
	current_dir, _ = os.Getwd()
	e_reader       = ereader.New()
	pages          = &tview.Pages{}
	app            = &tview.Application{}
)

func Run() {
	app = tview.NewApplication()
	main_f := newMainFrame(app)

	if err := app.SetRoot(main_f.MakeFrame(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
