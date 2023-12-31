package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleMarkup(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleMarkupIndex(c)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		handleMarkupCreate(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handleMarkupShow(c, second)
		return
	}
	if second != "" && third == "" && c.Method == "POST" {
		handleMarkupShowPost(c, second)
		return
	}
	if router.NotLoggedIn(c) {
		return
	}
	c.NotFound = true
}

func handleMarkupIndex(c *router.Context) {
	//list := c.All("markup", "where user_id=$1 order by created_at desc", "", c.User["id"])

	send := map[string]any{}
	//send["content"] = template.HTML(`<div class="p-3 border border-black"> </div>`)
	//send["top"] = markup.ToHTML(send, "top.mu")
	//send["content"] = template.HTML(markup.ToHTML(send, "index.mu"))
	c.SendContentInLayout("markup.html", send, 200)
}
