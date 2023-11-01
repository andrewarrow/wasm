package state

import "syscall/js"
import "html/template"
import "bytes"
import "wasm/network"

type State struct {
}

func NewState() *State {
	s := State{}
	return &s
}

func (e *State) WasmReady(this js.Value, p []js.Value) any {
	return js.Undefined()
}

func (e *State) WasmReady2(this js.Value, p []js.Value) any {
	list := js.Global().Get("document").Call("getElementById", "list")

	templateText := network.GetTemplate("list.html")

	list := []string{"", "", ""}
	vars := map[string]any{}
	vars["list"] = list

	t, _ := template.New("markup").Parse(templateText)
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	cb := content.Bytes()

	list.Set("innerHTML", string(cb))

	return js.Undefined()
}
