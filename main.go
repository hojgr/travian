package main

import (
	"fmt"

	"github.com/hojgr/travian/resources"
	"github.com/hojgr/travian/web"
)

func main() {
	web := web.NewClient("http://s5.zravian.com")
	web.Login("bond", "changeme")

	fmt.Println("Cookie: " + web.GetCookie())

	resp, _ := web.GetVillage1HTML()

	fields := resources.GetFields(resp)
	fmt.Printf("Fields: %d, First name: %s", len(fields), fields[0].Name)
}
