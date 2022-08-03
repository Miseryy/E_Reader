package ereader

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

type EReader struct {
	files map[string]*zip.File
}

func (self EReader) OpenEpub(file_path string) {
	f, err := os.Open("./mybook.epub")

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

	files := make(map[string]*zip.File)
	for _, ff := range z.File {
		files[ff.Name] = ff
	}

}

func openFile(files map[string]*zip.File, path string) ([]byte, error) {
	content_file, err := files[path].Open()
	if err != nil {
		return nil, nil
	}
	var b bytes.Buffer
	_, err = io.Copy(&b, content_file)
	return b.Bytes(), nil
}
