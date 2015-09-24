package resources

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Granary for Granary+Crop info
type Granary struct {
	Capacity int
	Crop     int
}

// Warehouse for Warehouse info
type Warehouse struct {
	Capacity int
	Lumber   int
	Clay     int
	Iron     int
}

// Production for production info
type Production struct {
	Lumber int
	Clay   int
	Iron   int
	Crop   int
}

// Resources for resources
type Resources struct {
	Granary    Granary
	Warehouse  Warehouse
	Production Production
}

// GetResources retrieves resource info
func GetResources(resp *http.Response) Resources {
	doc, _ := goquery.NewDocumentFromResponse(resp)

	resWrapperSelector := doc.Find("div#resWrap")

	lumber := resWrapperSelector.Find("td#l4")
	clay := resWrapperSelector.Find("td#l3")
	iron := resWrapperSelector.Find("td#l2")
	crop := resWrapperSelector.Find("td#l1")

	lumberProd, lumberStored, lumberCapa := parseResource(lumber)
	clayProd, clayStored, _ := parseResource(clay)
	ironProd, ironStored, _ := parseResource(iron)
	cropProd, cropStored, cropCapa := parseResource(crop)

	return Resources{
		Granary{
			Capacity: cropCapa,
			Crop:     cropStored,
		},
		Warehouse{
			Capacity: lumberCapa,
			Lumber:   lumberStored,
			Clay:     clayStored,
			Iron:     ironStored,
		},
		Production{
			lumberProd,
			clayProd,
			ironProd,
			cropProd,
		},
	}
}

func parseResource(s *goquery.Selection) (_production, _stored, _capacity int) {
	productionStr, _ := s.Attr("title")
	production, _ := strconv.Atoi(productionStr)

	status := s.Text()
	split := strings.Split(status, "/")

	stored, _ := strconv.Atoi(split[0])
	capacity, _ := strconv.Atoi(split[1])

	return production, stored, capacity
}
