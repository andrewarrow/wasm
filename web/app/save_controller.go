package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleSave(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleSaveIndex(c)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		handleSaveCreate(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handleSaveShow(c, second)
		return
	}
	if second != "" && third == "" && c.Method == "POST" {
		handleSaveShowPost(c, second)
		return
	}
	if router.NotLoggedIn(c) {
		return
	}
	c.NotFound = true
}

func handleSaveIndex(c *router.Context) {
	//list := c.All("save", "where user_id=$1 order by created_at desc", "", c.User["id"])

	send := map[string]any{}
	c.SendContentInLayout(".html", send, 200)
}
