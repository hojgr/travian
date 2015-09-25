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

	for {
		villageResp, _ := web.GetVillage1HTML()
		queue := web.GetBuildingQueue(villageResp)

		if len(queue.Tasks) > 0 {
			for _, e := range queue.Tasks {
				fmt.Printf("%s is upgrading from %d to %d. Will finish in %d\n",
					e.Name, e.OldLevel, e.NewLevel, e.TimeLeft)
				time.Sleep(1 * time.Second)
			}
		} else {

			composeResp, _ := web.GetComposeMessageHTML()
			key, _ := web.GetActionKey(composeResp)

			villageResp, _ = web.GetVillage1HTML()
			fields := resources.GetFields(villageResp)
			lowestField := resources.GetLowestLevelField(fields)

			web.UpgradeField(lowestField.Id, key)

			fmt.Printf("Upgrading %s to level %d\n", lowestField.Name, lowestField.Level+1)
		}
	}
}
