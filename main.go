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

type Package struct {
	XMLName          xml.Name `xml:"package"`
	Text             string   `xml:",chardata"`
	Xmlns            string   `xml:"xmlns,attr"`
	UniqueIdentifier string   `xml:"unique-identifier,attr"`
	Version          string   `xml:"version,attr"`
	Metadata         struct {
		Text       string `xml:",chardata"`
		Dc         string `xml:"dc,attr"`
		Identifier struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"identifier"`
		Title    string `xml:"title"`
		Language string `xml:"language"`
		Creator  struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
		} `xml:"creator"`
		Meta []struct {
			Text     string `xml:",chardata"`
			Refines  string `xml:"refines,attr"`
			Property string `xml:"property,attr"`
			Scheme   string `xml:"scheme,attr"`
			ID       string `xml:"id,attr"`
		} `xml:"meta"`
	} `xml:"metadata"`
	Manifest struct {
		Text string `xml:",chardata"`
		Item []struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			Href       string `xml:"href,attr"`
			MediaType  string `xml:"media-type,attr"`
			Properties string `xml:"properties,attr"`
		} `xml:"item"`
	} `xml:"manifest"`
	Spine struct {
		Text    string `xml:",chardata"`
		Toc     string `xml:"toc,attr"`
		Itemref struct {
			Text  string `xml:",chardata"`
			Idref string `xml:"idref,attr"`
		} `xml:"itemref"`
	} `xml:"spine"`
}

type aPackage struct {
	XMLName          xml.Name `xml:"package"`
	Text             string   `xml:",chardata"`
	Xmlns            string   `xml:"xmlns,attr"`
	UniqueIdentifier string   `xml:"unique-identifier,attr"`
	Version          string   `xml:"version,attr"`
	Lang             string   `xml:"lang,attr"`
	Metadata         struct {
		Text       string `xml:",chardata"`
		Dc         string `xml:"dc,attr"`
		Opf        string `xml:"opf,attr"`
		Identifier struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Scheme string `xml:"scheme,attr"`
		} `xml:"identifier"`
		Meta []struct {
			Text     string `xml:",chardata"`
			Refines  string `xml:"refines,attr"`
			Property string `xml:"property,attr"`
			Name     string `xml:"name,attr"`
			Content  string `xml:"content,attr"`
		} `xml:"meta"`
		Title     string `xml:"title"`
		Language  string `xml:"language"`
		Creator   string `xml:"creator"`
		Publisher string `xml:"publisher"`
		Date      string `xml:"date"`
	} `xml:"metadata"`
	Manifest struct {
		Text string `xml:",chardata"`
		Item []struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			Href       string `xml:"href,attr"`
			MediaType  string `xml:"media-type,attr"`
			Properties string `xml:"properties,attr"`
		} `xml:"item"`
	} `xml:"manifest"`
	Spine struct {
		Text                     string `xml:",chardata"`
		Toc                      string `xml:"toc,attr"`
		PageProgressionDirection string `xml:"page-progression-direction,attr"`
		Itemref                  []struct {
			Text   string `xml:",chardata"`
			Idref  string `xml:"idref,attr"`
			Linear string `xml:"linear,attr"`
		} `xml:"itemref"`
	} `xml:"spine"`
	Guide struct {
		Text      string `xml:",chardata"`
		Reference struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			Title string `xml:"title,attr"`
			Href  string `xml:"href,attr"`
		} `xml:"reference"`
	} `xml:"guide"`
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
	f, err := os.Open("./mybook.epub")
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
	//fmt.Println(files)

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
	_ = count
	d := dir + "/" + pp.Manifest.Item[count].Href
	fmt.Println(d)

	f4, err := files[d].Open()

	if err != nil {
		fmt.Println(err)
	}
	_ = f4
	var bb4 bytes.Buffer

	_, err = io.Copy(&bb4, f4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bb4.Bytes()))

}
