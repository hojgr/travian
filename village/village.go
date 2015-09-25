package village

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Parcel struct {
	Id       int
	Empty    bool
	Building Building
}

type Building struct {
	Name  string
	Level int
}

type Village struct {
	Buildings []Building
}

func GetBuildings(resp *http.Response) Village {
	doc, _ := goquery.NewDocumentFromResponse(resp)

	village := Village{}

	doc.Find("map#map2 area").EachWithBreak(func(index int, s *goquery.Selection) bool {
		parcelId := getIdFromURL(s.AttrOr("href", "fail"))

		name, level, empty := parseNameAndLevel(s.AttrOr("alt", "fail"))

		if empty {
			village.Buildings = append(village.Buildings, Parcel{
				Id:    parcelId,
				Empty: true,
			})
		} else {
			village.Buildings = append(village.Buildings, Parcel{
				Id:    parcelId,
				Empty: false,
				Building: Building{
					Name:  name,
					Level: level,
				},
			})
		}

		if parcelId == 40 { // 40 = Wall; it is repeated 3 times, this limits it to one
			return false
		} else {
			return true
		}
	})

	return village
}

func getIdFromURL(url string) int {
	idRegexp := regexp.MustCompile("id=([0-9]+)")

	matches := idRegexp.FindAllStringSubmatch(url, 1)

	id, _ := strconv.Atoi(matches[0][1])

	return id
}

func parseNameAndLevel(nameAndLevel string) (name string, level int, empty bool) {
	r := regexp.MustCompile("(.*) level ([0-9]+)")

	matches := r.FindAllStringSubmatch(nameAndLevel, 1)

	if len(matches) == 0 {
		return "", 0, true
	} else {
		level, _ := strconv.Atoi(matches[0][2])
		return matches[0][1], level, false
	}

	return "", 0, true
}
