package app





import (
	"net/http"

	"github.com/andrewarrow/feedback/router"
)

func handleWelcomeShowPost(c *router.Context, guid string) {
	c.ReadFormValuesIntoParams("file")
	returnPath := "/"

	//c.ValidateCreate("welcome")
	message := c.Update("welcome", "where guid=", guid)
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath+"/"+guid, 302)
		return
	}
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}

func handleWelcomeShow(c *router.Context, guid string) {
	item := c.One("welcome", "where guid=$1", guid)
	send := map[string]any{}
	send["item"] = item
	c.SendContentInLayout(".html", send, 200)
}