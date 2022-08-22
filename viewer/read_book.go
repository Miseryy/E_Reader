package viewer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type readBook struct {
	frame *tview.Grid
}

func newReadBook() *readBook {
	return &readBook{}
}

func (r *readBook) makeFrame() tview.Primitive {
	r.frame = tview.NewGrid()
	read_book_ele.table_contents = tview.NewTreeView()
	read_book_ele.table_contents.SetBorder(true)

	read_book_ele.text_view = tview.NewTextView()
	read_book_ele.text_view.SetText("sdfafa").SetBorder(true)
	read_book_ele.text_view.SetScrollable(true)

	command_text_view := tview.NewTextView()
	command_text_view.SetDynamicColors(true).SetRegions(true)
	command_string := "[red]<Tab>[white]::GoToBookList"
	command_text_view.SetText(command_string)

	r.frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			pages.SwitchToPage(p_book_list_name)

		}
		return event
	})

	r.frame.SetRows(0, 1).SetColumns(0, 0)
	r.frame.AddItem(read_book_ele.table_contents, 0, 0, 1, 1, 0, 0, true)
	r.frame.AddItem(read_book_ele.text_view, 0, 1, 1, 1, 0, 0, false)
	// under
	r.frame.AddItem(command_text_view, 1, 0, 1, 2, 0, 0, false)

	return r.frame
}

func (r *readBook) makeTreeView() {
	e_reader.MakeChapters()
	chaps := e_reader.GetChapters()
	_ = chaps
	read_book_ele.table_contents.SetBorder(true)
	nav := e_reader.GetNav()

	root := tview.NewTreeNode(nav.Title).SetColor(tcell.ColorRed).SetReference("ref")

	read_book_ele.table_contents = tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	// for _, d := range nav.Nav {
	// 	for _, l1 := range d.Li {
	// 		n := tview.NewTreeNode(l1.A.Href).SetReference(l1.A.Href).SetSelectable(true)
	// 		root.AddChild(n)

	// 	}
	// }

	for _, d := range e_reader.GetTableOfContents() {
		n := tview.NewTreeNode(d.ChapterName).SetReference(d.ChapterPath).SetSelectable(true)
		root.AddChild(n)
	}

	read_book_ele.table_contents.SetSelectedFunc(func(node *tview.TreeNode) {
		read_book_ele.text_view.Clear()
		ref := node.GetReference().(string)

		text, e := e_reader.GetChapterText(ref)
		if e != nil {
			read_book_ele.text_view.SetText(e.Error())
			return
		}

		read_book_ele.text_view.SetText(text)
		app.SetFocus(read_book_ele.text_view)
		// n := tview.NewTreeNode("test2").SetReference("ref").SetSelectable(true)
		// root.AddChild(n)
	})

}

func (r *readBook) refleshTreeView() {
	r.frame.RemoveItem(read_book_ele.table_contents)
	r.makeTreeView()
	r.frame.AddItem(read_book_ele.table_contents, 0, 0, 1, 1, 0, 0, true)
	app.SetFocus(read_book_ele.table_contents)

}

func (r *readBook) refleshTextview() {

}
