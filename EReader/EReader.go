package ereader

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"strings"
)

const (
	meta_container_path = "META-INF/container.xml"
)

type EReader struct {
	files   map[string]*zip.File
	chapter []string
	dir     string
	Package
}

type Package struct {
	Container *Container
	Content   *Content
	Nav       *Nav
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
	for _, i := range self.Package.Content.Items {
		if i.ID == "nav" {
			nav_path = i.Href
		}
	}

	err = self.setNav(self.dir + "/" + nav_path)
	if err != nil {
		return
	}
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
