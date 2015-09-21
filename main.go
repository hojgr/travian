package main

import (
	"fmt"
	"net/url"

	"github.com/hojgr/cook/web"
)

func main() {
	web := web.NewClient("http://s5.zravian.com")

	web.POST(web.BaseURL, url.Values{
		"name": {"bond"},
		"pass": {"changeme"},
	})

	info, _ := web.GetStatisticsHTML()

	fmt.Printf("Position: %s\nPop: %s\n", info["position"], info["pop"])

}
