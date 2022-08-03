package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/taylorskalyo/goreader/epub"
	"golang.org/x/net/html"
)

func my_teacher() {
	// rc, err := epub.OpenReader("./test.epub")
	// rc, err := epub.OpenReader("./mybook.epub")
	// rc, err := epub.OpenReader("./PrideAndPrejudice.epub")
	rc, err := epub.OpenReader("./gon_sample.epub")
	if err != nil {
		panic(err)
	}

	defer rc.Close()
	book := rc.Rootfiles
	fmt.Println("Root Count ", len(book))
	a := book[0]

	for i, itemref := range a.Spine.Itemrefs {
		fmt.Println(i + 1)
		fmt.Println(itemref.HREF)
	}

	return
	scanner := bufio.NewScanner(os.Stdin)
	for i, itemref := range a.Spine.Itemrefs {
		fmt.Println("Itemrefs ", i)
		f, err := itemref.Open()

		if err != nil {
			panic(err)
		}

		// fmt.Println(f)
		// fmt.Println(a.Manifest.Items)
		toke := html.NewTokenizer(f)
		count := 1
		fmt.Println(itemref.HREF)
		for {
			fmt.Println("Count ", count)
			ty := toke.Next()
			toke_n := toke.Token()
			_ = toke_n

			// fmt.Println("ty ", ty)
			fmt.Println("token ", toke_n)
			fmt.Println("token2 ", toke_n.DataAtom.String())
			// fmt.Println("attr ", toke_n.Attr)

			// for _, item := range toke_n.Attr {
			// 	println("key ", item.Key)
			// 	println("val ", item.Val)
			// 	println("name ", item.Namespace)
			// }

			switch ty {
			case html.StartTagToken:
				// fmt.Println("start ", toke_n.Data)
			case html.EndTagToken:
				// fmt.Println("end ", toke_n.Data)
			case html.ErrorToken:
				err = toke.Err()
			case html.TextToken:
				// A=LF
				// fmt.Printf("%#U", []byte(toke_n.Data)[0])
				// fmt.Println(toke_n.Data)
				// fmt.Println("end")
			}

			if err == io.EOF {
				fmt.Println("***********************")
				fmt.Println("*********EOF***********")
				fmt.Println("***********************")
				scanner.Scan()
				break
			} else if err != nil {
				fmt.Println("***********************")
				fmt.Println("*********ERR***********")
				fmt.Println("***********************")
				break
			}

			count++
		}
	}

	// fmt.Println(book.Title)

	// for _, item := range book.Spine.Itemrefs {
	// 	fmt.Println(item.ID)
	// 	fmt.Println(item.HREF)
	// 	fmt.Println(item.Item)

	// }

	// e := g_ep.NewEpub("TestE")
	// e.SetAuthor("TEST AUTH")

	// section1b := `<h1>Sect 1</h1>
	// <p>this is a pen</p>`
	// e.AddSection(section1b, "Sect1", "", "")
	// err := e.Write("test.epub")
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

const meta_container_path = "META-INF/container.xml"

func test() {
	f, err := os.Open("./gon_sample.epub")
	if err != nil {
		return
	}

	fi, err := f.Stat()
	if err != nil {
		return
	}

	//fmt.Println(fi.Size())
	z, err := zip.NewReader(f, fi.Size())
	if err != nil {
		return
	}
	//fmt.Println(z)

	files := make(map[string]*zip.File)
	for _, ff := range z.File {
		files[ff.Name] = ff
	}

	f2, err := files[meta_container_path].Open()
	if err != nil {
		return
	}

	var b bytes.Buffer
	_, err = io.Copy(&b, f2)
	if err != nil {
		return
	}

	c := new(Container)
	err = xml.Unmarshal(b.Bytes(), &c)
	full_path := c.Rootfiles.Rootfile.FullPath
	fmt.Println("FULL ", full_path)
	dir := strings.Split(full_path, "/")[0]

	f3, err := files[full_path].Open()
	if err != nil {
		return
	}

	var bb bytes.Buffer
	_, err = io.Copy(&bb, f3)
	if err != nil {
		return
	}

	pp := new(Package)
	err = xml.Unmarshal(bb.Bytes(), &pp)
	// fmt.Print(string(bb.Bytes()))

	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0
	for i, v := range pp.Manifest.Item {
		fmt.Printf("ID %s, Href %s\n", v.ID, v.Href)
		count = i
		if v.ID == "nav" {
			break
		}
	}

	d := dir + "/" + pp.Manifest.Item[count].Href
	fmt.Println(d)

	f4, err := files[d].Open()

	if err != nil {
		fmt.Println(err)
	}

	var bb4 bytes.Buffer
	_ = bb4

	fmt.Println("bb4")
	_, err = io.Copy(&bb4, f4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("nav")
	// nav
	err = xml.Unmarshal(bb4.Bytes(), &pp)

	fmt.Println("parse")

	nav_st := new(Html)

	err = xml.Unmarshal(bb4.Bytes(), &nav_st)

	// fmt.Println(nav_st.Body.Nav)
	fmt.Println(nav_st.Body.Nav.Ol)
	for _, vv := range nav_st.Body.Nav.Ol.Li {
		for u, li := range vv.Ol.Li {
			fmt.Println(u, li.Text)

		}
	}
}

func test2() {
	f, err := os.Open("./mybook.epub")

	if err != nil {
		panic(err)
	}

	fi, err := f.Stat()
	if err != nil {
		return
	}

	//fmt.Println(fi.Size())
	z, err := zip.NewReader(f, fi.Size())
	if err != nil {
		return
	}

	files := make(map[string]*zip.File)
	for _, ff := range z.File {
		files[ff.Name] = ff
	}

	container, err := OpenFile(files, meta_container_path)
	// write(b, "contain")
	if err != nil {
		return
	}

	container_st := new(Container01)
	err = xml.Unmarshal(container, container_st)
	if err != nil {
		return
	}
	full_path := container_st.Rootfiles.Rootfile.FullPath
	fmt.Println(full_path)
	dir := strings.Split(full_path, "/")[0]

	content_buff, err := OpenFile(files, full_path)
	if err != nil {
		return
	}

	// write(content_buff, "content")
	content_st := new(Content)
	xml.Unmarshal(content_buff, content_st)
	var nav_path string

	for _, item := range content_st.Manifest.Items {
		if item.ID == "nav" {
			nav_path = dir + "/" + item.Href
		}
	}

	fmt.Println(nav_path)

	nav_buff, err := OpenFile(files, nav_path)
	if err != nil {
		return
	}
	nav_st := new(Nav)
	xml.Unmarshal(nav_buff, nav_st)
	// write(nav_buff, "nav")
	for _, v := range nav_st.Body.Nav {
		if v.Type == "toc" {
			fmt.Println(v.Ol.Li)
		}
	}

}

func OpenFile(files map[string]*zip.File, path string) ([]byte, error) {
	content_file, err := files[path].Open()
	if err != nil {
		return nil, nil
	}
	var b bytes.Buffer
	_, err = io.Copy(&b, content_file)
	return b.Bytes(), nil
}

func main() {
	test2()

}

func write(b []byte, name string) {
	f, err := os.Create(name + ".xml")
	if err != nil {
		panic(err)
	}
	f.Write(b)
	f.Close()
}
