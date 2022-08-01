package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/taylorskalyo/goreader/epub"
	"golang.org/x/net/html"
)

type Container struct {
	XMLName   xml.Name `xml:"container"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Version   string   `xml:"version,attr"`
	Rootfiles struct {
		Text     string `xml:",chardata"`
		Rootfile struct {
			Text      string `xml:",chardata"`
			FullPath  string `xml:"full-path,attr"`
			MediaType string `xml:"media-type,attr"`
		} `xml:"rootfile"`
	} `xml:"rootfiles"`
}

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

const contain_path = "META-INF/container.xml"

func main() {
	f, err := os.Open("./test.epub")
	if err != nil {
		return
	}

	fi, err := f.Stat()
	if err != nil {
		return
	}

	fmt.Println(fi.Size())
	z, err := zip.NewReader(f, fi.Size())
	if err != nil {
		return
	}
	fmt.Println(z)

	files := make(map[string]*zip.File)
	for _, ff := range z.File {
		files[ff.Name] = ff
	}
	fmt.Println(files)

	f2, err := files[contain_path].Open()
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
	fmt.Println(c.Rootfiles.Rootfile.FullPath)

}
