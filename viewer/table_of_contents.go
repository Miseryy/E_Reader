package viewer

import (
	ereader "epub_reader/e-reader"

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

func (t *tableOfContents) noTocParam(root *tview.TreeNode) {
	content := e_reader.GetContent()
	items := content.Items
	for i, d := range items {
		if d.MediaType != "application/xhtml+xml" {
			continue
		}
		n := tview.NewTreeNode(d.ID).SetReference(i).SetSelectable(true)
		toc := ereader.TableOfContents{}
		toc.ChapterName = d.ID
		toc.ChapterPath = d.Href
		e_reader.TableOfContents = append(e_reader.TableOfContents, &toc)
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
		frame_objects.read_book.setTextView()

	})
}

func (t *tableOfContents) makeTreeView() {
	e_reader.MakeChapters()

	title := e_reader.GetContent().Metadata.Title
	root := tview.NewTreeNode(title).SetColor(tcell.ColorRed).SetReference("title")

	read_book_ele.table_contents = tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	tocs := e_reader.GetToCs()

	if len(tocs) == 0 {
		t.noTocParam(root)
		return
	}

	for i, d := range tocs {
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
		frame_objects.read_book.setTextView()
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
