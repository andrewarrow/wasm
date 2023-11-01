package state

import "syscall/js"
import "html/template"
import "bytes"
import "embed"

import "fmt"

var EmbeddedTemplates embed.FS

type State struct {
}

func NewState() *State {
	s := State{}
	return &s
}

func (e *State) WasmReady2(this js.Value, p []js.Value) any {
	return js.Undefined()
}

func (e *State) WasmReady(this js.Value, p []js.Value) any {
	list := js.Global().Get("document").Call("getElementById", "list")

	templateText, err := EmbeddedTemplates.ReadFile("views/" + "list.html")
	fmt.Println(err, string(templateText))

	listItems := []string{"", "", "", ""}
	vars := map[string]any{}
	vars["list"] = listItems

	t, err := template.New("markup").Parse(string(templateText))
	fmt.Println(err, t)
	content := new(bytes.Buffer)
	t.Execute(content, vars)
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
