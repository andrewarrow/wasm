package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleWelcome(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleWelcomeIndex(c)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		handleWelcomeCreate(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handleWelcomeShow(c, second)
		return
	}
	if second != "" && third == "" && c.Method == "POST" {
		handleWelcomeShowPost(c, second)
		return
	}
	if router.NotLoggedIn(c) {
		return
	}
	c.NotFound = true
}

func handleWelcomeIndex(c *router.Context) {
	//list := c.All("welcome", "where user_id=$1 order by created_at desc", "", c.User["id"])

	send := map[string]any{}
	c.SendContentInLayout("welcome.html", send, 200)
}
