package main

import (
	"embed"
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
	"wasm/state"
)

//go:embed views/*.html
var embeddedTemplates embed.FS

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Web Assembly")
	//editor := events.NewEditor()
	state.EmbeddedTemplates = embeddedTemplates
	state := state.NewState()
	js.Global().Set("WasmReady", js.FuncOf(state.WasmReady))

	select {}
}
