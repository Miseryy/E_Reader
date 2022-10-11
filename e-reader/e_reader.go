package ereader

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	meta_container_path = "META-INF/container.xml"
	chapter_media_type  = "application/xhtml+xml"
)

type EReader struct {
	files           map[string]*zip.File
	chapter         []string
	navPath         []*navParam
	dir             string
	middle_dir      string
	file_path       string
	TableOfContents []*TableOfContents
	iterator        *ToCIterator
	pack
}

type navParam struct {
	name string
	path string
}

type TableOfContents struct {
	ChapterName string
	ChapterPath string
}

type pack struct {
	container *Container
	content   *Content
	nav       *Nav
	chapter   []*Chapter
}

type ToCIterator struct {
	e_reader *EReader
	idx      int
}

func (c *ToCIterator) HasNext() bool {
	return c.idx < c.e_reader.GetToCSize()
}

func (c *ToCIterator) HasPrev() bool {
	return c.idx >= 0
}

func (c *ToCIterator) Next() *TableOfContents {
	c.idx++
	if c.idx >= c.e_reader.GetToCSize()-1 {
		c.idx = c.e_reader.GetToCSize() - 1
	}
	item := c.e_reader.GetToCAt(c.idx)
	return item
}

func (c *ToCIterator) Prev() *TableOfContents {
	c.idx--
	if c.idx < 0 {
		c.idx = 0
	}
	item := c.e_reader.GetToCAt(c.idx)
	return item
}

func (c *ToCIterator) Get(idx int) *TableOfContents {
	if c.e_reader.GetToCSize() <= idx {
		return nil
	}
	c.idx = idx
	return c.e_reader.TableOfContents[c.idx]
}

func (c *ToCIterator) SetIte(idx int) {
	c.idx = idx
}

func New() *EReader {
	m := &EReader{
		iterator: &ToCIterator{},
	}

	m.iterator.e_reader = m

	return m
}

func (e *EReader) OpenEpub(file_path string) {
	f, err := os.Open(file_path)
	e.file_path = file_path

	if err != nil {
		fmt.Println(err)
		return
	}

	fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	z, err := zip.NewReader(f, fi.Size())
	if err != nil {
		fmt.Println(err)
		return
	}

	e.files = make(map[string]*zip.File)
	for _, ff := range z.File {
		e.files[ff.Name] = ff
	}

}

func (e *EReader) setContainer() error {
	b, err := e.openFile(meta_container_path)
	if err != nil {
		return err
	}

	container := new(Container)
	xml.Unmarshal(b, container)
	e.pack.container = container

	return nil
}

func (e *EReader) setContent(path string) error {
	b, err := e.openFile(path)
	if err != nil {
		return err
	}

	content := new(Content)
	xml.Unmarshal(b, content)
	e.pack.content = content

	return nil

}

func (e *EReader) setNav(path string) error {
	b, err := e.openFile(path)
	if err != nil {
		return err
	}

	nav := new(Nav)
	xml.Unmarshal(b, nav)
	e.pack.nav = nav

	return nil

}

func (e *EReader) InitContainer() error {
	err := e.setContainer()
	if err != nil {
		return err
	}

	full_path := e.pack.container.Rootfiles[0].Rootfile.FullPath
	err = e.setContent(full_path)
	if err != nil {
		return err
	}

	e.dir = strings.Split(full_path, "/")[0]

	return nil
}

func (e *EReader) MakeChapters() error {
	var nav_path string
	var chapter_path_href string
	for _, i := range e.pack.content.Items {
		if i.ID == "nav" || i.Properties == "nav" {
			nav_path = i.Href
		}

		if i.ID != "nav" && i.MediaType == chapter_media_type {
			chapter_path_href = i.Href
		}

		if len(nav_path) > 0 && len(chapter_path_href) > 0 {
			break
		}
	}

	e.middle_dir = strings.Split(chapter_path_href, "/")[0]

	if len(nav_path) < 1 {
		return errors.New("Can't open")
	}

	err := e.setNav(e.dir + "/" + nav_path)
	if err != nil {
		return err
	}

	for _, n := range e.pack.nav.Nav {
		if n.Type != "toc" {
			continue
		}

		for _, l := range n.Li {
			tmp_param := navParam{
				name: l.A.Text,
				path: l.A.Href,
			}
			e.navPath = append(e.navPath, &tmp_param)
			for _, ll := range l.Li {
				tmp_param := navParam{
					name: ll.A.Text,
					path: ll.A.Href,
				}
				e.navPath = append(e.navPath, &tmp_param)
			}
		}
	}

	for _, nav := range e.navPath {
		p := strings.Split(nav.path, "#")[0]
		sp_array := strings.Split(p, "/")

		if len(sp_array) > 1 {
			p = e.dir + "/" + p
		} else {
			p = e.dir + "/" + e.middle_dir + "/" + p

		}

		toc := TableOfContents{}
		toc.ChapterName = nav.name
		toc.ChapterPath = p

		e.TableOfContents = append(e.TableOfContents, &toc)

	}

	return nil
}

func (e *EReader) GetChapterText(path string) (string, error) {
	b, err := e.openFile(path)
	if err != nil {
		return "", err
	}

	ch := &Chapter{}
	xml.Unmarshal(b, ch)
	d := GetTagHead(ch.Body.Data)

	var data string
	for _, dd := range d {
		if len(dd.Data) == 0 {
			continue
		}

		switch dd.Tag {
		case "h1", "h2", "h3", "h4", "h5", "h6":
			data += fmt.Sprintf("\t%s\n", dd.Data)
		default:
			data += fmt.Sprintf("%s\n", dd.Data)
		}
	}

	return data, nil
}

func (e *EReader) setChapter(path string) error {
	b, err := e.openFile(path)
	if err != nil {
		return err
	}

	ch := &Chapter{}
	xml.Unmarshal(b, ch)
	e.pack.chapter = append(e.pack.chapter, ch)

	return nil
}

func (e EReader) GetToCSize() int {
	return len(e.TableOfContents)
}

func (e EReader) GetContainer() Container {
	return *e.pack.container
}

func (e EReader) GetContent() Content {
	return *e.pack.content
}

func (e EReader) GetNav() Nav {
	return *e.pack.nav
}

func (e *EReader) GetFilePath() string {
	return e.file_path
}

func (e *EReader) GetToCs() []*TableOfContents {
	return e.TableOfContents
}

func (e *EReader) GetToCAt(idx int) *TableOfContents {
	if idx < 0 {
		return nil
	}

	if idx >= e.GetToCSize() {
		return nil
	}

	return e.iterator.Get(idx)
}

func (e EReader) HasTocNext() bool {
	return e.iterator.HasNext()
}

func (e EReader) HasTocPrev() bool {
	return e.iterator.HasPrev()
}

func (e EReader) TocNext() *TableOfContents {
	return e.iterator.Next()
}

func (e EReader) TocPrev() *TableOfContents {
	return e.iterator.Prev()
}

func (e *EReader) TocSetIte(idx int) {
	e.iterator.SetIte(idx)
}

func (e EReader) GetIdx() int {
	return e.iterator.idx
}

func (e EReader) GetCurrentContents() *TableOfContents {
	return e.iterator.Get(e.iterator.idx)
}

func (e EReader) openFile(path string) ([]byte, error) {
	content_file, err := e.files[path].Open()
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	_, err = io.Copy(&b, content_file)
	return b.Bytes(), nil
}
