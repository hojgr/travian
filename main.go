package main

import (
	"fmt"
	"time"

	"github.com/hojgr/travian/web"
)

func main() {
	web := web.NewClient("http://s5.zravian.com")
	web.Login("bond", "changeme")

	fmt.Println("Cookie: " + web.GetCookie())

	farmList := [...]string{
		"1016",
		"1015",
		"615",
		"1212",
		"2011",
	}

	counter := 0 // max 5

	for {
		resp, _ := web.GetComposeMessageHTML()
		key, _ := web.GetActionKey(resp)

		v := farmList[counter]
		counter++

		web.Raid(v, key)

		if counter == 5 {
			counter = 0
			time.Sleep(5 * time.Second)
		}
	}

}
