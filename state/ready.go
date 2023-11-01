package state

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"syscall/js"
	"wasm/network"
)

var EmbeddedTemplates embed.FS

type State struct {
}

func NewState() *State {
	s := State{}
	return &s
}

func dollar(name string) js.Value {
	return js.Global().Get("document").Call("getElementById", name)
}

func (e *State) Click(this js.Value, p []js.Value) any {
	id := this.Get("id").String()
	fmt.Println(id)

	modal := dollar("modal")
	if id == "b1" {
		removeClass(modal, "translate-x-full")
		removeClass(modal, "opacity-0")
		modal.Set("scrollTop", 0)
		go func() {
			modal.Set("innerHTML", runTemplate("form"))
		}()
	} else if id == "b2" {
		removeClass(modal, "translate-x-full")
		removeClass(modal, "opacity-0")
		modal.Set("scrollTop", 0)
		go func() {
			modal.Set("innerHTML", runTemplate("other_form"))
		}()
	} else if id == "cancel" {
		addClass(modal, "translate-x-full")
		addClass(modal, "opacity-0")
	} else if id == "save" {
		go network.Save()
		addClass(modal, "translate-x-full")
		addClass(modal, "opacity-0")
	}
	cancel := dollar("cancel")
	cancel.Set("onclick", js.FuncOf(e.Click))
	save := dollar("save")
	save.Set("onclick", js.FuncOf(e.Click))
	return js.Undefined()
}

func runTemplate(name string) string {
	templateText, _ := network.GetTemplate(name)

	vars := map[string]any{}

	t, _ := template.New("markup").Parse(string(templateText))
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	t.ExecuteTemplate(content, name, vars)
	cb := content.Bytes()
	return string(cb)
}

func runTemplate2(name string) string {
	templateText, _ := EmbeddedTemplates.ReadFile("views/" + name + ".html")

	vars := map[string]any{}

	t, _ := template.New("markup").Parse(string(templateText))
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	t.ExecuteTemplate(content, name, vars)
	cb := content.Bytes()
	return string(cb)
}

func (e *State) WasmReady(this js.Value, p []js.Value) any {
	list := js.Global().Get("document").Call("getElementById", "list")
	b1 := js.Global().Get("document").Call("getElementById", "b1")
	b2 := js.Global().Get("document").Call("getElementById", "b2")
	b1.Set("onclick", js.FuncOf(e.Click))
	b2.Set("onclick", js.FuncOf(e.Click))

	templateText, _ := EmbeddedTemplates.ReadFile("views/" + "list.html")

	listItems := []string{"", "", "", ""}
	vars := map[string]any{}
	vars["list"] = listItems

	t, _ := template.New("markup").Parse(string(templateText))
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	t.ExecuteTemplate(content, "list", vars)
	cb := content.Bytes()

	list.Set("innerHTML", string(cb))

	return js.Undefined()
}

/*
func MyGoFunc() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		requestUrl := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]

			go func() {
				res, err := http.DefaultClient.Get(requestUrl)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				arrayConstructor := js.Global().Get("Uint8Array")
				dataJS := arrayConstructor.New(len(data))
				js.CopyBytesToJS(dataJS, data)

				responseConstructor := js.Global().Get("Response")
				response := responseConstructor.New(dataJS)

				resolve.Invoke(response)
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}*/
