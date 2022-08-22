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
		pages.SwitchToPage(p_read_frame_name)

	})

}

func (t *tableOfContents) refleshTreeView() {
	t.frame.RemoveItem(read_book_ele.table_contents)
	t.makeTreeView()
	t.frame.AddItem(read_book_ele.table_contents, 0, 1, true)

}
