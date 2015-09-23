package resources

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const (
	WOODCUTTER = iota
	CROPLAND   = iota
	CLAY_PIT   = iota
	IRON_MINE  = iota
)

type Field struct {
	Id    int
	Level int
	Name  string
}

func GetFields(resp *http.Response) []Field {
	doc, _ := goquery.NewDocumentFromResponse(resp)

	altRegex := regexp.MustCompile("(.*) level ([0-9]+)")
	hrefRegex := regexp.MustCompile("\\?id=([0-9]+)")

	fields := []Field{}

	area := doc.Find("map area")
	area.Each(func(index int, s *goquery.Selection) {
		alt, _ := s.Attr("alt")
		altParsed := altRegex.FindAllStringSubmatch(alt, -1)

		if len(altParsed) > 0 {
			name := altParsed[0][1]

			level, _ := strconv.Atoi(altParsed[0][2])

			href, _ := s.Attr("href")
			hrefParsed := hrefRegex.FindAllStringSubmatch(href, -1)

			id, _ := strconv.Atoi(hrefParsed[0][1])

			fields = append(fields, Field{id, level, name})
		}
	})

	return fields
}
