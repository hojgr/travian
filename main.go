package main

import (
	"fmt"

	"github.com/hojgr/travian/statistics"
	"github.com/hojgr/travian/web"
)

func main() {
	web := web.NewClient("http://s5.zravian.com")
	web.Login("bond", "changeme")

	fmt.Println("Cookie: " + web.GetCookie())

	statsHtml, _ := web.GetStatisticsHTML()
	stats := statistics.GetStatistics(statsHtml)

	fmt.Printf("Position: %d\nPop: %d\n", stats.Position, stats.Population)

}
