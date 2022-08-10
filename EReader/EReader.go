package ereader

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
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
	files      map[string]*zip.File
	chapter    []string
	NavPath    []*navParam
	dir        string
	middle_dir string
	Package
}

type navParam struct {
	name string
	path string
}

type Package struct {
	Container *Container
	Content   *Content
	Nav       *Nav
	Chapter   []*Chapter
}

func New() *EReader {
	return &EReader{}
}

func (self *EReader) OpenEpub(file_path string) {
	f, err := os.Open(file_path)

	if err != nil {
		panic(err)
	}

	fi, err := f.Stat()
	if err != nil {
		return
	}

	z, err := zip.NewReader(f, fi.Size())
	if err != nil {
		return
	}

	self.files = make(map[string]*zip.File)
	for _, ff := range z.File {
		self.files[ff.Name] = ff
	}

	err = self.setContainer()
	if err != nil {
		return
	}

	full_path := self.Package.Container.Rootfiles[0].Rootfile.FullPath
	err = self.setContent(full_path)
	if err != nil {
		return
	}

	self.dir = strings.Split(full_path, "/")[0]

	var nav_path string
	var chapter_path_href string
	for _, i := range self.Package.Content.Items {
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

	self.middle_dir = strings.Split(chapter_path_href, "/")[0]

	if len(nav_path) < 1 {
		fmt.Println("error")
		return
	}

	err = self.setNav(self.dir + "/" + nav_path)
	if err != nil {
		return
	}

	for _, n := range self.Package.Nav.Nav {
		if n.Type != "toc" {
			continue
		}

		for _, l := range n.Li {
			tmp_param := navParam{
				name: l.A.Text,
				path: l.A.Href,
			}

			self.NavPath = append(self.NavPath, &tmp_param)
		}
	}

	for _, nav := range self.NavPath {
		p := strings.Split(nav.path, "#")[0]
		sp_array := strings.Split(p, "/")

		if len(sp_array) > 1 {
			p = self.dir + "/" + p
		} else {
			p = self.dir + "/" + self.middle_dir + "/" + p

		}
		self.setChapter(p)
	}

	// self.setChapter(path)

	// fmt.Println(path)
	fmt.Println(self.Package.Chapter[5].Body.Data)
	// // fmt.Println(self.Package.Chapter[0].Body.Data)
	// dd := strings.Split(self.Package.Chapter[0].Body.Data, "\n")
	// fmt.Println(dd)

	// fmt.Println(self.Package.Chapter[0].Title)

}

func (self *EReader) setContainer() error {
	b, err := self.openFile(meta_container_path)
	if err != nil {
		return err
	}

	container := new(Container)
	xml.Unmarshal(b, container)
	self.Package.Container = container

	return nil
}

func (self *EReader) setContent(path string) error {
	b, err := self.openFile(path)
	if err != nil {
		return err
	}

	content := new(Content)
	xml.Unmarshal(b, content)
	self.Package.Content = content

	return nil

}

func (self *EReader) setNav(path string) error {
	b, err := self.openFile(path)
	if err != nil {
		return err
	}

	nav := new(Nav)
	xml.Unmarshal(b, nav)
	self.Package.Nav = nav

	return nil

}

func (self *EReader) setChapter(path string) error {
	b, err := self.openFile(path)
	if err != nil {
		return err
	}

	ch := new(Chapter)
	xml.Unmarshal(b, ch)
	self.Package.Chapter = append(self.Package.Chapter, ch)

	return nil
}

func (self EReader) GetPackage() Package {
	return self.Package
}

func (self EReader) openFile(path string) ([]byte, error) {
	content_file, err := self.files[path].Open()
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	_, err = io.Copy(&b, content_file)
	return b.Bytes(), nil
}
