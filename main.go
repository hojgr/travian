package main

import (
	"fmt"

	"github.com/hojgr/travian/village"
	"github.com/hojgr/travian/web"
)

func main() {
	web := web.NewClient("http://s5.zravian.com")
	web.Login("bond", "changeme")

	fmt.Println("Cookie: " + web.GetCookie())

	resp, _ := web.GetVillage2HTML()
	village.GetBuildings(resp)
}
