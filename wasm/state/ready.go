package state

import "syscall/js"
import "html/template"
import "bytes"

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

	test := `<div>hi there {{index . "test"}}</div>`

	vars := map[string]any{"test": 123}
	t, _ := template.New("markup").Parse(test)
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	cb := content.Bytes()

	list.Set("innerHTML", string(cb))

	return js.Undefined()
}
