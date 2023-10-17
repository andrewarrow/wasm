package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
	"wasm/events"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Web Assembly")
	editor := events.NewEditor()
	js.Global().Set("key", js.FuncOf(editor.Key))

	select {}
}
