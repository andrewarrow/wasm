package app

import (
	"net/http"

	"github.com/andrewarrow/feedback/router"
)

func handleSaveCreate(c *router.Context) {
	c.ReadJsonBodyIntoParams()

	returnPath := "/"

	message := c.Insert("wasm")
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}
