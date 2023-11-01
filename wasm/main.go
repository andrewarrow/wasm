package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
	"wasm/state"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Web Assembly")
	//editor := events.NewEditor()
	state := state.NewState()
	js.Global().Set("WasmReady", js.FuncOf(state.WasmReady))

	select {}
}
