package viewer

import (
	"encoding/json"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type readBook struct {
	frame *tview.Grid
}

type bookSaveParam struct {
	Title string
	Pos   map[string]int
}

func newReadBook() *readBook {
	return &readBook{}
}

func (r *readBook) setTextView() {
	toc := e_reader.GetCurrentContents()

	text, e := e_reader.GetChapterText(toc.ChapterPath)
	if e != nil {
		read_book_ele.text_view.SetText(e.Error())
		return
	}

	read_book_ele.text_view.SetText(text)
	pages.SwitchToPage(p_read_frame_name)
	app.SetFocus(read_book_ele.text_view)
}

func (r *readBook) nextPage() {
	toc := e_reader.TocNext()
	if toc == nil {
		return
	}

	text, e := e_reader.GetChapterText(toc.ChapterPath)
	if e != nil {
		read_book_ele.text_view.SetText(e.Error())
	}

	read_book_ele.text_view.SetText(text)
	read_book_ele.text_view.ScrollToBeginning()

}

func (r *readBook) beforePage() {
	toc := e_reader.TocPrev()
	if toc == nil {
		return
	}

	text, e := e_reader.GetChapterText(toc.ChapterPath)
	if e != nil {
		read_book_ele.text_view.SetText(e.Error())
	}

	read_book_ele.text_view.SetText(text)
	read_book_ele.text_view.ScrollToBeginning()

}

func (r *readBook) saveCurrentPosJson() {
	title := e_reader.GetContent().Metadata.Title
	file, err := os.Create(current_dir + "/" + title + ".json")

	if err != nil {
		return
	}

	defer file.Close()

	row, colm := read_book_ele.text_view.GetScrollOffset()
	idx := e_reader.GetIdx()
	bsp := bookSaveParam{
		Title: title,
		Pos:   map[string]int{"row": row, "colm": colm, "idx": idx},
	}

	err = json.NewEncoder(file).Encode(bsp)
	if err != nil {
		panic(err)
	}
}

func (r *readBook) loadCurrentPosJson() {
	title := e_reader.GetContent().Metadata.Title
	file, err := os.Open(current_dir + "/" + title + ".json")
	if err != nil {
		return
	}

	json_data := bookSaveParam{}

	err = json.NewDecoder(file).Decode(&json_data)

	if err != nil {
		return
	}

	idx := json_data.Pos["idx"]
	row := json_data.Pos["row"]
	colm := json_data.Pos["colm"]

	read_book_ele.text_view.ScrollTo(row, colm)
	e_reader.TocSetIte(idx)

}

func (r *readBook) makeFrame() tview.Primitive {
	r.frame = tview.NewGrid()
	read_book_ele.table_contents = tview.NewTreeView()
	read_book_ele.table_contents.SetBorder(true)

	read_book_ele.text_view = tview.NewTextView()
	read_book_ele.text_view.SetBorder(true)
	read_book_ele.text_view.SetScrollable(true)

	command_text_view := tview.NewTextView()
	command_text_view.SetDynamicColors(true).SetRegions(true)
	command_string := "[red]<Tab>[white]::GoToBookList [red]<p>[white]::TableOfContents [red]<n>[white]::NextPage [red]<b>[white]::PrevPage [red]<s>[white]::Save"
	command_text_view.SetText(command_string)

	r.frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'n':
			r.nextPage()
		case 'b':
			r.beforePage()
		case 's':
			r.saveCurrentPosJson()
		}

		return event
	})

	r.frame.SetRows(0, 1).SetColumns(0)
	r.frame.AddItem(read_book_ele.text_view, 0, 0, 1, 2, 0, 0, false)
	// under
	r.frame.AddItem(command_text_view, 1, 0, 1, 2, 0, 0, false)

	return r.frame
}
