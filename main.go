package main

import (
	"fmt"

	"github.com/hojgr/travian/web"
)

func main() {
	web := web.NewClient("http://s5.zravian.com")
	web.Login("bond", "changeme")

	fmt.Println("Cookie: " + web.GetCookie())

	farmList := [...][4]string{
		[4]string{"5", "15", "11", "100"},
		[4]string{"5", "14", "11", "100"},
		[4]string{"3", "14", "11", "100"},
		[4]string{"6", "11", "11", "100"},
		[4]string{"10", "10", "11", "100"},
		[4]string{"1", "-1", "11", "100"},
	}

	counter := 0 // max 5

	for {
		v := farmList[counter]

		counter++
		if counter == 5 {
			counter = 0
		}

		resp, _ := web.GetComposeMessageHTML()
		key, _ := web.GetActionKey(resp)
		canAttack := web.CanAttack(v[0], v[1], key)

		if canAttack {
			resp, _ := web.GetComposeMessageHTML()
			key, _ := web.GetActionKey(resp)

			web.Raid(v[0], key)

		}
	}
}
