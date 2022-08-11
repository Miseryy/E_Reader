package viewer

import (
	"os"

	"github.com/rivo/tview"
)

var (
	current_dir, _ = os.Getwd()
)

func Run() {
	app := tview.NewApplication()
	main_f := newMainFrame(app)

	if err := app.SetRoot(main_f.MakeFrame(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
