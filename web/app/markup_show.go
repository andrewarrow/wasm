package app

import (
	"io/ioutil"
	"net/http"

	"github.com/andrewarrow/feedback/router"
)

func handleMarkupShow(c *router.Context, name string) {
	asBytes, _ := ioutil.ReadFile("views/" + name)
	contentType := "text/plain"
	c.Writer.Header().Set("Content-Type", contentType)
	c.Writer.Write(asBytes)
}

func handleMarkupShowPost(c *router.Context, guid string) {
	c.ReadFormValuesIntoParams("file")
	returnPath := "/"

	//c.ValidateCreate("markup")
	message := c.Update("markup", "where guid=", guid)
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath+"/"+guid, 302)
		return
	}
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}
