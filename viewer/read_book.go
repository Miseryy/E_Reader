package viewer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type readBook struct {
}

func newReadBook() *readBook {
	return &readBook{}
}

func (r *readBook) makeFrame() tview.Primitive {
	frame := tview.NewGrid()

	root := tview.NewTreeNode("test").SetColor(tcell.ColorRed).SetReference("ref")

	read_book_ele.table_contents = tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	read_book_ele.table_contents.SetSelectedFunc(func(node *tview.TreeNode) {
		n := tview.NewTreeNode("test2").SetReference("ref").SetSelectable(true)

		root.AddChild(n)

	}).SetBorder(true)

	frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			pages.SwitchToPage(p_book_list_name)

		}
		return event
	})

	read_book_ele.text_view = tview.NewTextView()
	read_book_ele.text_view.SetText("sdfafa").SetBorder(true)

	command_text_view := tview.NewTextView()
	command_text_view.SetDynamicColors(true).SetRegions(true)
	command_string := "[red]<Tab>[white]::GoToBookList"
	command_text_view.SetText(command_string)

	frame.SetRows(0, 1).SetColumns(0, 0)
	frame.AddItem(read_book_ele.table_contents, 0, 0, 1, 1, 0, 0, true)
	frame.AddItem(read_book_ele.text_view, 0, 1, 1, 1, 0, 0, true)
	frame.AddItem(command_text_view, 1, 0, 1, 2, 0, 0, true)

	return frame
}

func (r *readBook) refleshTreeView() {

}

func (r *readBook) refleshTextview() {

}
