package app

import (
	"html/template"

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
	one := c.One("wasm", "order by created_at desc")

	send := map[string]any{}
	send["content"] = template.HTML(`<div class="p-3 border border-black">hi </div>`)
	if len(one) > 0 {
		//send["content"] = template.HTML(one["content"].(string))
	}
	c.SendContentInLayout("welcome.html", send, 200)
}
