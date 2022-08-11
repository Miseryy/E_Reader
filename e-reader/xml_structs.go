package ereader

import "encoding/xml"

/*
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

type Nav1 struct {
	XMLName xml.Name `xml:"html"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Epub    string   `xml:"epub,attr"`
	Head    struct {
		Text string `xml:",chardata"`
		Meta []struct {
			Text    string `xml:",chardata"`
			Charset string `xml:"charset,attr"`
			Name    string `xml:"name,attr"`
			Content string `xml:"content,attr"`
		} `xml:"meta"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
			Href string `xml:"href,attr"`
		} `xml:"link"`
	} `xml:"head"`
	Body struct {
		Text string `xml:",chardata"`
		Nav  []struct {
			Text   string `xml:",chardata"`
			Type   string `xml:"type,attr"`
			ID     string `xml:"id,attr"`
			Hidden string `xml:"hidden,attr"`
			Xhtml  string `xml:"xhtml,attr"`
			H1     struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"h1"`
			Ol struct {
				Text  string `xml:",chardata"`
				Class string `xml:"class,attr"`
				Li    []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					A    struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
						Type string `xml:"type,attr"`
					} `xml:"a"`
				} `xml:"li"`
			} `xml:"ol"`
		} `xml:"nav"`
	} `xml:"body"`
}

type Html struct {
	XMLName xml.Name `xml:"html"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Epub    string   `xml:"epub,attr"`
	Lang    string   `xml:"lang,attr"`
	Head    struct {
		Text string `xml:",chardata"`
		Link struct {
			Text string `xml:",chardata"`
			Rel  string `xml:"rel,attr"`
			Href string `xml:"href,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Title string `xml:"title"`
	} `xml:"head"`
	Body struct {
		Text string `xml:",chardata"`
		Nav  struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			H1   string `xml:"h1"`
			Ol   struct {
				Text string `xml:",chardata"`
				Li   []struct {
					Text string `xml:",chardata"`
					A    struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"a"`
					Ol struct {
						Text string `xml:",chardata"`
						Li   []struct {
							Text string `xml:",chardata"`
							A    struct {
								Text string `xml:",chardata"`
								Href string `xml:"href,attr"`
							} `xml:"a"`
						} `xml:"li"`
					} `xml:"ol"`
				} `xml:"li"`
			} `xml:"ol"`
		} `xml:"nav"`
	} `xml:"body"`
}
*/

type Container struct {
	XMLName   xml.Name `xml:"container"`
	Xmlns     string   `xml:"xmlns,attr"`
	Version   string   `xml:"version,attr"`
	Rootfiles []struct {
		Rootfile struct {
			FullPath  string `xml:"full-path,attr"`
			MediaType string `xml:"media-type,attr"`
		} `xml:"rootfile"`
	} `xml:"rootfiles"`
}

type Content struct {
	XMLName  xml.Name `xml:"package"`
	Xmlns    string   `xml:"xmlns,attr"`
	Version  string   `xml:"version,attr"`
	Metadata struct {
		Dc         string `xml:"dc,attr"`
		Opf        string `xml:"opf,attr"`
		Identifier string `xml:"identifier"`
		Title      string `xml:"title"`
		Date       string `xml:"date"`
		Language   string `xml:"language"`
		Meta       struct {
			Name string `xml:"name,attr"`
			Text string `xml:",chardata"`
		} `xml:"meta"`
	} `xml:"metadata"`
	Items []struct {
		ID         string `xml:"id,attr"`
		Href       string `xml:"href,attr"`
		MediaType  string `xml:"media-type,attr"`
		Properties string `xml:"properties,attr"`
	} `xml:"manifest>item"`
	Spine struct {
		Toc      string `xml:"toc,attr"`
		Itemrefs []struct {
			Idref  string `xml:"idref,attr"`
			Linear string `xml:"linear,attr"`
		} `xml:"itemref"`
	} `xml:"spine"`
}

type Nav struct {
	Title string   `xml:"head>title"`
	Xmlns string   `xml:"xmlns,attr"`
	Epub  string   `xml:"epub,attr"`
	Lang  string   `xml:"lang,attr"`
	Id    []string `xml:"html>id"`
	Nav   []struct {
		Type string `xml:"type,attr"`
		H1   string `xml:"h1"`
		Li   []struct {
			Li []struct {
				A struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"a"`
			} `xml:"ol>li"`
			A struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"a"`
		} `xml:"ol>li"`
	} `xml:"body>nav"`
}

type Chapter struct {
	Title string `xml:"head>title"`
	Body  struct {
		Data string `xml:",innerxml"`
	} `xml:"body"`
}
