package ereader

import (
	"regexp"
	"strings"
)

type TagAndData struct {
	Tag  string
	Data string
}

func RemoveTag(str string) string {
	rep := regexp.MustCompile(`<("[^"]*"|'[^']*'|[^'">])*>`)
	str = rep.ReplaceAllString(str, "")
	return str
}

func ConvNewline(str, nlcode string) string {
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}

func GetTagHead(str string) []TagAndData {
	rep := regexp.MustCompile(`^*<([a-z\d]*)`)
	// <h1 id="contentIndex_chap_1">test</h1>
	str = ConvNewline(str, "\n")

	sp := strings.Split(str, "\n")
	tad := make([]TagAndData, 0)

	for _, s := range sp {
		t := TagAndData{}
		ss := rep.FindAllString(s, -1)
		if len(ss) == 0 {
			t.Tag = ""
			s = strings.TrimSpace(s)
			t.Data = s
			tad = append(tad, t)
			continue
		}

		t.Tag = ss[0][1:]
		t.Data = strings.TrimSpace(RemoveTag(s))
		tad = append(tad, t)
	}

	// str = rep.ReplaceAllString(str, "")
	return tad
}
