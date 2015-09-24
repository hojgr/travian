package main

import (
	"fmt"
	"time"

	"github.com/hojgr/travian/resources"
	"github.com/hojgr/travian/web"
)

func main() {
	web := web.NewClient("http://s5.zravian.com")
	web.Login("bond", "changeme")

	fmt.Println("Cookie: " + web.GetCookie())
	time.Sleep(5 * time.Second)
	for {
		villageResp, _ := web.GetVillage1HTML()

		composeResp, _ := web.GetComposeMessageHTML()
		key, _ := web.GetActionKey(composeResp)

		fields := resources.GetFields(villageResp)
		lowestField := resources.GetLowestLevelField(fields)

		web.UpgradeField(lowestField.Id, key)

		fmt.Printf("Upgrading %s to level %d\n", lowestField.Name, lowestField.Level+1)
	}
}
