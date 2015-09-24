package statistics

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Statistics for Village
type Statistics struct {
	Position     int
	Username     string
	Alliance     string
	Population   int
	VillageCount int
}

// GetStatistics retrieves statistics
func GetStatistics(resp *http.Response) Statistics {
	doc, _ := goquery.NewDocumentFromResponse(resp)

	s := doc.Find("tr.hl").First()

	positionStr := strings.TrimRight(s.Find(".ra").Text(), ".")
	position, _ := strconv.Atoi(positionStr)

	username := strings.Trim(s.Find(".pla").Find("a").Text(), " ")
	alliance := s.Find("al").Text()

	popStr := s.Find(".pop").Text()
	pop, _ := strconv.Atoi(popStr)

	villageCountStr := s.Find(".vil").Text()
	villageCount, _ := strconv.Atoi(villageCountStr)

	stats := Statistics{
		position,
		username,
		alliance,
		pop,
		villageCount,
	}

	return stats
}
