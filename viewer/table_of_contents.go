package viewer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type tableOfContents struct {
	frame *tview.Flex
}

func newToC() *tableOfContents {
	return &tableOfContents{}
}

func (t *tableOfContents) makeFrame() tview.Primitive {
	t.frame = tview.NewFlex()
	t.frame.SetDirection(tview.FlexColumn)
	t.frame.SetBorder(true)
	t.frame.SetTitle("TOC")
	return t.frame

}

func (t *tableOfContents) makeTreeView() {
	e_reader.MakeChapters()
	nav := e_reader.GetNav()
	root := tview.NewTreeNode(nav.Title).SetColor(tcell.ColorRed).SetReference("title")

	read_book_ele.table_contents = tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	for i, d := range e_reader.GetToCs() {
		n := tview.NewTreeNode(d.ChapterName).SetReference(i).SetSelectable(true)
		root.AddChild(n)
	}

	read_book_ele.table_contents.SetSelectedFunc(func(node *tview.TreeNode) {
		read_book_ele.text_view.Clear()
		var no int
		switch val := node.GetReference().(type) {
		case int:
			no = val
		default:
			return
		}

		e_reader.TocSetIte(no)
		toc := e_reader.TocNext()

		text, e := e_reader.GetChapterText(toc.ChapterPath)
		if e != nil {
			read_book_ele.text_view.SetText(e.Error())
			return
		}

		read_book_ele.text_view.SetText(text)
		read_book_ele.text_view.ScrollToBeginning()
		pages.SwitchToPage(p_read_frame_name)
		app.SetFocus(read_book_ele.text_view)

	})

}

func (t *tableOfContents) refleshTreeView() {
	t.frame.RemoveItem(read_book_ele.table_contents)
	t.makeTreeView()
	t.frame.AddItem(read_book_ele.table_contents, 0, 1, true)

}
