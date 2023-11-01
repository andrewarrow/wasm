package state

import "syscall/js"
import "fmt"

type State struct {
}

func NewState() *State {
	s := State{}
	return &s
}

func (e *State) WasmReady(this js.Value, p []js.Value) any {
	fmt.Println("wefwef")

	return js.Undefined()
}
